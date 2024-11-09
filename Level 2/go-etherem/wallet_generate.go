package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main_2() {
	// privateKey, err := crypto.GenerateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// fmt.Println(hexutil.Encode(privateKeyBytes[2:])) // 私钥 0x4945c1fbad24687b4689c7352fcf4f54ad54ffa4e9bfd0de38554eb827a1

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // b6a06606ab0f7432177636905ef062b149f944c624dc7ef0d362d0753ad9e5773d6b1729c93163829d583162f15e957e27b00d4d6e31d190452c274e75c765cc

	// address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// fmt.Println(address) // 0x81Be55aa8de0b4d20bC8AFd7615d9202B1FbaC64 公钥地址

	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x81be55aa8de0b4d20bc8afd7615d9202b1fbac64 公钥地址

	// 检查地址是否有效·································································································分割
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x81be55aa8de0b4d20bc8afd7615d9202b1fbac64")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	// 检查是否是只能合约地址
	// 0x Protocol Token (ZRX) smart contract address
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	address1 := common.HexToAddress("0x81be55aa8de0b4d20bc8afd7615d9202b1fbac64")
	bytecode, err := client.CodeAt(context.Background(), address1, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract) // is contract: true 合约地址

	// a random user account address
	address1 = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address1, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false 账户地址
}
