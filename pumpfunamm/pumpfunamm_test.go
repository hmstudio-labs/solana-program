package pumpfunamm

import (
	"encoding/hex"
	"testing"

	"github.com/test-go/testify/require"
)

func TestPumpfunammBuyData(t *testing.T) {
	unknow := SystemProgramId
	result := "66063d1201daebea9e5fb0000000000040420f0000000000"
	var amount uint64 = 11558814
	var maxSolCost uint64 = 1000000
	buy := NewBuyInstruction(
		amount, maxSolCost, unknow, unknow, unknow, unknow, unknow, unknow, unknow,
	)
	i := buy.Build()
	data, err := i.Data()

	// fmt.Println(hex.EncodeToString(data), err)
	require.NoError(t, err)
	require.Equal(t, result, hex.EncodeToString(data))
}

func TestPumpfunSellData(t *testing.T) {
	unknow := SystemProgramId
	result := "33e685a4017f83ade0abf351010000000000000000000000"
	var amount uint64 = 5669891040
	var minSolOutput uint64 = 0
	buy := NewSellInstruction(
		amount, minSolOutput, unknow, unknow, unknow, unknow, unknow, unknow, unknow,
	)
	i := buy.Build()
	data, err := i.Data()

	// fmt.Println(hex.EncodeToString(data), err)
	require.NoError(t, err)
	require.Equal(t, result, hex.EncodeToString(data))
}
