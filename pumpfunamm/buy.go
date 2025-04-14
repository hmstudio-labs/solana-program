package pumpfunamm

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type Buy struct {
	BaseAmountOut    *uint64
	MaxQuoteAmountIn *uint64

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
	// [6] = [WRITE] useQuoteTokenAccount
	//
	// [7] = [WRITE] poolBaseTokenAccount
	//
	// [8] = [WRITE] poolQuoteTokenAccount
	//
	// [9] = [] feeRecipient
	//
	// [10] = [WRITE] feeRecipientTokenAccount
	//
	// [11] = [] base token program
	//
	// [12] = [] quote token program
	//
	// [13] = [] system program
	//
	// [14] = [] associated token program
	//
	// [15] = [] event authority
	//
	// [16] = [] program
	sol.AccountMetaSlice `bin:"-"`
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewBuyInstruction(
	// Parameters:
	baseAmountOut uint64,
	maxQuoteAmountIn uint64,
	// Accounts:
	pool sol.PublicKey,
	mint sol.PublicKey,
	useBaseTokenAccount sol.PublicKey,
	userQuoteTokenAccount sol.PublicKey,
	poolBaseTokenAccount sol.PublicKey,
	poolQuoteTokenAccount sol.PublicKey,
	associatedBondingCurve sol.PublicKey,
	associatedUser sol.PublicKey,
	user sol.PublicKey) *Buy {
	return newBuyInstructionBuilder().
		setBaseAmountOut(baseAmountOut).
		setMaxQuoteAmountIn(maxQuoteAmountIn).
		setPoolAccount(pool).
		setMintAccount(mint).
		setUserBaseTokenAccount(useBaseTokenAccount).
		setUserQuoteTokenAccount(userQuoteTokenAccount).
		setPoolBaseTokenAccount(poolBaseTokenAccount).
		setPoolQuoteTokenAccount(poolQuoteTokenAccount).
		setUserAccount(user)
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func newBuyInstructionBuilder() *Buy {
	nd := &Buy{
		AccountMetaSlice: make(sol.AccountMetaSlice, 17),
	}
	return nd
}

// setBaseAmountOut sets the "baseAmountOut" parameter.
func (inst *Buy) setBaseAmountOut(baseAmountOut uint64) *Buy {
	inst.BaseAmountOut = &baseAmountOut
	return inst
}

// setMaxQuoteAmountIn sets the "maxQuoteAmountIn" parameter.
func (inst *Buy) setMaxQuoteAmountIn(maxQuoteAmountIn uint64) *Buy {
	inst.MaxQuoteAmountIn = &maxQuoteAmountIn
	return inst
}

// setPoolAccount sets the "pool" account.
func (inst *Buy) setPoolAccount(pool sol.PublicKey) *Buy {
	inst.AccountMetaSlice[0] = sol.Meta(pool)
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *Buy) setUserAccount(user sol.PublicKey) *Buy {
	inst.AccountMetaSlice[1] = sol.Meta(user).WRITE().SIGNER()
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *Buy) setMintAccount(mint sol.PublicKey) *Buy {
	inst.AccountMetaSlice[3] = sol.Meta(mint)
	return inst
}

// setUserBaseTokenAccount sets the "useBaseTokenAccount" account.
func (inst *Buy) setUserBaseTokenAccount(useBaseTokenAccount sol.PublicKey) *Buy {
	inst.AccountMetaSlice[5] = sol.Meta(useBaseTokenAccount).WRITE()
	return inst
}

// setUserQuoteTokenAccount sets the "userQuoteTokenAccount" account.
func (inst *Buy) setUserQuoteTokenAccount(userQuoteTokenAccount sol.PublicKey) *Buy {
	inst.AccountMetaSlice[6] = sol.Meta(userQuoteTokenAccount).WRITE()
	return inst
}

// setPoolBaseTokenAccount sets the "poolBaseTokenAccount" account.
func (inst *Buy) setPoolBaseTokenAccount(poolBaseTokenAccount sol.PublicKey) *Buy {
	inst.AccountMetaSlice[7] = sol.Meta(poolBaseTokenAccount).WRITE()
	return inst
}

// setPoolQuoteTokenAccount sets the "associatedUser" account.
func (inst *Buy) setPoolQuoteTokenAccount(poolQuoteTokenAccount sol.PublicKey) *Buy {
	inst.AccountMetaSlice[8] = sol.Meta(poolQuoteTokenAccount).WRITE()
	return inst
}

func (inst *Buy) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[2] = sol.Meta(Global)
	inst.AccountMetaSlice[4] = sol.Meta(WSOL)
	inst.AccountMetaSlice[9] = sol.Meta(FeeRecipient)
	inst.AccountMetaSlice[10] = sol.Meta(FeeRecipientTokenAccount).WRITE()
	inst.AccountMetaSlice[11] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[12] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[13] = sol.Meta(SystemProgramId)
	inst.AccountMetaSlice[14] = sol.Meta(AssociatedTokenProgramId)
	inst.AccountMetaSlice[15] = sol.Meta(EventAuthority)
	inst.AccountMetaSlice[16] = sol.Meta(PumpAMMProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Buy,
	}}
}

func (inst Buy) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	// Serialize `BaseAmountOut` param:
	err = encoder.Encode(inst.BaseAmountOut)
	if err != nil {
		return err
	}
	// Serialize `MaxQuoteAmountIn` param:
	err = encoder.Encode(inst.MaxQuoteAmountIn)
	if err != nil {
		return err
	}
	return nil
}
