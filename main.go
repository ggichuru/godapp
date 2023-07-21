package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ggichuru/godapp/configs"
	"github.com/ggichuru/godapp/wallet"
)

// define map to store env variable
// var myenv map[string]string

func main() {
	configs.LoadEnv()

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
	wallet.ImportKeystore()
}
