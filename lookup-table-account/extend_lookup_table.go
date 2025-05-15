package lookuptableaccount

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type ExtendLookupTable struct {
	Addresses []sol.PublicKey

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
func NewExtendLookupTableInstruction(
	addresses []sol.PublicKey,
	// Accounts:
	lookupTableAccount sol.PublicKey,
	authority sol.PublicKey,
	payer sol.PublicKey) *ExtendLookupTable {
	return newExtendLookupTableInstructionBuilder().
		SetLookupTableAccount(lookupTableAccount).
		SetLookupTableAuthority(authority).
		SetPayer(payer)
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func newExtendLookupTableInstructionBuilder() *ExtendLookupTable {
	nd := &ExtendLookupTable{
		AccountMetaSlice: make(sol.AccountMetaSlice, 4),
	}
	return nd
}

func (inst *ExtendLookupTable) SetAddresses(addresses []sol.PublicKey) *ExtendLookupTable {
	inst.Addresses = addresses
	return inst
}
func (inst *ExtendLookupTable) SetLookupTableAccount(lookupTableAccount sol.PublicKey) *ExtendLookupTable {
	inst.AccountMetaSlice[0] = sol.Meta(lookupTableAccount).WRITE()
	return inst
}
func (inst *ExtendLookupTable) SetLookupTableAuthority(authority sol.PublicKey) *ExtendLookupTable {
	inst.AccountMetaSlice[1] = sol.Meta(authority).WRITE().SIGNER()
	return inst
}

func (inst *ExtendLookupTable) SetPayer(payer sol.PublicKey) *ExtendLookupTable {
	inst.AccountMetaSlice[2] = sol.Meta(payer).WRITE().SIGNER()
	return inst
}

func (inst *ExtendLookupTable) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[3] = sol.Meta(SystemProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ExtendLookupTable,
	}}
}

func (inst ExtendLookupTable) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	addressesLen := uint64(len(inst.Addresses))
	err = encoder.Encode(addressesLen)
	if err != nil {
		return err
	}
	for _, address := range inst.Addresses {
		err = encoder.Encode(address)
		if err != nil {
			return err
		}
	}
	return nil
}
