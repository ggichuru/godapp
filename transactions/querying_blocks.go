package transactions

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getCurrentBlock(client ethclient.Client, ctx context.Context) int64 {
	// call header by number  by passing nill as the number
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return header.Number.Int64()

}

func getBlock(client ethclient.Client, ctx context.Context, block_number int) *types.Block {
	// call client.BlockByNumber
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(block_number)))
	if err != nil {
		log.Fatal(err)
	}

	return block
}

func TxCount(client ethclient.Client, ctx context.Context, block_hash common.Hash) uint {
	count, err := client.TransactionCount(ctx, block_hash)
	if err != nil {
		log.Fatal(err)
	}

	return count
}

func SimulateBlock(client ethclient.Client, ctx context.Context) *types.Block {
	latest_block := getCurrentBlock(client, ctx)
	fmt.Println("Latest Block:", latest_block)

	block := getBlock(client, ctx, int(latest_block))

	return block
}
