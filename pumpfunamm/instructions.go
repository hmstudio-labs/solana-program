package pumpfunamm

import (
	"bytes"
	"fmt"

	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

var (
	Global                   = sol.MustPublicKeyFromBase58("ADyA8hdefvWN2dbGGWFotbzWxrAvLW83WG6QCVXvJKqw")
	FeeRecipient             = sol.MustPublicKeyFromBase58("G5UZAVbAf46s7cKWoyKu8kYTip9DGTpbLZ2qa9Aq69dP")
	FeeRecipientTokenAccount = sol.MustPublicKeyFromBase58("BWXT6RUhit9FfJQM3pBmqeFLPYmuxgmyhMGC5sGr8RbA")
	SystemProgramId          = sol.MustPublicKeyFromBase58("11111111111111111111111111111111")
	TokenProgramId           = sol.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	AssociatedTokenProgramId = sol.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")
	RentProgramId            = sol.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")
	EventAuthority           = sol.MustPublicKeyFromBase58("GS4CU59F31iL7aR2Q8zVS8DRrcRnXX1yjQ66TqNVQnaR")
	PumpAMMProgramId         = sol.MustPublicKeyFromBase58("pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA")
	WSOL                     = sol.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112")
)

var (
	// Creates the global state.
	// Instruction_Initialize = bin.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	// Sets the global state parameters.
	// Instruction_SetParams = bin.TypeID([8]byte{27, 234, 178, 52, 147, 2, 187, 141})

	// Creates a new coin and bonding curve.
	// Instruction_Create = bin.TypeID([8]byte{24, 30, 200, 40, 5, 28, 7, 119})

	// Buys tokens from a bonding curve.
	Instruction_Buy = bin.TypeID([8]byte{102, 6, 61, 18, 1, 218, 235, 234})

	// Sells tokens into a bonding curve.
	Instruction_Sell = bin.TypeID([8]byte{51, 230, 133, 164, 1, 127, 131, 173})

	// Allows the admin to withdraw liquidity for a migration once the bonding curve completes
	// Instruction_Withdraw = bin.TypeID([8]byte{183, 18, 70, 156, 148, 109, 161, 34})
)

type Instruction struct {
	bin.BaseVariant
}

func (inst *Instruction) ProgramID() sol.PublicKey {
	return PumpAMMProgramId
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
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}
