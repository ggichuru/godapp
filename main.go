package main

import (
	"context"
	"fmt"

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
	acc_addr := transactions.ExtractPublicKey()

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
	block := transactions.SimulateBlock(client, ctx)
	fmt.Println("Block Hash: ", block.Hash())

	fmt.Println()

	// tx query
	// transactions.GetTransactions(block, client, ctx)

	fmt.Println()

	// Transfer eth
	// transactions.SetupTransfer(client, ctx)

	// Transfer ERC20
	to_address := common.HexToAddress("0x3F92A2952746be63f8E22D58997A9A56c95ed2D1")
	transactions.TransferErc20(client, ctx, acc_addr, to_address, 20)

	// fmt.Println(transactions.ToWei(0, 20, 18))
}
