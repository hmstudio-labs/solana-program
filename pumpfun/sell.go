package pumpfun

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

// Sells tokens into a bonding curve.
type Sell struct {
	Amount       *uint64
	MinSolOutput *uint64

	// [0] = [] global
	//
	// [1] = [WRITE] feeRecipient
	//
	// [2] = [] mint
	//
	// [3] = [WRITE] bondingCurve
	//
	// [4] = [WRITE] associatedBondingCurve
	//
	// [5] = [WRITE] associatedUser
	//
	// [6] = [WRITE, SIGNER] user
	//
	// [7] = [] systemProgram
	//
	// [8] = [] associatedTokenProgram
	//
	// [9] = [] tokenProgram
	//
	// [10] = [] eventAuthority
	//
	// [11] = [] program
	sol.AccountMetaSlice `bin:"-"`
}

// NewSellInstruction declares a new Sell instruction with the provided parameters and accounts.
func NewSellInstruction(
	// Parameters:
	amount uint64,
	minSolOutput uint64,
	// Accounts:
	mint sol.PublicKey,
	bondingCurve sol.PublicKey,
	associatedBondingCurve sol.PublicKey,
	associatedUser sol.PublicKey,
	user sol.PublicKey) *Sell {
	return newSelInstructionBuilder().
		setAmount(amount).
		SetMinSolOutput(minSolOutput).
		setMintAccount(mint).
		setBondingCurveAccount(bondingCurve).
		setAssociatedBondingCurveAccount(associatedBondingCurve).
		setAssociatedUserAccount(associatedUser).
		setUserAccount(user)
}

// NewSellInstructionBuilder creates a new `Sell` instruction builder.
func newSelInstructionBuilder() *Sell {
	nd := &Sell{
		AccountMetaSlice: make(sol.AccountMetaSlice, 12),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *Sell) setAmount(amount uint64) *Sell {
	inst.Amount = &amount
	return inst
}

// SetMinSolOutput sets the "minSolOutput" parameter.
func (inst *Sell) SetMinSolOutput(minSolOutput uint64) *Sell {
	inst.MinSolOutput = &minSolOutput
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *Sell) setMintAccount(mint sol.PublicKey) *Sell {
	inst.AccountMetaSlice[2] = sol.Meta(mint)
	return inst
}

// SetBondingCurveAccount sets the "bondingCurve" account.
func (inst *Sell) setBondingCurveAccount(bondingCurve sol.PublicKey) *Sell {
	inst.AccountMetaSlice[3] = sol.Meta(bondingCurve).WRITE()
	return inst
}

// SetAssociatedBondingCurveAccount sets the "associatedBondingCurve" account.
func (inst *Sell) setAssociatedBondingCurveAccount(associatedBondingCurve sol.PublicKey) *Sell {
	inst.AccountMetaSlice[4] = sol.Meta(associatedBondingCurve).WRITE()
	return inst
}

// SetAssociatedUserAccount sets the "associatedUser" account.
func (inst *Sell) setAssociatedUserAccount(associatedUser sol.PublicKey) *Sell {
	inst.AccountMetaSlice[5] = sol.Meta(associatedUser).WRITE()
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *Sell) setUserAccount(user sol.PublicKey) *Sell {
	inst.AccountMetaSlice[6] = sol.Meta(user).WRITE().SIGNER()
	return inst
}

func (inst *Sell) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[0] = sol.Meta(Global)
	inst.AccountMetaSlice[1] = sol.Meta(FeeRecipient).WRITE()
	inst.AccountMetaSlice[7] = sol.Meta(SystemProgramId)
	inst.AccountMetaSlice[8] = sol.Meta(AssociatedTokenProgramId)
	inst.AccountMetaSlice[9] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[10] = sol.Meta(EventAuthority)
	inst.AccountMetaSlice[11] = sol.Meta(PumpProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Sell,
	}}
}

func (inst Sell) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(inst.Amount)
	if err != nil {
		return err
	}
	// Serialize `MaxSolCost` param:
	err = encoder.Encode(inst.MinSolOutput)
	if err != nil {
		return err
	}
	return nil
}
