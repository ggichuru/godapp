package transactions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SetupTransfer(client ethclient.Client, ctx context.Context) {
	private_key := LoadPrivateKey()
	public_key := ExtractPublicKey()
	fmt.Println(public_key)

	nonce := SetNonce(client, ctx, public_key)
	fmt.Println("nonce:", nonce)

	gasLimit := uint64(21000)

	gasPrice := EstimateGas(client, ctx)
	fmt.Println("suggested_gas:", gasPrice)

	value := big.NewInt(10000000000000000) // in wei (0.01 eth)

	to_address := common.HexToAddress("0x3F92A2952746be63f8E22D58997A9A56c95ed2D1")

	// Generate unsigned tx
	tx := types.NewTransaction(nonce, to_address, value, gasLimit, big.NewInt(int64(gasLimit)), nil)

	// Sign transaction
	signed_tx := SignTransaction(client, ctx, tx, private_key)
	fmt.Println("trascation_has:", signed_tx.Hash().Hex())
}
