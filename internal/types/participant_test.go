package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestParticipantsThreshold(t *testing.T) {
	assert.Equal(t, CalculateThreshold(3), 1)   // (1 + 1) / 3 = 2/3
	assert.Equal(t, CalculateThreshold(4), 2)   // (2 + 1) / 4 = 3/4 > 2/3
	assert.Equal(t, CalculateThreshold(5), 3)   // (3 + 1) / 5 = 4/5 > 2/3
	assert.Equal(t, CalculateThreshold(6), 3)   // (3 + 1) / 6 = 4/6 = 2/3
	assert.Equal(t, CalculateThreshold(7), 4)   // (4 + 1) / 7 = 5/7 > 2/3
	assert.Equal(t, CalculateThreshold(8), 5)   // (5 + 1) / 8 = 6/8 > 2/3
	assert.Equal(t, CalculateThreshold(9), 5)   // (5 + 1) / 9 = 6/9 = 2/3
	assert.Equal(t, CalculateThreshold(10), 6)  // (6 + 1) / 10 = 7/10 > 2/3
	assert.Equal(t, CalculateThreshold(11), 7)  // (7 + 1) / 11 = 8/11 > 2/3
	assert.Equal(t, CalculateThreshold(12), 7)  // (7 + 1) / 12 = 8/12 = 2/3
	assert.Equal(t, CalculateThreshold(13), 8)  // (8 + 1) / 13 = 9/13 > 2/3
	assert.Equal(t, CalculateThreshold(15), 9)  // (9 + 1) / 15 = 10/15 = 2/3
	assert.Equal(t, CalculateThreshold(20), 13) // (13 + 1) / 20 = 14/20 = 7/10 > 2/3
}

func TestGroupID(t *testing.T) {
	group := Participants{
		common.HexToAddress("0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4"),
		common.HexToAddress("0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"),
		common.HexToAddress("0xf6D37CE75dB465DcDb4c7097bEB9c1D46b171037"),
		common.HexToAddress("0x1D2cd50A3cF3c55a7982AD54F9f364C1e953Bc57"),
		common.HexToAddress("0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2"),
	}
	one := group.GroupID()
	t.Log(one)

	group = Participants{
		common.HexToAddress("0x1D2cd50A3cF3c55a7982AD54F9f364C1e953Bc57"),
		common.HexToAddress("0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4"),
		common.HexToAddress("0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2"),
		common.HexToAddress("0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"),
		common.HexToAddress("0xf6D37CE75dB465DcDb4c7097bEB9c1D46b171037"),
	}
	two := group.GroupID()
	t.Log(two)
	assert.Equal(t, two, one)

	group = Participants{
		common.HexToAddress("0x1D2cd50A3cF3c55a7982AD54F9f364C1e953Bc57"),
		common.HexToAddress("0x3a818294ca1F3C27d7588b123Ec43F2546fa07f4"),
		common.HexToAddress("0x5091FC3cb4E4BB014141Aa41375d8Dd73b34AfA2"),
		common.HexToAddress("0x04d9389Cf937b1e6F2258d842e7237E955d6ab04"),
	}
	three := group.GroupID()
	t.Log(three)
	assert.NotEqual(t, three, one)
}
