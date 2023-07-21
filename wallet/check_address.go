package wallet

import (
	"context"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsEvmAddr(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	return re.MatchString(address)
}

func IsContract(client ethclient.Client, ctx context.Context, address common.Address) bool {
	// an address is a smart contract if there's bytecode
	bytecode, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		log.Fatal(err)
	}

	return len(bytecode) > 0
}
