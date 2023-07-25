package transactions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func TransferErc20(client ethclient.Client, ctx context.Context, from_address common.Address, to_address common.Address, amount int) {
	token := common.HexToAddress("0x1D194b8cc47f8dE0B89A82b52EE996970a0D2279")

	// Define the transfer signature for ERC20
	abi := []byte("transfer(address,uint256)")

	// Generate MethodID
	hash := sha3.NewLegacyKeccak256()
	hash.Write(abi)
	// get the first 4 bytes of the resulting has
	methodID := hash.Sum(nil)[:4]

	// left pad the account address
	padded_addr := common.LeftPadBytes(to_address.Bytes(), 32)

	// set the value tokens to send as *big.Int
	amnt_wei := fmt.Sprint(ToWei(nil, float64(amount), 18))
	fmt.Println(amnt_wei)
	_amount := new(big.Int)
	_amount.SetString(amnt_wei, 10)

	// left pad the amount
	padded_amnt := common.LeftPadBytes(_amount.Bytes(), 32)

	// concatenate the methodID, padded (addr, amnt) into a byte slice
	var data []byte
	data = append(data, methodID...)
	data = append(data, padded_addr...)
	data = append(data, padded_amnt...)

	// compute gas limit
	// gas_limit := EstimateGas(client, ctx, gasOpts{Data: data, To: token})
	gas_limit := uint64(200000)

	// compute gas price
	gas_price := big.NewInt(int64(EstimateGas(client, ctx)))

	// get nonce
	nonce := SetNonce(client, ctx, from_address)

	// create transaction
	// tx := types.NewTransaction(nonce, token, big.NewInt(0), gas_limit, gas_price, data)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gas_price,
		Gas:      gas_limit,
		To:       &token,
		Value:    nil,
		Data:     data,
		V:        nil,
		R:        nil,
		S:        nil,
	})

	// get PK
	private_key := LoadPrivateKey()

	sign_tx := SignTransaction(client, ctx, tx, private_key)

	fmt.Printf("Transfering %d WK1 Tokens to %v. \n Hash: %s ", amount, to_address, sign_tx.Hash().Hex())
}
