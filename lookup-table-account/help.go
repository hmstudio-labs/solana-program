package lookuptableaccount

import (
	"bytes"
	"encoding/binary"
	"math/big"

	sol "github.com/gagliardetto/solana-go"
)

func CreateLookupTableQuick(
	recentSlot uint64,
	authority sol.PublicKey,
	payer sol.PublicKey) (sol.PublicKey, *Instruction) {
	recentSlotBytes := encodeI64LE(big.NewInt(int64(recentSlot)))
	lookupTableAddress, bumpSeed, _ := sol.FindProgramAddress([][]byte{
		authority[:],
		recentSlotBytes,
	}, ProgramID)

	ix := NewInstruction(recentSlot, bumpSeed, lookupTableAddress, authority, payer).Build()
	return lookupTableAddress, ix
}

func ExtendLookupTableQuick(
	lookupTable sol.PublicKey,
	authority sol.PublicKey,
	payer sol.PublicKey,
	addresses []sol.PublicKey) *Instruction {
	return NewExtendLookupTableInstruction(addresses, lookupTable, authority, payer).Build()
}

func encodeI64LE(n *big.Int) []byte {
	// 将 n 限定在 int64 范围
	i64 := n.Int64() // 会截断超出范围
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, i64)
	return buf.Bytes()
}
