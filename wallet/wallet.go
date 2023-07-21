package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func createWallet() (string, string) {
	// Create private key
	_privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(_privateKey)
	privateKey := hexutil.Encode(privateKeyBytes)[2:]

	_ = privateKey

	// Create public key
	_publicKey := _privateKey.Public()
	publicKeyECDSA, ok := _publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("Error Casting public key to ECSDA")
	}
	publicKey := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKey, publicKey
}

func PreviewWallet(acc_addr common.Address, client ethclient.Client) {
	private_key, public_key := createWallet()
	fmt.Printf("Public Key: %s\n", public_key)
	fmt.Printf("Private Key: %s\n", private_key)

	fmt.Println()
}
