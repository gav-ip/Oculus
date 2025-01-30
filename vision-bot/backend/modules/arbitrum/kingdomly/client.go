package arbitrum

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
	CHAIN_ID = 42161
)

func NewKingdomlyClient(privateKeyHex, contractAddress string) (*KingdomlyClient, error) {
	// Make the ETH client
	endpoint := os.Getenv("ARB_ENDPOINT")
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the arbitrum network: %v", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Parse the contract address
	address := common.HexToAddress(contractAddress)

	// Initalize the instance
	instance, err := NewKingdomly(address, client)
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

	// Return the Kingdomly Client
	return &KingdomlyClient{
		Contract:          *instance,
		EthClient:         client,
		ContractAddress:   address,
		PrivateKey:        privateKey,
		TransactionSigner: auth,
	}, nil
}
