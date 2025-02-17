package config

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/mapstructure"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Env                 string          `validate:"required" yaml:"env"` // dev、test、prod
	MasterChainId       uint64          `validate:"required" yaml:"master_chain_id"`
	SubmitterPrivateKey string          `validate:"required" yaml:"submitter_private_key"`
	DB                  DataBase        `validate:"required" yaml:"db"`
	Log                 Log             `validate:"required" yaml:"log"`
	Tss                 TssConfig       `validate:"required" yaml:"tss"`
	P2P                 P2PConfig       `validate:"required" yaml:"p2p"`
	Contract            ContractAddress `validate:"required" yaml:"contract"`
	Chains              []Chain         `validate:"required" yaml:"chains"`
}

func (c *Config) IsEVM(chainId uint64) bool {
	for _, chain := range c.Chains {
		if chain.ChainId == chainId {
			return chain.ChainType == 1 // todo
		}
	}

	return false
}

type DataBase struct {
	DbRootDir string `validate:"required" yaml:"db_root_dir"`
}

type Chain struct {
	Network       string        `yaml:"network"` // devnet、testnet、mainnet
	ChainId       uint64        `validate:"required"    yaml:"chain_id"`
	ChainType     int           `validate:"required"    yaml:"chain_type"`
	Rpc           Rpc           `validate:"required"    yaml:"rpc"`
	StartHeight   int           `yaml:"start_height"`
	Confirmations int           `yaml:"confirmations"`
	MaxBlockRange int           `yaml:"max_block_range"`
	ScanInterval  time.Duration `yaml:"scan_interval"`
}

type Rpc struct {
	Url       string `validate:"required" yaml:"url"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	JwtSecret string `yaml:"jwt_secret"`
}

func (c *Config) IsProd() bool {
	return c.Env == "prod"
}

func (c *Config) IsMaster(chainId uint64) bool {
	return c.MasterChainId == chainId
}

func (c *Config) MasterChainInfo() *Chain {
	for _, chain := range c.Chains {
		if chain.ChainId == c.MasterChainId {
			return &chain
		}
	}

	panic("master chain not found")
}

func (c *Config) Validator() {
	master := c.MasterChainInfo()
	if master.ChainType != 1 { // types.ChainEthereum {
		panic("invalid master chain type")
	}

	if c.Contract.Voter == "" {
		panic("missing config for contract voter address")
	}
}

type Log struct {
	Level string `yaml:"level"`
}

type P2PConfig struct {
	Port      int      `yaml:"port"`
	BootNodes []string `yaml:"boot_nodes"`
}

type TssConfig struct {
	Threshold   int           `yaml:"threshold"`
	SignTimeout time.Duration `yaml:"sign_timeout"`
	PublicKeys  []string      `yaml:"public_keys"`
}

type ContractAddress struct {
	Voter        string `yaml:"voter"`
	Account      string `yaml:"account"`
	TaskManager  string `yaml:"task_manager"`
	Participant  string `yaml:"participant"`
	Deposit      string `yaml:"deposit"`
	AssetHandler string `yaml:"asset_handler"`
}

var (
	AppConfig           Config
	TssPublicKeys       []*ecdsa.PublicKey
	SubmitterPrivateKey *ecdsa.PrivateKey
	MasterChainId       *big.Int
)

func Submitter() common.Address {
	return crypto.PubkeyToAddress(SubmitterPrivateKey.PublicKey)
}

func InitConfig(configPath string) {
	// viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("/etc")
	viper.SetConfigFile(configPath)
	logrus.Info("load yaml config")

	err := viper.ReadInConfig()
	utils.Assert(err)
	err = viper.Unmarshal(
		&AppConfig,
		func(decoderConfig *mapstructure.DecoderConfig) { decoderConfig.TagName = "yaml" },
	)
	utils.Assert(err)

	setL2ChainId(AppConfig.MasterChainId)
	setTssPublicKeys(AppConfig.Tss.PublicKeys)

	if len(AppConfig.SubmitterPrivateKey) > 0 {
		setL2PrivateKey(AppConfig.SubmitterPrivateKey)
	} else {
		logrus.Info("load .env config")
		viper.AutomaticEnv()
		// viper.Debug()
		viper.SetConfigType("env")
		viper.AddConfigPath(".")

		setL2PrivateKey(viper.GetString("SUBMITTER_PRIVATE_KEY"))
	}

	setLogLevel()
}

func setTssPublicKeys(tssList []string) {
	tssPublicKeys, err := ParseECDSAPublicKeys(tssList)
	if err != nil {
		logrus.Fatalf("Failed to parse tss public keys: %v", err)
	}

	TssPublicKeys = tssPublicKeys
}

func setL2ChainId(chainId uint64) {
	MasterChainId = big.NewInt(int64(chainId))
}

func setL2PrivateKey(pk string) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		logrus.Fatalf("Failed to load l2 private key: %v, given length %d", err, len(pk))
	}

	SubmitterPrivateKey = privateKey
}

func setLogLevel() {
	logrus.SetOutput(os.Stdout)

	logLvl, err := logrus.ParseLevel(AppConfig.Log.Level)
	if err != nil {
		logLvl = logrus.WarnLevel
	}

	logrus.SetLevel(logLvl)

	if !AppConfig.IsProd() {
		logrus.SetReportCaller(true)
	}
}

// ParseECDSAPublicKeys parses a comma-separated string of 132-character public keys with '0x' prefix
// into an array of *ecdsa.PublicKey. It uses the secp256k1 elliptic curve.
func ParseECDSAPublicKeys(publicKeyHexArray []string) ([]*ecdsa.PublicKey, error) {
	publicKeys := make([]*ecdsa.PublicKey, len(publicKeyHexArray))

	for i, keyHex := range publicKeyHexArray {
		if len(keyHex) != 66 {
			return nil, errors.New("invalid compressed public key length, expected 33 bytes")
		}

		pubBytes, err := hex.DecodeString(keyHex)
		if err != nil {
			return nil, errors.New("failed to decode public key hex: " + err.Error())
		}

		pubKey, err := crypto.DecompressPubkey(pubBytes)
		if err != nil {
			return nil, errors.New("failed to decompress public key: " + err.Error())
		}

		publicKeys[i] = pubKey
	}

	return publicKeys, nil
}
