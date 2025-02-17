package utils

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeString(t *testing.T) {
	data := EncodeString("0x1ED76b7CDf23A597d97c5734104B7fEbE32EB17e")
	dataStr := hex.EncodeToString(data)
	t.Logf("data: %x", data)

	expectData := "000000000000000000000000000000000000000000000000000000000000002a30783145443736623743446632334135393764393763353733343130344237664562453332454231376500000000000000000000000000000000000000000000"
	assert.Equal(t, expectData, dataStr)
}
