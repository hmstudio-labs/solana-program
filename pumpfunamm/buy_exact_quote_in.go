package pumpfunamm

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type BuyExactQuoteIn struct {
	BaseAmountIn     *uint64
	MinBaseAmountOut *uint64

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
	//
	// [17] = [] Coin Creator Vault Ata
	//
	// [18] = [] Coin Creator Vault Authority
	//
	// [19] = [] global volume accumulator
	//
	// [20] = [] user volume accumulator
	sol.AccountMetaSlice `bin:"-"`
}

// NewBuyExactQuoteInInstruction declares a new BuyExactQuoteIn instruction with the provided parameters and accounts.
func NewBuyExactQuoteInInstruction(
	// Parameters:
	baseAmountIn uint64,
	minBaseAmountOut uint64,
	// Accounts:
	pool sol.PublicKey,
	baseMint sol.PublicKey,
	quoteMint sol.PublicKey,
	baseTokenProgram sol.PublicKey,
	useBaseTokenAccount sol.PublicKey,
	userQuoteTokenAccount sol.PublicKey,
	poolBaseTokenAccount sol.PublicKey,
	poolQuoteTokenAccount sol.PublicKey,
	feeRecipient sol.PublicKey,
	feeRecipientTokenAccount sol.PublicKey,
	coinCreatorVaultAta sol.PublicKey,
	coinCreatorVaultAuthority sol.PublicKey,
	userVolumeAccumulator sol.PublicKey,
	user sol.PublicKey) *BuyExactQuoteIn {
	return newBuyExactQuoteInInstructionBuilder().
		setBaseAmountIn(baseAmountIn).
		setMinBaseAmountOut(minBaseAmountOut).
		setPoolAccount(pool).
		setBaseMintAccount(baseMint).
		setQuoteMintAccount(quoteMint).
		setBaseTokenProgram(baseTokenProgram).
		setUserBaseTokenAccount(useBaseTokenAccount).
		setUserQuoteTokenAccount(userQuoteTokenAccount).
		setPoolBaseTokenAccount(poolBaseTokenAccount).
		setPoolQuoteTokenAccount(poolQuoteTokenAccount).
		setProtocolFeeRecipient(feeRecipient).
		setProtocolFeeRecipientTokenAccount(feeRecipientTokenAccount).
		setCoinCreatorVaultAta(coinCreatorVaultAta).
		setCoinCreatorVaultAuthority(coinCreatorVaultAuthority).
		setUserVolumeAccumulator(userVolumeAccumulator).
		setUserAccount(user)
}

// NewBuyExactQuoteInInstructionBuilder creates a new `BuyExactQuoteIn` instruction builder.
func newBuyExactQuoteInInstructionBuilder() *BuyExactQuoteIn {
	nd := &BuyExactQuoteIn{
		AccountMetaSlice: make(sol.AccountMetaSlice, 23),
	}
	return nd
}

// setBaseAmountIn sets the "baseAmountIn" parameter.
func (inst *BuyExactQuoteIn) setBaseAmountIn(baseAmountIn uint64) *BuyExactQuoteIn {
	inst.BaseAmountIn = &baseAmountIn
	return inst
}

// setMinBaseAmountOut sets the "minBaseAmountOut" parameter.
func (inst *BuyExactQuoteIn) setMinBaseAmountOut(minBaseAmountOut uint64) *BuyExactQuoteIn {
	inst.MinBaseAmountOut = &minBaseAmountOut
	return inst
}

// setPoolAccount sets the "pool" account.
func (inst *BuyExactQuoteIn) setPoolAccount(pool sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[0] = sol.Meta(pool).WRITE()
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *BuyExactQuoteIn) setUserAccount(user sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[1] = sol.Meta(user).WRITE().SIGNER()
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *BuyExactQuoteIn) setBaseMintAccount(baseMint sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[3] = sol.Meta(baseMint)
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *BuyExactQuoteIn) setQuoteMintAccount(quoteMint sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[4] = sol.Meta(quoteMint)
	return inst
}

