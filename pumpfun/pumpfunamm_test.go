package pumpfun

import (
	"encoding/hex"
	"testing"

	"github.com/test-go/testify/require"
)

func TestPumpfunBuyData(t *testing.T) {
	unknow := SystemProgramId
	result := "66063d1201daebea858c2a7821010000f0d8700200000000"
	var amount uint64 = 1243261602949
	var maxSolCost uint64 = 40950000
	buy := NewBuyInstruction(
		amount, maxSolCost, unknow, unknow, unknow, unknow, unknow,
	)
	i := buy.Build()
	data, err := i.Data()

	// fmt.Println(hex.EncodeToString(data), err)
	require.NoError(t, err)
	require.Equal(t, result, hex.EncodeToString(data))
}

func TestPumpfunSellData(t *testing.T) {
	unknow := SystemProgramId
	result := "33e685a4017f83ad00c0afd6913600000000000000000000"
	var amount uint64 = 60000000000000
	var minSolOutput uint64 = 0
	buy := NewSellInstruction(
		amount, minSolOutput, unknow, unknow, unknow, unknow, unknow,
	)
	i := buy.Build()
	data, err := i.Data()

	// fmt.Println(hex.EncodeToString(data), err)
	require.NoError(t, err)
	require.Equal(t, result, hex.EncodeToString(data))
}
