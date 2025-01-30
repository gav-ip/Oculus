package base

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MagicEdenClient struct {
	PrivateKey        *ecdsa.PrivateKey
	Contract          Magiceden
	EthClient         *ethclient.Client
	ContractAddress   common.Address
	TransactionSigner *bind.TransactOpts
}
