package systempro

import (
	sol "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
)

func NewWSOLAccountAndInstructions(owner sol.PublicKey, lamports uint64) (*sol.PublicKey, sol.Instruction, sol.Instruction, sol.Instruction, error) {
	seed := sol.NewWallet().PublicKey().String()[0:32]
	wrappedSolAccount, err := sol.CreateWithSeed(
		owner,
		seed,
		sol.TokenProgramID,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	createAccountWithSeedIx, err := system.NewCreateAccountWithSeedInstruction(
		owner,
		seed,
		lamports,
		165,
		sol.TokenProgramID,
		owner,
		wrappedSolAccount,
		owner,
	).ValidateAndBuild()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	initTokenAccount, err := token.NewInitializeAccountInstruction(
		wrappedSolAccount,
		sol.WrappedSol,
		owner,
		sol.SysVarRentPubkey,
	).ValidateAndBuild()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	closeAccInst, err := token.NewCloseAccountInstruction(
		wrappedSolAccount,
		owner,
		owner,
		[]sol.PublicKey{},
	).ValidateAndBuild()

	if err != nil {
		return nil, nil, nil, nil, err
	}
	return &wrappedSolAccount, createAccountWithSeedIx, initTokenAccount, closeAccInst, nil
}
