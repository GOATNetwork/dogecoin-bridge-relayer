package tss

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	tssCrypto "github.com/bnb-chain/tss-lib/v2/crypto"
	ecdsaKeygen "github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	eddsaKeygen "github.com/bnb-chain/tss-lib/v2/eddsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/edwards/v2"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/nuvosphere/nudex-voter/internal/crypto"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/types/address"
	"github.com/nuvosphere/nudex-voter/internal/types/party"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	log "github.com/sirupsen/logrus"
)

type PartyData struct {
	basePath string
	rw       sync.RWMutex
	datas    map[crypto.CurveType]*LocalPartySaveData
}

func NewPartyData(basePath string) *PartyData {
	return &PartyData{
		basePath: basePath,
		rw:       sync.RWMutex{},
		datas:    make(map[crypto.CurveType]*LocalPartySaveData),
	}
}

func (p *PartyData) ECDSALocalData() *LocalPartySaveData {
	return p.GetData(crypto.ECDSA)
}

func (p *PartyData) EDDSALocalData() *LocalPartySaveData {
	return p.GetData(crypto.EDDSA)
}

func (p *PartyData) GetData(ec crypto.CurveType) *LocalPartySaveData {
	p.rw.RLock()
	data, ok := p.datas[ec]
	p.rw.RUnlock()

	if ok {
		return data
	}

	p.rw.Lock()
	data, err := p.loadTSSData(ec)
	utils.Assert(err)

	p.datas[ec] = data
	p.rw.Unlock()

	return data
}

func (p *PartyData) GetDataByChain(chainType uint8) *LocalPartySaveData {
	return p.GetData(types.GetCurveTypeByChain(chainType))
}

func (p *PartyData) GenerateNewLocalPartySaveData(ec crypto.CurveType, parties types.Participants) *LocalPartySaveData {
	switch ec {
	case crypto.ECDSA:
		save := ecdsaKeygen.NewLocalPartySaveData(parties.Len())
		localData := p.EDDSALocalData()

		if localData != nil && localData.ECDSAData() != nil {
			save.LocalPreParams = localData.ECDSAData().LocalPreParams // new node join party
		}

		BuildECDSALocalPartySaveData().SetData(&save)
	case crypto.EDDSA:
		save := eddsaKeygen.NewLocalPartySaveData(parties.Len())
		return BuildEDDSALocalPartySaveData().SetData(&save)
	}

	return nil
}

func (p *PartyData) LoadData() bool {
	p.rw.Lock()
	defer p.rw.Unlock()

	data, err := p.loadTSSData(crypto.ECDSA)
	if err != nil {
		return false
	}

	p.datas[data.CurveType()] = data

	data, err = p.loadTSSData(crypto.EDDSA)
	if err != nil {
		return false
	}

	p.datas[data.CurveType()] = data

	return true
}

func (p *PartyData) SaveLocalData(data *LocalPartySaveData) error {
	p.rw.Lock()
	defer p.rw.Unlock()
	p.datas[data.CurveType()] = data

	return p.saveTSSData(data)
}

func (p *PartyData) saveTSSData(data *LocalPartySaveData) error {
	curveType := data.CurveType()

	dataDir := filepath.Join(p.basePath, "tss_data", curveType.CurveName())
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		log.Errorf("Failed to create TSS data directory: %v", err)
		return err
	}

	dataBytes := data.GetJsonData()

	filePath := filepath.Join(dataDir, "tss_key_data.json")
	if err := os.WriteFile(filePath, dataBytes, 0o600); err != nil {
		log.Errorf("Failed to save TSS data to file: %v", err)
		return err
	}

	log.Infof("TSS data successfully saved to: %s", filePath)

	return nil
}

func (p *PartyData) loadTSSData(ec crypto.CurveType) (*LocalPartySaveData, error) {
	filePath := filepath.Join(p.basePath, "tss_data", ec.CurveName(), "tss_key_data.json")

	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read TSS data file: %w", err)
	}

	switch ec {
	case crypto.ECDSA:
		var data ecdsaKeygen.LocalPartySaveData
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return nil, fmt.Errorf("unable to deserialize TSS data: %w", err)
		}

		return BuildECDSALocalPartySaveData().SetData(&data), nil
	case crypto.EDDSA:
		var data eddsaKeygen.LocalPartySaveData
		if err := json.Unmarshal(dataBytes, &data); err != nil {
			return nil, fmt.Errorf("unable to deserialize TSS data: %w", err)
		}

		return BuildEDDSALocalPartySaveData().SetData(&data), nil
	}

	return nil, fmt.Errorf("unknown elliptic curve")
}

var (
	ZeroSessionID           party.SessionID
	senateSessionID         = ethCrypto.Keccak256Hash([]byte("The voter senate session，one and only one"))
	SenateProposal          = senateSessionID.Big()
	SenateSessionIDOfECDSA  = ethCrypto.Keccak256Hash([]byte("ECDSA:The voter senate session，one and only one"))
	SenateProposalIDOfECDSA = SenateSessionIDOfECDSA.String()
	SenateSessionIDOfEDDSA  = ethCrypto.Keccak256Hash([]byte("EDDSA:The voter senate session，one and only one"))
	SenateProposalIDOfEDDSA = SenateSessionIDOfEDDSA.String()
)

