package eth

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (mc *MagicEdenClient) GetStageInfo() (IERC721MMintStageInfo, error) {
	// Create call options
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	// Get current stage number depending on the time
	timestamp := uint64(time.Now().Unix())
	currentStage, err := mc.Contract.GetActiveStageFromTimestamp(callOpts, timestamp)
	if err != nil {
		return IERC721MMintStageInfo{}, err
	}

	// Get the stage info
	info, _, _, err := mc.Contract.GetStageInfo(callOpts, currentStage)
	if err != nil {
		return IERC721MMintStageInfo{}, err
	}

	return info, err
}
