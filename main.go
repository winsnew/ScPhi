package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	rpcUrl          = "https://goerli.infura.io/v3/YOUR_INFURA_PROJECT_ID"
	contractAddress = ""
)

func main() {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to Ethereum Node")
	}

	defer client.Close()
	fmt.Println("Connected to Ethereum Node")

	walletJSON := `{}`
	password := ""
	key, err := keystore.DecryptKey([]byte(walletJSON), password)

	if err != nil {
		log.Fatalf("Failed to load wallet: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, big.NewInt(5))
	if err != nil {
		log.Fatalf("Failed to authorized transactor: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	instance, err := bindings.NewSimpleContract(contractAddr, client)
	if err != nil {
		log.Fatalf("Failed to load smart contract: %v", err)
	}

	message, err := instance.Message(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve message: %v", err)
	}
	fmt.Println("Message: ", message)

	tx, err := instance.SetMessage(auth, "Test")
	if err != nil {
		log.Fatalf("Failed to update message: %v", err)
	}
	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
}
