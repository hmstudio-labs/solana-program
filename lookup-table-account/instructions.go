package lookuptableaccount

import (
	"bytes"
	"fmt"

	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

var (
	SystemProgramId = sol.MustPublicKeyFromBase58("11111111111111111111111111111111")
	ProgramID       = sol.MustPublicKeyFromBase58("AddressLookupTab1e1111111111111111111111111")
)

var (
	Instruction_CreateLookupTable     = bin.TypeID([8]byte{0, 0, 0, 0, 0, 0, 0, 0})
	Instruction_FreezeLookupTable     = bin.TypeID([8]byte{1, 0, 0, 0, 0, 0, 0, 0})
	Instruction_ExtendLookupTable     = bin.TypeID([8]byte{2, 0, 0, 0, 0, 0, 0, 0})
	Instruction_DeactivateLookupTable = bin.TypeID([8]byte{3, 0, 0, 0, 0, 0, 0, 0})
	Instruction_CloseLookupTable      = bin.TypeID([8]byte{4, 0, 0, 0, 0, 0, 0, 0})
)

type Instruction struct {
	bin.BaseVariant
}

func (inst *Instruction) ProgramID() sol.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*sol.AccountMeta) {
	return inst.Impl.(sol.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bin.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst Instruction) MarshalWithEncoder(encoder *bin.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes()[0:4], false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}
