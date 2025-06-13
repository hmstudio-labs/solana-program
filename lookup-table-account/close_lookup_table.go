package lookuptableaccount

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type CloseLookupTable struct {
	// [0] = [WRITE] lookupTableAccount
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [WRITE, SIGNER] payer
	sol.AccountMetaSlice `bin:"-"`
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewCloseLookupTableInstruction(
	// Accounts:
	lookupTableAccount sol.PublicKey,
	authority sol.PublicKey,
	payer sol.PublicKey) *CloseLookupTable {
	return newCloseLookupTableInstructionBuilder().
		SetLookupTableAccount(lookupTableAccount).
		SetLookupTableAuthority(authority).
		SetPayer(payer)
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func newCloseLookupTableInstructionBuilder() *CloseLookupTable {
	nd := &CloseLookupTable{
		AccountMetaSlice: make(sol.AccountMetaSlice, 3),
	}
	return nd
}

func (inst *CloseLookupTable) SetLookupTableAccount(lookupTableAccount sol.PublicKey) *CloseLookupTable {
	inst.AccountMetaSlice[0] = sol.Meta(lookupTableAccount).WRITE()
	return inst
}
func (inst *CloseLookupTable) SetLookupTableAuthority(authority sol.PublicKey) *CloseLookupTable {
	inst.AccountMetaSlice[1] = sol.Meta(authority).WRITE().SIGNER()
	return inst
}

func (inst *CloseLookupTable) SetPayer(payer sol.PublicKey) *CloseLookupTable {
	inst.AccountMetaSlice[2] = sol.Meta(payer).WRITE().SIGNER()
	return inst
}

func (inst *CloseLookupTable) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[3] = sol.Meta(SystemProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CloseLookupTable,
	}}
}

func (inst CloseLookupTable) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	return nil
}
