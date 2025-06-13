package lookuptableaccount

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type DeactivateLookupTable struct {
	// [0] = [WRITE] lookupTableAccount
	//
	// [1] = [WRITE, SIGNER] authority
	sol.AccountMetaSlice `bin:"-"`
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewDeactivateLookupTableInstruction(
	// Accounts:
	lookupTableAccount sol.PublicKey,
	authority sol.PublicKey) *DeactivateLookupTable {
	return newDeactivateLookupTableInstructionBuilder().
		SetLookupTableAccount(lookupTableAccount).
		SetLookupTableAuthority(authority)
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func newDeactivateLookupTableInstructionBuilder() *DeactivateLookupTable {
	nd := &DeactivateLookupTable{
		AccountMetaSlice: make(sol.AccountMetaSlice, 2),
	}
	return nd
}

func (inst *DeactivateLookupTable) SetLookupTableAccount(lookupTableAccount sol.PublicKey) *DeactivateLookupTable {
	inst.AccountMetaSlice[0] = sol.Meta(lookupTableAccount).WRITE()
	return inst
}
func (inst *DeactivateLookupTable) SetLookupTableAuthority(authority sol.PublicKey) *DeactivateLookupTable {
	inst.AccountMetaSlice[1] = sol.Meta(authority).WRITE().SIGNER()
	return inst
}

func (inst *DeactivateLookupTable) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[3] = sol.Meta(SystemProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DeactivateLookupTable,
	}}
}

func (inst DeactivateLookupTable) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	return nil
}
