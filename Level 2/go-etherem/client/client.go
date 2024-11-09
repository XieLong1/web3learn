package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func EthereumClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://rpc.ankr.com/eth_sepolia/2b2a1b783dede8d6d12cd65f273a2f8055f5bc08f5cf2bd3631944b5f745875c")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %w", err)
	}
	return client, nil
}
