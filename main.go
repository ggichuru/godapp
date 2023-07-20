package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// define map to store env variable
// var myenv map[string]string

// constant highlighting the env source file.
const envloc = ".env"

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loadind env from: %s \n %v", envloc, err)
	}
}

func connect_ethclient() ethclient.Client {
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

func main() {
	loadEnv()

	client := connect_ethclient()
	fmt.Print(client)

	// convert address to common.address
	address := common.HexToAddress(os.Getenv("PUBLIC_KEY"))

	fmt.Println(address.Hex())
}