type SessionContext[T, M any] struct {
	Group
	SeqId      uint64          `json:"seq_id"`
	SessionID  party.SessionID `json:"sessionID,omitempty"`
	Proposer   common.Address  `json:"proposer,omitempty"`    // current submitter
	Signer     string          `json:"signer,omitempty"`      // current signer
	ProposalID T               `json:"proposal_id,omitempty"` // msg id
	Proposal   M               `json:"proposal,omitempty"`
	Data       []byte          `json:"data,omitempty"`
}

type Group struct {
	EC          crypto.CurveType   `json:"ec,omitempty"`
	ChainId     uint64             `json:"chain_id,omitempty"`
	AllPartners types.Participants `json:"all_partners,omitempty"` // all submitter
}

func (g *Group) GroupID() party.GroupID {
	return g.AllPartners.GroupID()
}

type LocalPartySaveData struct {
	ty   crypto.CurveType // 0:secp256k1; 1:ed25519
	data []byte           // json byte
}

func BuildECDSALocalPartySaveData() *LocalPartySaveData {
	return &LocalPartySaveData{ty: crypto.ECDSA}
}

func BuildEDDSALocalPartySaveData() *LocalPartySaveData {
	return &LocalPartySaveData{ty: crypto.EDDSA}
}

func (d *LocalPartySaveData) SetData(data any) *LocalPartySaveData {
	raw, err := json.Marshal(data)
	utils.Assert(err)

	d.data = raw

	return d
}

func (d *LocalPartySaveData) GetJsonData() []byte {
	return d.data
}

func (d *LocalPartySaveData) EC() elliptic.Curve {
	switch d.ty {
	case crypto.EDDSA:
		return tss.Edwards()
	case crypto.ECDSA:
		return tss.S256()
	default:
		panic("unknown elliptic curve")
	}
}

func (d *LocalPartySaveData) CurveType() crypto.CurveType {
	return d.ty
}

func (d *LocalPartySaveData) ECDSAData() *ecdsaKeygen.LocalPartySaveData {
	if d.ty == crypto.ECDSA && len(d.data) != 0 {
		data := &ecdsaKeygen.LocalPartySaveData{}
		err := json.Unmarshal(d.data, data)
		utils.Assert(err)

		return data
	}

	return nil
}

func (d *LocalPartySaveData) EDDSAData() *eddsaKeygen.LocalPartySaveData {
	if d.ty == crypto.EDDSA && len(d.data) != 0 {
		data := &eddsaKeygen.LocalPartySaveData{}
		err := json.Unmarshal(d.data, data)
		utils.Assert(err)

		return data
	}

	return nil
}

func (d *LocalPartySaveData) TssSigner() string {
	switch d.ty {
	case crypto.ECDSA:
		return ethCrypto.PubkeyToAddress(*d.ECDSAData().ECDSAPub.ToECDSAPubKey()).String()
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) ToECDSAPubKey() *ecdsa.PublicKey {
	switch d.ty {
	case crypto.ECDSA:
		return d.ECDSAData().ECDSAPub.ToECDSAPubKey()
	case crypto.EDDSA:
		return d.EDDSAData().EDDSAPub.ToECDSAPubKey()
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) ECPoint() *tssCrypto.ECPoint {
	switch d.ty {
	case crypto.ECDSA:
		return d.ECDSAData().ECDSAPub
	case crypto.EDDSA:
		return d.EDDSAData().EDDSAPub
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) PublicKeyBase58() string {
	switch d.ty {
	case crypto.ECDSA:
		pubKey := d.ECDSAData().ECDSAPub.ToECDSAPubKey()
		pubKeyBytes := ethCrypto.FromECDSAPub(pubKey)

		return base58.Encode(pubKeyBytes)
	case crypto.EDDSA:
		pubKey := d.EDDSAData().EDDSAPub.ToECDSAPubKey()
		pubKeyBytes := ethCrypto.FromECDSAPub(pubKey)

		return base58.Encode(pubKeyBytes)
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) CompressedPublicKey() string {
	return hex.EncodeToString(d.PublicKey().SerializeCompressed())
}

func (d *LocalPartySaveData) PublicKey() crypto.PublicKey {
	p := d.ECPoint()

	switch d.ty {
	case crypto.ECDSA:
		var (
			x = &btcec.FieldVal{}
			y = &btcec.FieldVal{}
		)

		x.SetByteSlice(p.X().Bytes())
		y.SetByteSlice(p.Y().Bytes())

		return secp256k1.NewPublicKey(x, y)
	case crypto.EDDSA:
		return edwards.NewPublicKey(p.X(), p.Y())
	default:
		panic("implement me")
	}
}

func (d *LocalPartySaveData) Address(chainType uint8) string {
	return address.GenerateAddressByECPoint(d.ECPoint(), types.GetCoinTypeByChain(chainType))
}
