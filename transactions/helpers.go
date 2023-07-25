package transactions

import (
	"context"
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gasOpts struct {
	Data []byte
	To   common.Address
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

func EstimateGas(client ethclient.Client, ctx context.Context, opts ...gasOpts) uint64 {
	if len(opts) > 0 {
		data := opts[0].Data
		gas_limit, err := client.EstimateGas(ctx, ethereum.CallMsg{
			To:   &opts[0].To,
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		return gas_limit
	}

	gas, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return gas.Uint64()
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
