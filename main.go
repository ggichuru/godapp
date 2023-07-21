package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

func _getBalance(acc_addr common.Address, client ethclient.Client, decimals ...int) (*big.Float, *big.Float) {
	ctx := context.Background()

	// assign decimals
	decimal := 18
	if len(decimals) > 0 {
		decimal = decimals[0]
	}

	_balance, err := client.BalanceAt(ctx, acc_addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	_pending_bal, err := client.PendingBalanceAt(ctx, acc_addr)
	if err != nil {
		log.Fatal(err)
	}

	// convert wei to float
	balance := new(big.Float)
	balance.SetString(_balance.String())

	// balance in plaform token
	bal_eth := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(decimal)))

	balance.SetString(_pending_bal.String())

	// pending balance
	bal_pending := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(decimal)))
	return bal_eth, bal_pending
}

func createWallet() (string, string) {
	// Create private key
	_privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(_privateKey)
	privateKey := hexutil.Encode(privateKeyBytes)[2:]

	_ = privateKey

	// Create public key
	_publicKey := _privateKey.Public()
	publicKeyECDSA, ok := _publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("Error Casting public key to ECSDA")
	}
	publicKey := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKey, publicKey
}

func main() {
	loadEnv()

	/** GLOBALS*/
	client := connect_ethclient()

	// convert acc_addr to common.Address
	acc_addr := common.HexToAddress(os.Getenv("PUBLIC_KEY"))

	/** FN CALLS*/
	// balances
	bal_eth, bal_pending := _getBalance(acc_addr, client)
	fmt.Printf("Wallet Balance: %g MATIC \n", bal_eth)
	fmt.Printf("Pending Balance: %g MATIC \n", bal_pending)

	fmt.Println()

	// Wallet
	private_key, public_key := createWallet()
	fmt.Printf("Public Key: %s\n", public_key)
	fmt.Printf("Private Key: %s\n", private_key)
}
