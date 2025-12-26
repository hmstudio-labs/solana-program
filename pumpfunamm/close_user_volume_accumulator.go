package pumpfunamm

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type CloseUserVolumeAccumulator struct {

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [WRITE] user volume accumulator
	//
	// [2] = [] event authority
	//
	// [3] = [] program
	sol.AccountMetaSlice `bin:"-"`
}

// NewCloseUserVolumeAccumulatorInstruction declares a new CloseUserVolumeAccumulator instruction with the provided parameters and accounts.
func NewCloseUserVolumeAccumulatorInstruction(
	// Accounts:
	userVolumeAccumulator sol.PublicKey,
	user sol.PublicKey) *CloseUserVolumeAccumulator {
	return newCloseUserVolumeAccumulatorInstructionBuilder().
		setUserVolumeAccumulator(userVolumeAccumulator).
		setUserAccount(user)
}

// NewCloseUserVolumeAccumulatorInstructionBuilder creates a new `CloseUserVolumeAccumulator` instruction builder.
func newCloseUserVolumeAccumulatorInstructionBuilder() *CloseUserVolumeAccumulator {
	nd := &CloseUserVolumeAccumulator{
		AccountMetaSlice: make(sol.AccountMetaSlice, 4),
	}
	return nd
}

// SetUserAccount sets the "user" account.
func (inst *CloseUserVolumeAccumulator) setUserAccount(user sol.PublicKey) *CloseUserVolumeAccumulator {
	inst.AccountMetaSlice[0] = sol.Meta(user).WRITE().SIGNER()
	return inst
}

func (inst *CloseUserVolumeAccumulator) setUserVolumeAccumulator(userVolumeAccumulator sol.PublicKey) *CloseUserVolumeAccumulator {
	inst.AccountMetaSlice[1] = sol.Meta(userVolumeAccumulator).WRITE()
	return inst
}
func (inst *CloseUserVolumeAccumulator) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[2] = sol.Meta(EventAuthority)
	inst.AccountMetaSlice[3] = sol.Meta(PumpAMMProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CloseUserVolumeAccumulator,
	}}
}

func (inst CloseUserVolumeAccumulator) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	return nil
}
