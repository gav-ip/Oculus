package base

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CHAIN_ID = 8453
)

func NewMagicEdenClient(privateKeyHex, contractAddress string) (*MagicEdenClient, error) {
	// Make the ETH client
	endpoint := os.Getenv("BASE_ENDPOINT")
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum Mainnet network: %v", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Parse contract address
	address := common.HexToAddress(contractAddress)

	// Initalize the instance
	instance, err := NewMagiceden(address, client)
	if err != nil {
		return nil, err
	}

	// Create the transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(CHAIN_ID))
	if err != nil {
		return nil, err
	}
	auth.Nonce = nil    // Use latest nonce
	auth.GasPrice = nil // Use suggested gas price
	auth.GasLimit = 0   // Use estimate gas limit

	// Return the MagicEden client
	return &MagicEdenClient{
		Contract:          *instance,
		EthClient:         client,
		ContractAddress:   address,
		PrivateKey:        privateKey,
		TransactionSigner: auth,
	}, nil
}
