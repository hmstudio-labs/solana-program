package systempro

import (
	sol "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	token2022 "github.com/gagliardetto/solana-go/programs/token-2022"
)

func NewTokenAccount(owner sol.PublicKey, mint sol.PublicKey, lamports uint64) (*sol.PublicKey, error) {
	seed := mint.String()[0:32]
	tokenAccount, err := sol.CreateWithSeed(
		owner,
		seed,
		sol.TokenProgramID,
	)
	if err != nil {
		return nil, err
	}
	return &tokenAccount, nil
}

func NewTokenAccountV2(owner sol.PublicKey, mint sol.PublicKey, tokenProgram sol.PublicKey) (*sol.PublicKey, error) {
	seed := mint.String()[0:32]
	tokenAccount, err := sol.CreateWithSeed(owner, seed, tokenProgram)
	if err != nil {
		return nil, err
	}
	return &tokenAccount, nil
}

func NewAccountAndInstructions(owner sol.PublicKey, mint sol.PublicKey, lamports uint64) (*sol.PublicKey, sol.Instruction, sol.Instruction, sol.Instruction, error) {
	seed := mint.String()[0:32]
	tokenAccount, err := sol.CreateWithSeed(
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
		tokenAccount,
		owner,
	).ValidateAndBuild()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	initTokenAccount, err := token.NewInitializeAccountInstruction(
		tokenAccount,
		mint,
		owner,
		sol.SysVarRentPubkey,
	).ValidateAndBuild()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	closeAccInst, err := token.NewCloseAccountInstruction(
		tokenAccount,
		owner,
		owner,
		[]sol.PublicKey{},
	).ValidateAndBuild()

	if err != nil {
		return nil, nil, nil, nil, err
	}
	return &tokenAccount, createAccountWithSeedIx, initTokenAccount, closeAccInst, nil
}

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

func NewAccountAndInstructionsV2(owner sol.PublicKey, mint sol.PublicKey, tokenProgram sol.PublicKey) (*sol.PublicKey, sol.Instruction, sol.Instruction, sol.Instruction, error) {
	var lamports uint64 = 2039290
	var space uint64 = 165
	if tokenProgram.Equals(sol.Token2022ProgramID) {
		lamports = 2157600
		space = 182
	}
	seed := mint.String()[0:32]
	tokenAccount, err := sol.CreateWithSeed(
		owner,
		seed,
		tokenProgram,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	createAccountWithSeedIx, err := system.NewCreateAccountWithSeedInstruction(
		owner,
		seed,
		lamports,
		space,
		tokenProgram,
		owner,
		tokenAccount,
		owner,
	).ValidateAndBuild()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if tokenProgram.Equals(sol.TokenProgramID) {
		initTokenAccount, err := token.NewInitializeAccountInstruction(
			tokenAccount,
			mint,
			owner,
			sol.SysVarRentPubkey,
		).ValidateAndBuild()
		if err != nil {
			return nil, nil, nil, nil, err
		}
		closeAccInst, err := token.NewCloseAccountInstruction(
			tokenAccount,
			owner,
			owner,
			[]sol.PublicKey{},
		).ValidateAndBuild()

		if err != nil {
			return nil, nil, nil, nil, err
		}
		return &tokenAccount, createAccountWithSeedIx, initTokenAccount, closeAccInst, nil
	}
	initTokenAccount, err := token2022.NewInitializeAccount3Instruction(
		owner,
		tokenAccount,
		mint,
	).ValidateAndBuild()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	closeAccInst, err := token2022.NewCloseAccountInstruction(
		tokenAccount,
		owner,
		owner,
		[]sol.PublicKey{},
	).ValidateAndBuild()

	if err != nil {
		return nil, nil, nil, nil, err
	}
	return &tokenAccount, createAccountWithSeedIx, initTokenAccount, closeAccInst, nil
}

func NewCloseAccountInstruction(tokenAccount sol.PublicKey, owner sol.PublicKey, tokenProgram sol.PublicKey) sol.Instruction {
	if tokenProgram.Equals(sol.TokenProgramID) {
		return token.NewCloseAccountInstruction(
			tokenAccount,
			owner,
			owner,
			[]sol.PublicKey{},
		).Build()
	}
	return token2022.NewCloseAccountInstruction(
		tokenAccount,
		owner,
		owner,
		[]sol.PublicKey{},
	).Build()
}
