package abstract

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (mc *MagicEdenClient) GetStageInfo() (MintStageInfo, error) {
	// Create call options
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	// Get current stage number depending on the time
	timestamp := time.Now().Unix()
	currentStage, err := mc.Contract.GetActiveStageFromTimestamp(callOpts, big.NewInt(timestamp))
	if err != nil {
		return MintStageInfo{}, err
	}

	// Get the stage info
	info, _, _, err := mc.Contract.GetStageInfo(callOpts, currentStage)
	if err != nil {
		return MintStageInfo{}, err
	}

	return info, nil
}
