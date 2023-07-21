package wallet

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeystore() {
	ks := keystore.NewKeyStore("./_cfg/users", keystore.StandardScryptN, keystore.StandardScryptP)

	password := "secret"

	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	acc_addr := account.Address.Hex()
	fmt.Println(acc_addr)
}

func ImportKeystore() {
	file := "./_cfg/users/UTC--2023-07-21T01-48-40.555564000Z--2fb0ea17b5813c67075a0bb174bea39cae4b08ea"

	ks := keystore.NewKeyStore("./_cfg/users", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	acc_addr := account.Address.Hex()
	fmt.Println(acc_addr)

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
