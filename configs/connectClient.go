package configs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectEthClient() ethclient.Client {
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

func LoadPrivateKey() *ecdsa.PrivateKey {
	private_key, err := crypto.HexToECDSA(os.Getenv("ACCOUNT_PK"))
	if err != nil {
		log.Fatal(err)
	}
	return private_key
}

func ExtractPublicKey() common.Address {
	// extract public key from private key
	public_key := LoadPrivateKey().Public()
	publicKeyECSDA, ok := public_key.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECSDA")
	}

	return crypto.PubkeyToAddress(*publicKeyECSDA)
}

func SetNonce(client ethclient.Client, ctx context.Context, from_addr common.Address) uint64 {
	nonce, err := client.PendingNonceAt(ctx, from_addr)
	if err != nil {
		log.Fatal(err)
	}
	return nonce
}

func SuggestGas(client ethclient.Client, ctx context.Context) *big.Int {
	gas, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return gas
}

func SignTransaction(client ethclient.Client, ctx context.Context, tx *types.Transaction, private_key *ecdsa.PrivateKey) *types.Transaction {
	chainId, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	signed_tx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), private_key)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(ctx, signed_tx)
	if err != nil {
		log.Fatal(err)
	}

	return signed_tx
}
