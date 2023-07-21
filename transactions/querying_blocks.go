package transactions

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetCurrentBlock(client ethclient.Client, ctx context.Context) string {
	// call header by number  by passing nill as the number
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return header.Number.String()

}

func GetBlock(client ethclient.Client, ctx context.Context, blocknumber int) *types.Block {
	// call client.BlockByNumber
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blocknumber)))
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
