package wallet

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(acc_addr common.Address, client ethclient.Client, decimals ...int) (*big.Float, *big.Float) {
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

func PreviewBalance(acc_addr common.Address, client ethclient.Client) {
	bal_eth, bal_pending := GetBalance(acc_addr, client)
	fmt.Printf("Wallet Balance: %g MATIC \n", bal_eth)
	fmt.Printf("Pending Balance: %g MATIC \n", bal_pending)

	fmt.Println()
}
