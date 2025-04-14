package pumpfun

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type Buy struct {
	Amount     *uint64
	MaxSolCost *uint64

	// [0] = [] pool
	//
	// [1] = [WRITE, SIGNER] user
	//
	// [2] = [] global
	//
	// [3] = [] mint
	//
	// [4] = [] wsol
	//
	// [5] = [WRITE] useBaseTokenAccount
	//
	// [6] = [WRITE] user
	//
	// [7] = [] systemProgram
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] rent
	//
	// [10] = [] eventAuthority
	//
	// [11] = [] program
	sol.AccountMetaSlice `bin:"-"`
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewBuyInstruction(
	// Parameters:
	amount uint64,
	maxSolCost uint64,
	// Accounts:
	mint sol.PublicKey,
	bondingCurve sol.PublicKey,
	associatedBondingCurve sol.PublicKey,
	associatedUser sol.PublicKey,
	user sol.PublicKey) *Buy {
	return newBuyInstructionBuilder().
		setAmount(amount).
		setMaxSolCost(maxSolCost).
		setMintAccount(mint).
		setBondingCurveAccount(bondingCurve).
		setAssociatedBondingCurveAccount(associatedBondingCurve).
		setAssociatedUserAccount(associatedUser).
		setUserAccount(user)
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func newBuyInstructionBuilder() *Buy {
	nd := &Buy{
		AccountMetaSlice: make(sol.AccountMetaSlice, 12),
	}
	return nd
}

// setAmount sets the "amount" parameter.
func (inst *Buy) setAmount(amount uint64) *Buy {
	inst.Amount = &amount
	return inst
}

// setMaxSolCost sets the "maxSolCost" parameter.
func (inst *Buy) setMaxSolCost(maxSolCost uint64) *Buy {
	inst.MaxSolCost = &maxSolCost
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *Buy) setMintAccount(mint sol.PublicKey) *Buy {
	inst.AccountMetaSlice[2] = sol.Meta(mint)
	return inst
}

// SetBondingCurveAccount sets the "bondingCurve" account.
func (inst *Buy) setBondingCurveAccount(bondingCurve sol.PublicKey) *Buy {
	inst.AccountMetaSlice[3] = sol.Meta(bondingCurve).WRITE()
	return inst
}

// SetAssociatedBondingCurveAccount sets the "associatedBondingCurve" account.
func (inst *Buy) setAssociatedBondingCurveAccount(associatedBondingCurve sol.PublicKey) *Buy {
	inst.AccountMetaSlice[4] = sol.Meta(associatedBondingCurve).WRITE()
	return inst
}

// SetAssociatedUserAccount sets the "associatedUser" account.
func (inst *Buy) setAssociatedUserAccount(associatedUser sol.PublicKey) *Buy {
	inst.AccountMetaSlice[5] = sol.Meta(associatedUser).WRITE()
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *Buy) setUserAccount(user sol.PublicKey) *Buy {
	inst.AccountMetaSlice[6] = sol.Meta(user).WRITE().SIGNER()
	return inst
}

func (inst *Buy) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[0] = sol.Meta(Global)
	inst.AccountMetaSlice[1] = sol.Meta(FeeRecipient).WRITE()
	inst.AccountMetaSlice[7] = sol.Meta(SystemProgramId)
	inst.AccountMetaSlice[8] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[9] = sol.Meta(RentProgramId)
	inst.AccountMetaSlice[10] = sol.Meta(EventAuthority)
	inst.AccountMetaSlice[11] = sol.Meta(PumpProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Buy,
	}}
}

func (inst Buy) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(inst.Amount)
	if err != nil {
		return err
	}
	// Serialize `MaxSolCost` param:
	err = encoder.Encode(inst.MaxSolCost)
	if err != nil {
		return err
	}
	return nil
}
