# Solana Program SDK

A Go SDK for interacting with various Solana programs including Pumpfun, PumpfunAMM, Jito, and Raydium. This SDK provides a simple interface to execute common operations on Solana blockchain.

## Features

- **ComputeBudgetProgram**: Set compute unit limits for transactions
- **Pumpfun Program**: Bonding curve trading operations
  - Buy tokens from bonding curve
  - Sell tokens to bonding curve
- **PumpfunAMM Program**: Automated Market Maker operations
  - Buy tokens from liquidity pool
  - Sell tokens to liquidity pool
- **Raydium Program**: Integration with Raydium protocol
  - Execute buy orders
  - Execute sell orders
- **Jito**: Additional Solana program integrations

## Installation

```bash
go get github.com/hmstudio-labs/solana-program
```

## Usage

### Pumpfun Program
```go
package main

import (
	"github.com/hmstudio-labs/solana-program/pumpfun"
	"github.com/gagliardetto/solana-go"
)

func main() {
	// Initialize client
	client := pumpfun.NewClient()
	
	// Example: Sell tokens
	tx, err := client.Sell(
		solana.MustPublicKeyFromBase58("tokenMintAddress"),
		solana.MustPublicKeyFromBase58("bondingCurveAddress"),
		solana.MustPublicKeyFromBase58("associatedBondingCurve"),
		solana.MustPublicKeyFromBase58("associatedUser"),
		solana.MustPublicKeyFromBase58("userWallet"),
		1000000, // amount
		500000,  // minSolOutput
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction: %s\n", tx)
}
```

### PumpfunAMM Program
```go
package main

import (
	"github.com/hmstudio-labs/solana-program/pumpfunamm"
	"github.com/gagliardetto/solana-go"
)

func main() {
	// Example: Buy from liquidity pool
	tx, err := pumpfunamm.Buy(
		solana.MustPublicKeyFromBase58("poolAddress"),
		solana.MustPublicKeyFromBase58("userWallet"),
		1000000, // amount
		500000,  // maxSolCost
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction: %s\n", tx)
}
```

### Raydium Program
```go
package main

import (
	"github.com/hmstudio-labs/solana-program/raydium"
	"github.com/gagliardetto/solana-go"
)

func main() {
	// Example: Execute swap
	tx, err := raydium.Swap(
		solana.MustPublicKeyFromBase58("marketAddress"),
		solana.MustPublicKeyFromBase58("userWallet"),
		1000000, // amountIn
		500000,  // minAmountOut
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction: %s\n", tx)
}
```

### Jito Program
```go
package main

import (
	"github.com/hmstudio-labs/solana-program/jito"
	"github.com/gagliardetto/solana-go"
)

func main() {
	// Example: Stake tokens
	tx, err := jito.Stake(
		solana.MustPublicKeyFromBase58("validatorAddress"),
		solana.MustPublicKeyFromBase58("userWallet"),
		1000000, // amount
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction: %s\n", tx)
}
```
