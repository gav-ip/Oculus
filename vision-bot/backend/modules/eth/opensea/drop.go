package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (oc *OpenseaClient) GetPublicDrop() (PublicDrop, error) {
	// Create call options
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	publicDrop, err := oc.Contract.GetPublicDrop(callOpts, oc.ContractAddress)
	if err != nil {
		return PublicDrop{}, err
	}

	return publicDrop, nil
}
