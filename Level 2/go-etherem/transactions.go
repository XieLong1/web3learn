package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main_4() {
	client, err := ethclient.Dial("https://rpc.ankr.com/eth/2b2a1b783dede8d6d12cd65f273a2f8055f5bc08f5cf2bd3631944b5f745875c")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(21141430)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex())
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(receipt.Status)
	}
	// // 您还可以使用 TransactionByHash 在给定具体事务哈希值的情况下直接查询单个事务。
	// txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	// tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	// fmt.Println(isPending)       // false
}
