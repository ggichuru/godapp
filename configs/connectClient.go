package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectEthClient() ethclient.Client {
	gateway_endpoint := os.Getenv("RPC_URL") + os.Getenv("INFURA_APIKEY")

	// initialize go-ethereum package
	client, err := ethclient.Dial(gateway_endpoint)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\t\t\t**Gateway connection established**")
	fmt.Println()

	return *client
}
