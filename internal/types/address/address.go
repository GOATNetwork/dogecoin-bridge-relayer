package address

import (
	"math/big"

	tssCrypto "github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/crypto/ckd"
	"github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/nuvosphere/nudex-voter/internal/wallet/bip44"
)

func init() {
	addressRegister = make(map[int]func(*tssCrypto.ECPoint) string)
}

var addressRegister map[int]func(*tssCrypto.ECPoint) string

func RegisterAddress(coinType int, f func(*tssCrypto.ECPoint) string) {
	addressRegister[coinType] = f
}

func GenerateAddressByECPoint(point *tssCrypto.ECPoint, coinType int) string {
	m, ok := addressRegister[coinType]
	if ok {
		return m(point)
	}

	panic("invalid coin type")
}

func GenerateAddressByPath(masterPubKey *tssCrypto.ECPoint, coinType, account, index uint32) string {
	address, _, _ := GenerateSignerByPath(masterPubKey, coinType, account, index)
	return address
}

func GenerateSignerByPath(masterPubKey *tssCrypto.ECPoint, coinType, account, index uint32) (string, *tssCrypto.ECPoint, *big.Int) {
	bip44Path := bip44.Bip44DerivationPath(coinType, account, index)

	p, err := bip44Path.ToParams()
	utils.Assert(err)

	indexes := p.Indexes()

	chainCode := big.NewInt(int64(coinType)).Bytes() // todo
	curveType := types.GetCurveTypeByCoinType(int(coinType))
	derivingKey, extendKey, err := ckd.DerivingPubkeyFromPath(masterPubKey, chainCode, indexes, curveType.EC())
	utils.Assert(err)

	return GenerateAddressByECPoint(extendKey.PublicKey, int(coinType)), extendKey.PublicKey, derivingKey
}

func HotAddressOfBtc(masterPubKey *tssCrypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeBTC, 0, 0)
}

func HotAddressOfEth(masterPubKey *tssCrypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeEVM, 0, 0)
}

func HotAddressOfSolana(masterPubKey *tssCrypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeSOL, 0, 0)
}

func HotAddressOfSui(masterPubKey *tssCrypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeSUI, 0, 0)
}

func HotAddressOfDog(masterPubKey *tssCrypto.ECPoint) string {
	return GenerateAddressByPath(masterPubKey, types.CoinTypeDOG, 0, 0)
}
