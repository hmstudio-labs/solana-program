package lookuptableaccount

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type CreateLookupTable struct {
	RecentSlot *uint64
	BumpSeed   *uint8

	// [0] = [WRITE] lookupTableAccount
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [] system program
	sol.AccountMetaSlice `bin:"-"`
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewInstruction(
	// Parameters:
	recentSlot uint64,
	bumpSeed uint8,
	// Accounts:
	lookupTableAccount sol.PublicKey,
	authority sol.PublicKey,
	payer sol.PublicKey) *CreateLookupTable {
	return newCreateLookupTableInstructionBuilder().
		SetRecentSlot(recentSlot).
		SetBumpSeed(bumpSeed).
		SetLookupTableAccount(lookupTableAccount).
		SetLookupTableAuthority(authority).
		SetPayer(payer)
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func newCreateLookupTableInstructionBuilder() *CreateLookupTable {
	nd := &CreateLookupTable{
		AccountMetaSlice: make(sol.AccountMetaSlice, 4),
	}
	return nd
}

func (inst *CreateLookupTable) SetRecentSlot(recentSlot uint64) *CreateLookupTable {
	inst.RecentSlot = &recentSlot
	return inst
}
func (inst *CreateLookupTable) SetBumpSeed(bumpSeed uint8) *CreateLookupTable {
	inst.BumpSeed = &bumpSeed
	return inst
}
func (inst *CreateLookupTable) SetLookupTableAccount(lookupTableAccount sol.PublicKey) *CreateLookupTable {
	inst.AccountMetaSlice[0] = sol.Meta(lookupTableAccount).WRITE()
	return inst
}
func (inst *CreateLookupTable) SetLookupTableAuthority(authority sol.PublicKey) *CreateLookupTable {
	inst.AccountMetaSlice[1] = sol.Meta(authority).WRITE().SIGNER()
	return inst
}

func (inst *CreateLookupTable) SetPayer(payer sol.PublicKey) *CreateLookupTable {
	inst.AccountMetaSlice[2] = sol.Meta(payer).WRITE().SIGNER()
	return inst
}

func (inst *CreateLookupTable) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[3] = sol.Meta(SystemProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateLookupTable,
	}}
}

func (inst CreateLookupTable) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	// Serialize `RecentSlot` param:
	err = encoder.Encode(inst.RecentSlot)
	if err != nil {
		return err
	}
	// Serialize `BumpSeed` param:
	err = encoder.Encode(inst.BumpSeed)
	if err != nil {
		return err
	}
	return nil
}
