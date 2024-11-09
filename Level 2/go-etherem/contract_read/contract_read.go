package main

import (
	"etherem/client"
	store "etherem/contracts"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	client, err := client.EthereumClient()
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x5F8EbC14152f832E20cFc89c8FFe3b6348C655AE")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}
