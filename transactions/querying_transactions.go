package transactions

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getBlockTx(block *types.Block) types.Transactions {
	return block.Transactions()
}

func getTx(client ethclient.Client, ctx context.Context, tx_hash common.Hash) (*types.Transaction, bool) {
	tx, isPending, err := client.TransactionByHash(ctx, tx_hash)
	if err != nil {
		log.Fatal(err)
	}

	return tx, isPending
}

func checkTxReceipt(client ethclient.Client, ctx context.Context, tx_hash common.Hash) uint64 {
	if tx_receipt, err := client.TransactionReceipt(ctx, tx_hash); err == nil {
		return tx_receipt.Status
	} else {
		log.Fatal(err)
		return 0
	}

}

func GetTransactions(block *types.Block, client ethclient.Client, ctx context.Context) {

	for _, tx := range getBlockTx(block) {

		// Read tx receipt
		tx_receipt := checkTxReceipt(client, ctx, tx.Hash())
		_, isPending := getTx(client, ctx, tx.Hash())

		if tx_receipt == 1 {
			fmt.Println()
			fmt.Println("hash:", tx.Hash().Hex())
			fmt.Println("value:", tx.Value().String())
			fmt.Println("gas:", tx.Gas())
			fmt.Println("GasPrice:", tx.GasPrice().Uint64())
			fmt.Println("nonce:", tx.Nonce())
			// fmt.Println("Data:", tx.Data())
			fmt.Println("to:", tx.To().Hex())

			fmt.Println("tx status:", tx_receipt)
			fmt.Println("tx isPending:", isPending)
		}

	}
}
