package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func clent() {
// 	client, err := ethclient.Dial("https://cloudflare-eth.com")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("we have a connection")
// 	return client
// }

func account() {
	address := common.HexToAddress("0xfeF8088ee49Fc9f8b5d2ABb2D1271b2fD261811A")
	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())
	fmt.Println(common.HexToHash("0xfeF8088ee49Fc9f8b5d2ABb2D1271b2fD261811A"))
}

func balance(client *ethclient.Client) {
	account := common.HexToAddress("0xfeF8088ee49Fc9f8b5d2ABb2D1271b2fD261811A")
	// balance, err := client.BalanceAt(context.Background(), account, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(balance)

	// 获取本地节点最新区块
	// 获取当前区块高度
	blockNumber, err := client.BlockNumber(context.Background()) // nil 表示最新的区块
	if err != nil {
		log.Fatalf("获取区块高度失败: %v", err)
	}
	// 打印区块高度
	fmt.Printf("当前区块高度: %d\n", blockNumber)
	// 获取当前区块高度时的账户余额
	blockNumber1 := big.NewInt(0)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	// 处理余额小数点问题
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethVal := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethVal)

	// 余额加待处理中的余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pendingBalance) // 25729324269165216042
}

func main_1() {
	// 链接
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	// 账户
	// account()

	// 账户金额
	balance(client)
}
