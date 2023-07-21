package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ggichuru/godapp/configs"
	"github.com/ggichuru/godapp/transactions"
	"github.com/ggichuru/godapp/wallet"
)

// define map to store env variable
// var myenv map[string]string

func main() {
	configs.LoadEnv()
	ctx := context.Background()

	/** GLOBALS*/
	client := configs.ConnectEthClient()

	// convert acc_addr to common.Address
	acc_addr := common.HexToAddress(os.Getenv("PUBLIC_KEY"))

	/** FN CALLS*/
	// balances
	wallet.PreviewBalance(acc_addr, client)

	// Wallet
	wallet.PreviewWallet(acc_addr, client)

	// Keystore
	// wallet.CreateKeystore()
	// wallet.ImportKeystore()

	// Address Checks
	if isEvm := wallet.IsEvmAddr("0x323b5d4c32345ced77393b3530b1eed0f346429d"); isEvm {
		fmt.Println("address is evm")
	}

	if isContract := wallet.IsContract(client, ctx, acc_addr); isContract {
		fmt.Println("address is Contract")
	} else {
		fmt.Println("address is EOA")
	}

	fmt.Println()

	// Block Query
	latest_block := transactions.GetCurrentBlock(client, ctx)
	fmt.Println("Latest Block:", latest_block)

}
