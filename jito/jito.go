package jito

import (
	"math/rand"
	"time"

	sol "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
)

var jitoTipAccounts = []string{
	"96gYZGLnJYVFmbjzopPSU6QiEV5fGqZNyN9nmNhvrZU5",
	"HFqU5x63VTqvQss8hp11i4wVV8bD44PvwucfZ2bU7gRe",
	"Cw8CFyM9FkoMi7K7Crf6HNQqf4uEMzpKw6QNghXLvLkY",
	"ADaUMid9yfUytqMBgopwjb2DTLSokTSzL1zt6iGPaS49",
	"DfXygSm4jCyNCybVYYK6DwvWqjKee8pbDmJGcLWNDXjh",
	"ADuUkR4vqLUMWXxW9gh6D6L8pMSawimctcNZ5pGwDcEt",
	"DttWaMuVvTiduZRnguLF7jNxTgiMBZ1hyAumKUiL2KRL",
	"3AVi9Tg9Uo68tJfuvoKvqKNWKkC5wPdSSdeBnizKZ6jT",
}

// 获取随机 Jito Tip 公钥
func getJitoTipPublicKey() sol.PublicKey {
	// 初始化随机数种子
	// 初始化一个新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 随机选择一个账户
	pickAccount := r.Intn(len(jitoTipAccounts))
	jitoTipAccount := jitoTipAccounts[pickAccount]

	// 转换为 PublicKey 类型
	publicKey := sol.MustPublicKeyFromBase58(jitoTipAccount)
	return publicKey
}

func NewJitoIx(amount uint64, owner sol.PublicKey) sol.Instruction {
	to := getJitoTipPublicKey()
	return system.NewTransferInstruction(
		amount,
		owner,
		to,
	).Build()
}
