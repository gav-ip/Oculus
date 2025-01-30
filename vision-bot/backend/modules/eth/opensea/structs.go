package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type OpenseaClient struct {
	PrivateKey             *ecdsa.PrivateKey
	SeaDropContractAddress common.Address
	Contract               Opensea
	EthClient              *ethclient.Client
	ContractAddress        common.Address
	TransactionSigner      *bind.TransactOpts
}
