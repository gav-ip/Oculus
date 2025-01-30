package arbitrum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (kc *KingdomlyClient) QuoteFees(amountToMint int) (*big.Int, error) {
	// Set up call options
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	// Inputs for the QuoteBatchMint function
	mintId := big.NewInt(0)
	amount := big.NewInt(int64(amountToMint))

	// Call the function
	quote, err := kc.Contract.QuoteBatchMint(callOpts, mintId, amount)
	if err != nil {
		return nil, err
	}

	return quote.TotalCostWithFee, nil
}
