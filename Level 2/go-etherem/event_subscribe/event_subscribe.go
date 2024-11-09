package main

import (
	"context"
	"etherem/client"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	client, err := client.EthereumClient()
	if err != nil {
		log.Fatal(err)
	}

	constractAddress := common.HexToAddress("0x5F8EbC14152f832E20cFc89c8FFe3b6348C655AE")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{constractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
