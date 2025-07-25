package pumpfunamm

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

// Sells tokens into a bonding curve.
type Sell struct {
	BaseAmountIn      *uint64
	MinQuoteAmountOut *uint64

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
	//
	// [17] = [] Coin Creator Vault Ata
	//
	// [18] = [] Coin Creator Vault Authority
	sol.AccountMetaSlice `bin:"-"`
}

// NewSellInstruction declares a new Sell instruction with the provided parameters and accounts.
func NewSellInstruction(
	// Parameters:
	baseAmountIn uint64,
	minQuoteAmountOut uint64,
	// Accounts:
	pool sol.PublicKey,
	baseMint sol.PublicKey,
	quoteMint sol.PublicKey,
	useBaseTokenAccount sol.PublicKey,
	userQuoteTokenAccount sol.PublicKey,
	poolBaseTokenAccount sol.PublicKey,
	poolQuoteTokenAccount sol.PublicKey,
	feeRecipientTokenAccount sol.PublicKey,
	coinCreatorVaultAta sol.PublicKey,
	coinCreatorVaultAuthority sol.PublicKey,
	user sol.PublicKey) *Sell {
	return newSelInstructionBuilder().
		setBaseAmountIn(baseAmountIn).
		setMinQuoteAmountOut(minQuoteAmountOut).
		setPoolAccount(pool).
		setBaseMintAccount(baseMint).
		setQuoteMintAccount(quoteMint).
		setUserBaseTokenAccount(useBaseTokenAccount).
		setUserQuoteTokenAccount(userQuoteTokenAccount).
		setPoolBaseTokenAccount(poolBaseTokenAccount).
		setPoolQuoteTokenAccount(poolQuoteTokenAccount).
		setProtocolFeeRecipientTokenAccount(feeRecipientTokenAccount).
		setCoinCreatorVaultAta(coinCreatorVaultAta).
		setCoinCreatorVaultAuthority(coinCreatorVaultAuthority).
		setUserAccount(user)
}

// NewSellInstructionBuilder creates a new `Sell` instruction builder.
func newSelInstructionBuilder() *Sell {
	nd := &Sell{
		AccountMetaSlice: make(sol.AccountMetaSlice, 19),
	}
	return nd
}

// setBaseAmountIn sets the "baseAmountIn" parameter.
func (inst *Sell) setBaseAmountIn(baseAmountIn uint64) *Sell {
	inst.BaseAmountIn = &baseAmountIn
	return inst
}

// setMinQuoteAmountOut sets the "minQuoteAmountOut" parameter.
func (inst *Sell) setMinQuoteAmountOut(minQuoteAmountOut uint64) *Sell {
	inst.MinQuoteAmountOut = &minQuoteAmountOut
	return inst
}

// setPoolAccount sets the "pool" account.
func (inst *Sell) setPoolAccount(pool sol.PublicKey) *Sell {
	inst.AccountMetaSlice[0] = sol.Meta(pool)
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *Sell) setUserAccount(user sol.PublicKey) *Sell {
	inst.AccountMetaSlice[1] = sol.Meta(user).WRITE().SIGNER()
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *Sell) setBaseMintAccount(baseMint sol.PublicKey) *Sell {
	inst.AccountMetaSlice[3] = sol.Meta(baseMint)
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *Sell) setQuoteMintAccount(quoteMint sol.PublicKey) *Sell {
	inst.AccountMetaSlice[4] = sol.Meta(quoteMint)
	return inst
}

// setUserBaseTokenAccount sets the "useBaseTokenAccount" account.
func (inst *Sell) setUserBaseTokenAccount(useBaseTokenAccount sol.PublicKey) *Sell {
	inst.AccountMetaSlice[5] = sol.Meta(useBaseTokenAccount).WRITE()
	return inst
}

// setUserQuoteTokenAccount sets the "userQuoteTokenAccount" account.
func (inst *Sell) setUserQuoteTokenAccount(userQuoteTokenAccount sol.PublicKey) *Sell {
	inst.AccountMetaSlice[6] = sol.Meta(userQuoteTokenAccount).WRITE()
	return inst
}

// setPoolBaseTokenAccount sets the "poolBaseTokenAccount" account.
func (inst *Sell) setPoolBaseTokenAccount(poolBaseTokenAccount sol.PublicKey) *Sell {
	inst.AccountMetaSlice[7] = sol.Meta(poolBaseTokenAccount).WRITE()
	return inst
}

// setPoolQuoteTokenAccount sets the "associatedUser" account.
func (inst *Sell) setPoolQuoteTokenAccount(poolQuoteTokenAccount sol.PublicKey) *Sell {
	inst.AccountMetaSlice[8] = sol.Meta(poolQuoteTokenAccount).WRITE()
	return inst
}
func (inst *Sell) setProtocolFeeRecipientTokenAccount(protocolFeeRecipientTokenAccount sol.PublicKey) *Sell {
	inst.AccountMetaSlice[10] = sol.Meta(protocolFeeRecipientTokenAccount).WRITE()
	return inst
}
func (inst *Sell) setCoinCreatorVaultAta(coinCreatorVaultAta sol.PublicKey) *Sell {
	inst.AccountMetaSlice[17] = sol.Meta(coinCreatorVaultAta).WRITE()
	return inst
}
func (inst *Sell) setCoinCreatorVaultAuthority(coinCreatorVaultAuthority sol.PublicKey) *Sell {
	inst.AccountMetaSlice[18] = sol.Meta(coinCreatorVaultAuthority)
	return inst
}

func (inst *Sell) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[2] = sol.Meta(Global)
	inst.AccountMetaSlice[9] = sol.Meta(FeeRecipient)
	// inst.AccountMetaSlice[10] = sol.Meta(FeeRecipientTokenAccount).WRITE()
	inst.AccountMetaSlice[11] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[12] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[13] = sol.Meta(SystemProgramId)
	inst.AccountMetaSlice[14] = sol.Meta(AssociatedTokenProgramId)
	inst.AccountMetaSlice[15] = sol.Meta(EventAuthority)
	inst.AccountMetaSlice[16] = sol.Meta(PumpAMMProgramId)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Sell,
	}}
}

func (inst Sell) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	// Serialize `BaseAmountIn` param:
	err = encoder.Encode(inst.BaseAmountIn)
	if err != nil {
		return err
	}
	// Serialize `MinQuoteAmountOut` param:
	err = encoder.Encode(inst.MinQuoteAmountOut)
	if err != nil {
		return err
	}
	return nil
}
