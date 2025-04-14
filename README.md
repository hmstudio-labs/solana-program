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
go get github.com/your-username/solana-program
```

## Usage

### Setting Compute Unit Limit

```go
// Set compute unit limit for transaction
ComputeBudgetProgram.setComputeUnitLimit()
```

### Pumpfun Trading

```go
// Buy tokens from bonding curve
PumpfunProgram.setBuy()

// Sell tokens to bonding curve
PumpfunProgram.setSell()
```

### PumpfunAMM Trading

```go
// Buy tokens from AMM pool
PumpfunAMMProgram.setBuy()

// Sell tokens to AMM pool
PumpfunAMMProgram.setSell()
```

### Raydium Trading

```go
// Execute buy order on Raydium
RaydiumProgram.setBuy()

// Execute sell order on Raydium
RaydiumProgram.setSell()
```

## License

MIT