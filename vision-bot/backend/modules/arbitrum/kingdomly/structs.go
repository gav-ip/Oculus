package arbitrum

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type KingdomlyClient struct {
	PrivateKey        *ecdsa.PrivateKey
	Contract          Kingdomly
	EthClient         *ethclient.Client
	ContractAddress   common.Address
	TransactionSigner *bind.TransactOpts
}