// setUserBaseTokenAccount sets the "useBaseTokenAccount" account.
func (inst *BuyExactQuoteIn) setUserBaseTokenAccount(useBaseTokenAccount sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[5] = sol.Meta(useBaseTokenAccount).WRITE()
	return inst
}

// setUserQuoteTokenAccount sets the "userQuoteTokenAccount" account.
func (inst *BuyExactQuoteIn) setUserQuoteTokenAccount(userQuoteTokenAccount sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[6] = sol.Meta(userQuoteTokenAccount).WRITE()
	return inst
}

// setPoolBaseTokenAccount sets the "poolBaseTokenAccount" account.
func (inst *BuyExactQuoteIn) setPoolBaseTokenAccount(poolBaseTokenAccount sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[7] = sol.Meta(poolBaseTokenAccount).WRITE()
	return inst
}

// setPoolQuoteTokenAccount sets the "associatedUser" account.
func (inst *BuyExactQuoteIn) setPoolQuoteTokenAccount(poolQuoteTokenAccount sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[8] = sol.Meta(poolQuoteTokenAccount).WRITE()
	return inst
}
func (inst *BuyExactQuoteIn) setProtocolFeeRecipient(protocolFeeRecipient sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[9] = sol.Meta(protocolFeeRecipient)
	return inst
}
func (inst *BuyExactQuoteIn) setProtocolFeeRecipientTokenAccount(protocolFeeRecipientTokenAccount sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[10] = sol.Meta(protocolFeeRecipientTokenAccount).WRITE()
	return inst
}
func (inst *BuyExactQuoteIn) setBaseTokenProgram(tokenProgram sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[11] = sol.Meta(tokenProgram)
	return inst
}
func (inst *BuyExactQuoteIn) setCoinCreatorVaultAta(coinCreatorVaultAta sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[17] = sol.Meta(coinCreatorVaultAta).WRITE()
	return inst
}
func (inst *BuyExactQuoteIn) setCoinCreatorVaultAuthority(coinCreatorVaultAuthority sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[18] = sol.Meta(coinCreatorVaultAuthority)
	return inst
}
func (inst *BuyExactQuoteIn) setUserVolumeAccumulator(userVolumeAccumulator sol.PublicKey) *BuyExactQuoteIn {
	inst.AccountMetaSlice[20] = sol.Meta(userVolumeAccumulator).WRITE()
	return inst
}
func (inst *BuyExactQuoteIn) Build() *Instruction {
	// 构建accounts
	inst.AccountMetaSlice[2] = sol.Meta(Global)
	// inst.AccountMetaSlice[9] = sol.Meta(FeeRecipient)
	// inst.AccountMetaSlice[10] = sol.Meta(FeeRecipientTokenAccount).WRITE()
	// inst.AccountMetaSlice[11] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[12] = sol.Meta(TokenProgramId)
	inst.AccountMetaSlice[13] = sol.Meta(SystemProgramId)
	inst.AccountMetaSlice[14] = sol.Meta(AssociatedTokenProgramId)
	inst.AccountMetaSlice[15] = sol.Meta(EventAuthority)
	inst.AccountMetaSlice[16] = sol.Meta(PumpAMMProgramId)
	inst.AccountMetaSlice[19] = sol.Meta(GlobalVolumeAccumulator).WRITE()
	inst.AccountMetaSlice[21] = sol.Meta(FeeConfig)
	inst.AccountMetaSlice[22] = sol.Meta(FeeProgram)
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BuyExactQuoteIn,
	}}
}

func (inst BuyExactQuoteIn) MarshalWithEncoder(encoder *bin.Encoder) (err error) {
	// Serialize `BaseAmountIn` param:
	err = encoder.Encode(inst.BaseAmountIn)
	if err != nil {
		return err
	}
	// Serialize `MinBaseAmountOut` param:
	err = encoder.Encode(inst.MinBaseAmountOut)
	if err != nil {
		return err
	}
	return nil
}
