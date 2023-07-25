package transactions

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
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

// ToWei decimals to wei
func ToWei(iamount interface{}, amnt float64, decimals int) *big.Int {
	amount := decimal.NewFromFloat(amnt)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
