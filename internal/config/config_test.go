package config

import (
	"os"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test.yaml")
	require.NoError(t, err)
	defer func(name string) {
		err := os.Remove(name)
		assert.NoError(t, err)

		err = tmpfile.Close()
		assert.NoError(t, err)
	}(tmpfile.Name())

	testData := `
env: dev # dev、test、prod
master_chain_id: 12 # contract on chain id
submitter_private_key: '76cbb08e5321cec5f584b2b40b4666d9bbbee59eb3022e80d804e8310b17a105'

db:
  db_root_dir: 'data/db'

log:
  level: debug

tss:
  threshold: 1
  sign_timeout: 60 # 60s
  public_keys:
    - '020b537f46c6da81f84824ce1409bab1f9825fb58b57dcafbf4f4b074e90a0c040'
    - '02a8fd23c439e9226f422e94911f06788e0019aa1f8efd4f498f75e4f1d5ef7c0a'
    - '02f82403b0337c908478d381f88582e1051c2a9da22a34cd0a1a5b1d10a85b6256'

p2p:
  port: 4001
  boot_nodes:
    - '/ip4/127.0.0.1/tcp/4001/p2p/16Uiu2HAkvBtJw6RmmPfF4nqFZNzK7cGgscmZ7AdP7Fjji3MuKvdM'
  
contract:
  voter: '0x49fd2BE640DB2910c2fAb69bB8531Ab6E76127ff'
  account: '0xe8D2A1E88c91DCd5433208d4152Cc4F399a7e91d'
  task_manager: ''
  participant: ''
  deposit: ''
  asset_handler: ''
  
chains:
- chain_id: 1
  network: testnet
  chain_type: 0
  start_height: 1
  confirmations: 1
  max_block_range: 2
  scan_interval: 10 # 10s
  rpc: 
    url: 'http://localhost:8545'
    user: 'test'
    password: 'test'
    jwt_secret: 'test'

- chain_id: 2
  network: testnet
  chain_type: 0
  start_height: 1
  confirmations: 1
  max_block_range: 2
  scan_interval: 10 # 10s
  rpc:
    url: 'http://localhost:8545'
    user: ''
    password: ''
    jwt_secret: ''
 
- chain_id: 12 # master chain
  network: testnet
  chain_type: 0
  start_height: 1
  confirmations: 1
  max_block_range: 2
  scan_interval: 10 # 10s
  rpc:
    url: 'http://localhost:8545'
    user: ''
    password: ''
    jwt_secret: ''
`
	filePath := tmpfile.Name()
	t.Log(filePath)

	data := []byte(testData)
	err = os.WriteFile(tmpfile.Name(), data, 0o600)
	require.NoError(t, err)

	InitConfig(tmpfile.Name())
	t.Log(utils.FormatJSON(AppConfig))
	assert.Equal(t, uint64(12), AppConfig.MasterChainId)
	assert.Equal(t, 3, len(AppConfig.Chains))
	assert.Equal(t, time.Duration(60), AppConfig.Tss.SignTimeout)
	assert.Equal(t, "76cbb08e5321cec5f584b2b40b4666d9bbbee59eb3022e80d804e8310b17a105", AppConfig.SubmitterPrivateKey)
	AppConfig.MasterChainInfo()
	assert.Equal(t, "dev", AppConfig.Env)
	assert.Equal(t, "data/db", AppConfig.DB.DbRootDir)

	err = validator.New().Struct(AppConfig)
	assert.NoError(t, err)
}
