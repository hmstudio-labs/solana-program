package lookuptableaccount

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type FreezeLookupTable struct {
	// [0] = [WRITE] lookupTableAccount
	//
	// [1] = [WRITE, SIGNER] authority
	sol.AccountMetaSlice `bin:"-"`
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewFreezeLookupTableInstruction(
	// Accounts:
	lookupTableAccount sol.PublicKey,
	authority sol.PublicKey) *FreezeLookupTable {
	return newFreezeLookupTableInstructionBuilder().
		SetLookupTableAccount(lookupTableAccount).
		SetLookupTableAuthority(authority)
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func newFreezeLookupTableInstructionBuilder() *FreezeLookupTable {
	nd := &FreezeLookupTable{
		AccountMetaSlice: make(sol.AccountMetaSlice, 2),
	}
	return nd
}

func (inst *FreezeLookupTable) SetLookupTableAccount(lookupTableAccount sol.PublicKey) *FreezeLookupTable {
	inst.AccountMetaSlice[0] = sol.Meta(lookupTableAccount).WRITE()
	return inst
}
func (inst *FreezeLookupTable) SetLookupTableAuthority(authority sol.PublicKey) *FreezeLookupTable {
	inst.AccountMetaSlice[1] = sol.Meta(authority).WRITE().SIGNER()
	return inst
}

func (inst *FreezeLookupTable) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[3] = sol.Meta(SystemProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FreezeLookupTable,
	}}
}

func (inst FreezeLookupTable) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	return nil
}
