package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SEADROP_CONTRACT_ADDR = "0x00005EA00Ac477B1030CE78506496e8C2dE24bf5"
)

func NewOpenseaClient(client *ethclient.Client, privateKeyHex, contractAddress string) (*OpenseaClient, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Parse the contract address
	address := common.HexToAddress(SEADROP_CONTRACT_ADDR)

	// Initalize the instance
	instance, err := NewOpensea(address, client)
	if err != nil {
		return nil, err
	}

	// Create the transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // Replace privateKey and chainID
	if err != nil {
		return nil, err
	}
	auth.Nonce = nil    // Use latest nonce
	auth.GasPrice = nil // Use suggested gas price
	auth.GasLimit = 0   // Use estimate gas limit

	// Return the Opensea Client
	return &OpenseaClient{
		SeaDropContractAddress: address,
		Contract:               *instance,
		EthClient:              client,
		ContractAddress:        common.HexToAddress(contractAddress),
		PrivateKey:             privateKey,
		TransactionSigner:      auth,
	}, nil
}
