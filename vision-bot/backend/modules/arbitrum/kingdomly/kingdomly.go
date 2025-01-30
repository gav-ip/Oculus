// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DaBrainersBaseVariables is an auto generated low-level Go binding around an user-defined struct.
type DaBrainersBaseVariables struct {
	Name               string
	Symbol             string
	OwnerPayoutAddress common.Address
	InitialBaseURI     string
	MaxSupply          *big.Int
}

// KingdomlyMetaData contains all meta data concerning the Kingdomly contract.
var KingdomlyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerPayoutAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"initialBaseURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"internalType\":\"structDa_Brainers.BaseVariables\",\"name\":\"_baseVariables\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"_maxMintPerWallet\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_maxSupplyPerMintGroup\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mintPrice\",\"type\":\"uint256[]\"},{\"internalType\":\"uint96\",\"name\":\"_royaltyPercentage\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_kingdomlyAdmin\",\"type\":\"address\"},{\"internalType\":\"contractKingdomlyFeeContract\",\"name\":\"_kingdomlyFeeContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidDefaultRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidDefaultRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidTokenRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidTokenRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxMintGroupSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowed\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxPerWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowed\",\"type\":\"uint256\"}],\"name\":\"ExceedsMintQuota\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"}],\"name\":\"InsufficientEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKingdomlyFeeContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"MintGroupDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"MintGroupInactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintInactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCompatibleWithSpotMints\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"NotInPresale\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialMintExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialUpToTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SpotMintTokenIdTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"}],\"name\":\"KingdomlyFeeContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxMintPerWallet\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintGroupId\",\"type\":\"uint256\"}],\"name\":\"MaxMintPerWalletChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintGroupId\",\"type\":\"uint256\"}],\"name\":\"PreSaleMintScheduledStartTimestampChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintGroupId\",\"type\":\"uint256\"}],\"name\":\"PreSaleMintStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"SalePriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hotWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"TokensDelegateMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"TokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeMintGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"airdropNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCharge\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"batchMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCostWithFee\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterAddress\",\"type\":\"address\"}],\"name\":\"canDelegateMintCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"minterAddress\",\"type\":\"address\"}],\"name\":\"canMintCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"changeMintStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newmintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"changeSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPendingBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"checkPendingBalanceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractMintLive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractPresaleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"delegatedMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCostWithFee\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKingdomlyFeeContract\",\"outputs\":[{\"internalType\":\"contractKingdomlyFeeContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kingdomlyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kingdomlyFeeContract\",\"outputs\":[{\"internalType\":\"contractKingdomlyFeeContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maxMintPerWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maxSupplyPerMintGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintGroupMints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintLive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintQuotas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerPayoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"presaleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"presaleScheduledStartTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"quoteAirdropFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAirdropCostWithFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"quoteBatchMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCostWithFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"schedulePresaleMintStart\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduledMintLiveTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxMintPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintGroupId\",\"type\":\"uint256\"}],\"name\":\"setMaxMintPerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setMintLiveTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addressToAdd\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_mintQuotas\",\"type\":\"uint256[]\"}],\"name\":\"setMintQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractKingdomlyFeeContract\",\"name\":\"_kingdomlyFeeContract\",\"type\":\"address\"}],\"name\":\"setNewKingdomlyFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setNewMaxPerMintGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"presaleStatus\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mintId\",\"type\":\"uint256\"}],\"name\":\"stopOrStartpresaleMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threeDollarsEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threeDollarsInCents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMintFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KingdomlyABI is the input ABI used to generate the binding from.
// Deprecated: Use KingdomlyMetaData.ABI instead.
var KingdomlyABI = KingdomlyMetaData.ABI

// Kingdomly is an auto generated Go binding around an Ethereum contract.
type Kingdomly struct {
	KingdomlyCaller     // Read-only binding to the contract
	KingdomlyTransactor // Write-only binding to the contract
	KingdomlyFilterer   // Log filterer for contract events
}

// KingdomlyCaller is an auto generated read-only Go binding around an Ethereum contract.
type KingdomlyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomlyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KingdomlyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomlyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KingdomlyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KingdomlySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KingdomlySession struct {
	Contract     *Kingdomly        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KingdomlyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KingdomlyCallerSession struct {
	Contract *KingdomlyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KingdomlyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KingdomlyTransactorSession struct {
	Contract     *KingdomlyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KingdomlyRaw is an auto generated low-level Go binding around an Ethereum contract.
type KingdomlyRaw struct {
	Contract *Kingdomly // Generic contract binding to access the raw methods on
}

// KingdomlyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KingdomlyCallerRaw struct {
	Contract *KingdomlyCaller // Generic read-only contract binding to access the raw methods on
}

// KingdomlyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KingdomlyTransactorRaw struct {
	Contract *KingdomlyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKingdomly creates a new instance of Kingdomly, bound to a specific deployed contract.
func NewKingdomly(address common.Address, backend bind.ContractBackend) (*Kingdomly, error) {
	contract, err := bindKingdomly(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kingdomly{KingdomlyCaller: KingdomlyCaller{contract: contract}, KingdomlyTransactor: KingdomlyTransactor{contract: contract}, KingdomlyFilterer: KingdomlyFilterer{contract: contract}}, nil
}

// NewKingdomlyCaller creates a new read-only instance of Kingdomly, bound to a specific deployed contract.
func NewKingdomlyCaller(address common.Address, caller bind.ContractCaller) (*KingdomlyCaller, error) {
	contract, err := bindKingdomly(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KingdomlyCaller{contract: contract}, nil
}

// NewKingdomlyTransactor creates a new write-only instance of Kingdomly, bound to a specific deployed contract.
func NewKingdomlyTransactor(address common.Address, transactor bind.ContractTransactor) (*KingdomlyTransactor, error) {
	contract, err := bindKingdomly(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KingdomlyTransactor{contract: contract}, nil
}

// NewKingdomlyFilterer creates a new log filterer instance of Kingdomly, bound to a specific deployed contract.
func NewKingdomlyFilterer(address common.Address, filterer bind.ContractFilterer) (*KingdomlyFilterer, error) {
	contract, err := bindKingdomly(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KingdomlyFilterer{contract: contract}, nil
}

// bindKingdomly binds a generic wrapper to an already deployed contract.
func bindKingdomly(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KingdomlyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kingdomly *KingdomlyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kingdomly.Contract.KingdomlyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kingdomly *KingdomlyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kingdomly.Contract.KingdomlyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kingdomly *KingdomlyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kingdomly.Contract.KingdomlyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kingdomly *KingdomlyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kingdomly.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kingdomly *KingdomlyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kingdomly.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kingdomly *KingdomlyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kingdomly.Contract.contract.Transact(opts, method, params...)
}

// ActiveMintGroups is a free data retrieval call binding the contract method 0x483f0a82.
//
// Solidity: function activeMintGroups(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) ActiveMintGroups(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "activeMintGroups", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveMintGroups is a free data retrieval call binding the contract method 0x483f0a82.
//
// Solidity: function activeMintGroups(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlySession) ActiveMintGroups(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.ActiveMintGroups(&_Kingdomly.CallOpts, arg0)
}

// ActiveMintGroups is a free data retrieval call binding the contract method 0x483f0a82.
//
// Solidity: function activeMintGroups(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) ActiveMintGroups(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.ActiveMintGroups(&_Kingdomly.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kingdomly *KingdomlySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Kingdomly.Contract.BalanceOf(&_Kingdomly.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Kingdomly.Contract.BalanceOf(&_Kingdomly.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Kingdomly *KingdomlyCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Kingdomly *KingdomlySession) BaseURI() (string, error) {
	return _Kingdomly.Contract.BaseURI(&_Kingdomly.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Kingdomly *KingdomlyCallerSession) BaseURI() (string, error) {
	return _Kingdomly.Contract.BaseURI(&_Kingdomly.CallOpts)
}

// CanDelegateMintCheck is a free data retrieval call binding the contract method 0xf72171f1.
//
// Solidity: function canDelegateMintCheck(uint256 amount, uint256 mintId, address vault, address minterAddress) view returns(bool)
func (_Kingdomly *KingdomlyCaller) CanDelegateMintCheck(opts *bind.CallOpts, amount *big.Int, mintId *big.Int, vault common.Address, minterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "canDelegateMintCheck", amount, mintId, vault, minterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDelegateMintCheck is a free data retrieval call binding the contract method 0xf72171f1.
//
// Solidity: function canDelegateMintCheck(uint256 amount, uint256 mintId, address vault, address minterAddress) view returns(bool)
func (_Kingdomly *KingdomlySession) CanDelegateMintCheck(amount *big.Int, mintId *big.Int, vault common.Address, minterAddress common.Address) (bool, error) {
	return _Kingdomly.Contract.CanDelegateMintCheck(&_Kingdomly.CallOpts, amount, mintId, vault, minterAddress)
}

// CanDelegateMintCheck is a free data retrieval call binding the contract method 0xf72171f1.
//
// Solidity: function canDelegateMintCheck(uint256 amount, uint256 mintId, address vault, address minterAddress) view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) CanDelegateMintCheck(amount *big.Int, mintId *big.Int, vault common.Address, minterAddress common.Address) (bool, error) {
	return _Kingdomly.Contract.CanDelegateMintCheck(&_Kingdomly.CallOpts, amount, mintId, vault, minterAddress)
}

// CanMintCheck is a free data retrieval call binding the contract method 0xf34657ec.
//
// Solidity: function canMintCheck(uint256 amount, uint256 mintId, address minterAddress) view returns(bool)
func (_Kingdomly *KingdomlyCaller) CanMintCheck(opts *bind.CallOpts, amount *big.Int, mintId *big.Int, minterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "canMintCheck", amount, mintId, minterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanMintCheck is a free data retrieval call binding the contract method 0xf34657ec.
//
// Solidity: function canMintCheck(uint256 amount, uint256 mintId, address minterAddress) view returns(bool)
func (_Kingdomly *KingdomlySession) CanMintCheck(amount *big.Int, mintId *big.Int, minterAddress common.Address) (bool, error) {
	return _Kingdomly.Contract.CanMintCheck(&_Kingdomly.CallOpts, amount, mintId, minterAddress)
}

// CanMintCheck is a free data retrieval call binding the contract method 0xf34657ec.
//
// Solidity: function canMintCheck(uint256 amount, uint256 mintId, address minterAddress) view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) CanMintCheck(amount *big.Int, mintId *big.Int, minterAddress common.Address) (bool, error) {
	return _Kingdomly.Contract.CanMintCheck(&_Kingdomly.CallOpts, amount, mintId, minterAddress)
}

// CheckPendingBalance is a free data retrieval call binding the contract method 0xfea414b6.
//
// Solidity: function checkPendingBalance() view returns(uint256)
func (_Kingdomly *KingdomlyCaller) CheckPendingBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "checkPendingBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckPendingBalance is a free data retrieval call binding the contract method 0xfea414b6.
//
// Solidity: function checkPendingBalance() view returns(uint256)
func (_Kingdomly *KingdomlySession) CheckPendingBalance() (*big.Int, error) {
	return _Kingdomly.Contract.CheckPendingBalance(&_Kingdomly.CallOpts)
}

// CheckPendingBalance is a free data retrieval call binding the contract method 0xfea414b6.
//
// Solidity: function checkPendingBalance() view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) CheckPendingBalance() (*big.Int, error) {
	return _Kingdomly.Contract.CheckPendingBalance(&_Kingdomly.CallOpts)
}

// CheckPendingBalanceFor is a free data retrieval call binding the contract method 0x6e75e2e2.
//
// Solidity: function checkPendingBalanceFor(address user) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) CheckPendingBalanceFor(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "checkPendingBalanceFor", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckPendingBalanceFor is a free data retrieval call binding the contract method 0x6e75e2e2.
//
// Solidity: function checkPendingBalanceFor(address user) view returns(uint256)
func (_Kingdomly *KingdomlySession) CheckPendingBalanceFor(user common.Address) (*big.Int, error) {
	return _Kingdomly.Contract.CheckPendingBalanceFor(&_Kingdomly.CallOpts, user)
}

// CheckPendingBalanceFor is a free data retrieval call binding the contract method 0x6e75e2e2.
//
// Solidity: function checkPendingBalanceFor(address user) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) CheckPendingBalanceFor(user common.Address) (*big.Int, error) {
	return _Kingdomly.Contract.CheckPendingBalanceFor(&_Kingdomly.CallOpts, user)
}

// ContractMintLive is a free data retrieval call binding the contract method 0xe690f9eb.
//
// Solidity: function contractMintLive() view returns(bool)
func (_Kingdomly *KingdomlyCaller) ContractMintLive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "contractMintLive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractMintLive is a free data retrieval call binding the contract method 0xe690f9eb.
//
// Solidity: function contractMintLive() view returns(bool)
func (_Kingdomly *KingdomlySession) ContractMintLive() (bool, error) {
	return _Kingdomly.Contract.ContractMintLive(&_Kingdomly.CallOpts)
}

// ContractMintLive is a free data retrieval call binding the contract method 0xe690f9eb.
//
// Solidity: function contractMintLive() view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) ContractMintLive() (bool, error) {
	return _Kingdomly.Contract.ContractMintLive(&_Kingdomly.CallOpts)
}

// ContractPresaleActive is a free data retrieval call binding the contract method 0xdb0dbb71.
//
// Solidity: function contractPresaleActive(uint256 ) view returns(bool)
func (_Kingdomly *KingdomlyCaller) ContractPresaleActive(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "contractPresaleActive", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractPresaleActive is a free data retrieval call binding the contract method 0xdb0dbb71.
//
// Solidity: function contractPresaleActive(uint256 ) view returns(bool)
func (_Kingdomly *KingdomlySession) ContractPresaleActive(arg0 *big.Int) (bool, error) {
	return _Kingdomly.Contract.ContractPresaleActive(&_Kingdomly.CallOpts, arg0)
}

// ContractPresaleActive is a free data retrieval call binding the contract method 0xdb0dbb71.
//
// Solidity: function contractPresaleActive(uint256 ) view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) ContractPresaleActive(arg0 *big.Int) (bool, error) {
	return _Kingdomly.Contract.ContractPresaleActive(&_Kingdomly.CallOpts, arg0)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Kingdomly *KingdomlyCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Kingdomly *KingdomlySession) FeeAddress() (common.Address, error) {
	return _Kingdomly.Contract.FeeAddress(&_Kingdomly.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Kingdomly *KingdomlyCallerSession) FeeAddress() (common.Address, error) {
	return _Kingdomly.Contract.FeeAddress(&_Kingdomly.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kingdomly *KingdomlyCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kingdomly *KingdomlySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Kingdomly.Contract.GetApproved(&_Kingdomly.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kingdomly *KingdomlyCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Kingdomly.Contract.GetApproved(&_Kingdomly.CallOpts, tokenId)
}

// GetKingdomlyFeeContract is a free data retrieval call binding the contract method 0x93a24841.
//
// Solidity: function getKingdomlyFeeContract() view returns(address)
func (_Kingdomly *KingdomlyCaller) GetKingdomlyFeeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "getKingdomlyFeeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKingdomlyFeeContract is a free data retrieval call binding the contract method 0x93a24841.
//
// Solidity: function getKingdomlyFeeContract() view returns(address)
func (_Kingdomly *KingdomlySession) GetKingdomlyFeeContract() (common.Address, error) {
	return _Kingdomly.Contract.GetKingdomlyFeeContract(&_Kingdomly.CallOpts)
}

// GetKingdomlyFeeContract is a free data retrieval call binding the contract method 0x93a24841.
//
// Solidity: function getKingdomlyFeeContract() view returns(address)
func (_Kingdomly *KingdomlyCallerSession) GetKingdomlyFeeContract() (common.Address, error) {
	return _Kingdomly.Contract.GetKingdomlyFeeContract(&_Kingdomly.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Kingdomly *KingdomlyCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Kingdomly *KingdomlySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Kingdomly.Contract.IsApprovedForAll(&_Kingdomly.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Kingdomly.Contract.IsApprovedForAll(&_Kingdomly.CallOpts, owner, operator)
}

// KingdomlyAdmin is a free data retrieval call binding the contract method 0x63691c93.
//
// Solidity: function kingdomlyAdmin() view returns(address)
func (_Kingdomly *KingdomlyCaller) KingdomlyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "kingdomlyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KingdomlyAdmin is a free data retrieval call binding the contract method 0x63691c93.
//
// Solidity: function kingdomlyAdmin() view returns(address)
func (_Kingdomly *KingdomlySession) KingdomlyAdmin() (common.Address, error) {
	return _Kingdomly.Contract.KingdomlyAdmin(&_Kingdomly.CallOpts)
}

// KingdomlyAdmin is a free data retrieval call binding the contract method 0x63691c93.
//
// Solidity: function kingdomlyAdmin() view returns(address)
func (_Kingdomly *KingdomlyCallerSession) KingdomlyAdmin() (common.Address, error) {
	return _Kingdomly.Contract.KingdomlyAdmin(&_Kingdomly.CallOpts)
}

// KingdomlyFeeContract is a free data retrieval call binding the contract method 0xefe82328.
//
// Solidity: function kingdomlyFeeContract() view returns(address)
func (_Kingdomly *KingdomlyCaller) KingdomlyFeeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "kingdomlyFeeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KingdomlyFeeContract is a free data retrieval call binding the contract method 0xefe82328.
//
// Solidity: function kingdomlyFeeContract() view returns(address)
func (_Kingdomly *KingdomlySession) KingdomlyFeeContract() (common.Address, error) {
	return _Kingdomly.Contract.KingdomlyFeeContract(&_Kingdomly.CallOpts)
}

// KingdomlyFeeContract is a free data retrieval call binding the contract method 0xefe82328.
//
// Solidity: function kingdomlyFeeContract() view returns(address)
func (_Kingdomly *KingdomlyCallerSession) KingdomlyFeeContract() (common.Address, error) {
	return _Kingdomly.Contract.KingdomlyFeeContract(&_Kingdomly.CallOpts)
}

// MaxMintPerWallet is a free data retrieval call binding the contract method 0x16da3bc6.
//
// Solidity: function maxMintPerWallet(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) MaxMintPerWallet(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "maxMintPerWallet", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMintPerWallet is a free data retrieval call binding the contract method 0x16da3bc6.
//
// Solidity: function maxMintPerWallet(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlySession) MaxMintPerWallet(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MaxMintPerWallet(&_Kingdomly.CallOpts, arg0)
}

// MaxMintPerWallet is a free data retrieval call binding the contract method 0x16da3bc6.
//
// Solidity: function maxMintPerWallet(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) MaxMintPerWallet(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MaxMintPerWallet(&_Kingdomly.CallOpts, arg0)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Kingdomly *KingdomlyCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Kingdomly *KingdomlySession) MaxSupply() (*big.Int, error) {
	return _Kingdomly.Contract.MaxSupply(&_Kingdomly.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) MaxSupply() (*big.Int, error) {
	return _Kingdomly.Contract.MaxSupply(&_Kingdomly.CallOpts)
}

// MaxSupplyPerMintGroup is a free data retrieval call binding the contract method 0x24a663c3.
//
// Solidity: function maxSupplyPerMintGroup(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) MaxSupplyPerMintGroup(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "maxSupplyPerMintGroup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupplyPerMintGroup is a free data retrieval call binding the contract method 0x24a663c3.
//
// Solidity: function maxSupplyPerMintGroup(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlySession) MaxSupplyPerMintGroup(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MaxSupplyPerMintGroup(&_Kingdomly.CallOpts, arg0)
}

// MaxSupplyPerMintGroup is a free data retrieval call binding the contract method 0x24a663c3.
//
// Solidity: function maxSupplyPerMintGroup(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) MaxSupplyPerMintGroup(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MaxSupplyPerMintGroup(&_Kingdomly.CallOpts, arg0)
}

// MintGroupMints is a free data retrieval call binding the contract method 0xe5fd1145.
//
// Solidity: function mintGroupMints(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) MintGroupMints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "mintGroupMints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintGroupMints is a free data retrieval call binding the contract method 0xe5fd1145.
//
// Solidity: function mintGroupMints(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlySession) MintGroupMints(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MintGroupMints(&_Kingdomly.CallOpts, arg0)
}

// MintGroupMints is a free data retrieval call binding the contract method 0xe5fd1145.
//
// Solidity: function mintGroupMints(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) MintGroupMints(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MintGroupMints(&_Kingdomly.CallOpts, arg0)
}

// MintLive is a free data retrieval call binding the contract method 0xe8656fcc.
//
// Solidity: function mintLive() view returns(bool)
func (_Kingdomly *KingdomlyCaller) MintLive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "mintLive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintLive is a free data retrieval call binding the contract method 0xe8656fcc.
//
// Solidity: function mintLive() view returns(bool)
func (_Kingdomly *KingdomlySession) MintLive() (bool, error) {
	return _Kingdomly.Contract.MintLive(&_Kingdomly.CallOpts)
}

// MintLive is a free data retrieval call binding the contract method 0xe8656fcc.
//
// Solidity: function mintLive() view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) MintLive() (bool, error) {
	return _Kingdomly.Contract.MintLive(&_Kingdomly.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0xe6a72acf.
//
// Solidity: function mintPrice(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) MintPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "mintPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0xe6a72acf.
//
// Solidity: function mintPrice(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlySession) MintPrice(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MintPrice(&_Kingdomly.CallOpts, arg0)
}

// MintPrice is a free data retrieval call binding the contract method 0xe6a72acf.
//
// Solidity: function mintPrice(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) MintPrice(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.MintPrice(&_Kingdomly.CallOpts, arg0)
}

// MintQuotas is a free data retrieval call binding the contract method 0x01d2718e.
//
// Solidity: function mintQuotas(uint256 , address ) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) MintQuotas(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "mintQuotas", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintQuotas is a free data retrieval call binding the contract method 0x01d2718e.
//
// Solidity: function mintQuotas(uint256 , address ) view returns(uint256)
func (_Kingdomly *KingdomlySession) MintQuotas(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Kingdomly.Contract.MintQuotas(&_Kingdomly.CallOpts, arg0, arg1)
}

// MintQuotas is a free data retrieval call binding the contract method 0x01d2718e.
//
// Solidity: function mintQuotas(uint256 , address ) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) MintQuotas(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Kingdomly.Contract.MintQuotas(&_Kingdomly.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kingdomly *KingdomlyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kingdomly *KingdomlySession) Name() (string, error) {
	return _Kingdomly.Contract.Name(&_Kingdomly.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kingdomly *KingdomlyCallerSession) Name() (string, error) {
	return _Kingdomly.Contract.Name(&_Kingdomly.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kingdomly *KingdomlyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kingdomly *KingdomlySession) Owner() (common.Address, error) {
	return _Kingdomly.Contract.Owner(&_Kingdomly.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kingdomly *KingdomlyCallerSession) Owner() (common.Address, error) {
	return _Kingdomly.Contract.Owner(&_Kingdomly.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kingdomly *KingdomlyCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kingdomly *KingdomlySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Kingdomly.Contract.OwnerOf(&_Kingdomly.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kingdomly *KingdomlyCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Kingdomly.Contract.OwnerOf(&_Kingdomly.CallOpts, tokenId)
}

// OwnerPayoutAddress is a free data retrieval call binding the contract method 0xae4e4942.
//
// Solidity: function ownerPayoutAddress() view returns(address)
func (_Kingdomly *KingdomlyCaller) OwnerPayoutAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "ownerPayoutAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerPayoutAddress is a free data retrieval call binding the contract method 0xae4e4942.
//
// Solidity: function ownerPayoutAddress() view returns(address)
func (_Kingdomly *KingdomlySession) OwnerPayoutAddress() (common.Address, error) {
	return _Kingdomly.Contract.OwnerPayoutAddress(&_Kingdomly.CallOpts)
}

// OwnerPayoutAddress is a free data retrieval call binding the contract method 0xae4e4942.
//
// Solidity: function ownerPayoutAddress() view returns(address)
func (_Kingdomly *KingdomlyCallerSession) OwnerPayoutAddress() (common.Address, error) {
	return _Kingdomly.Contract.OwnerPayoutAddress(&_Kingdomly.CallOpts)
}

// PresaleActive is a free data retrieval call binding the contract method 0x4a5bd2fd.
//
// Solidity: function presaleActive(uint256 mintId) view returns(bool)
func (_Kingdomly *KingdomlyCaller) PresaleActive(opts *bind.CallOpts, mintId *big.Int) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "presaleActive", mintId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PresaleActive is a free data retrieval call binding the contract method 0x4a5bd2fd.
//
// Solidity: function presaleActive(uint256 mintId) view returns(bool)
func (_Kingdomly *KingdomlySession) PresaleActive(mintId *big.Int) (bool, error) {
	return _Kingdomly.Contract.PresaleActive(&_Kingdomly.CallOpts, mintId)
}

// PresaleActive is a free data retrieval call binding the contract method 0x4a5bd2fd.
//
// Solidity: function presaleActive(uint256 mintId) view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) PresaleActive(mintId *big.Int) (bool, error) {
	return _Kingdomly.Contract.PresaleActive(&_Kingdomly.CallOpts, mintId)
}

// PresaleScheduledStartTimestamp is a free data retrieval call binding the contract method 0xe6087db0.
//
// Solidity: function presaleScheduledStartTimestamp(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCaller) PresaleScheduledStartTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "presaleScheduledStartTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleScheduledStartTimestamp is a free data retrieval call binding the contract method 0xe6087db0.
//
// Solidity: function presaleScheduledStartTimestamp(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlySession) PresaleScheduledStartTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.PresaleScheduledStartTimestamp(&_Kingdomly.CallOpts, arg0)
}

// PresaleScheduledStartTimestamp is a free data retrieval call binding the contract method 0xe6087db0.
//
// Solidity: function presaleScheduledStartTimestamp(uint256 ) view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) PresaleScheduledStartTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.PresaleScheduledStartTimestamp(&_Kingdomly.CallOpts, arg0)
}

// QuoteAirdropFees is a free data retrieval call binding the contract method 0xb80f8fb4.
//
// Solidity: function quoteAirdropFees(uint256 amount) view returns(uint256 totalAirdropCostWithFee)
func (_Kingdomly *KingdomlyCaller) QuoteAirdropFees(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "quoteAirdropFees", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteAirdropFees is a free data retrieval call binding the contract method 0xb80f8fb4.
//
// Solidity: function quoteAirdropFees(uint256 amount) view returns(uint256 totalAirdropCostWithFee)
func (_Kingdomly *KingdomlySession) QuoteAirdropFees(amount *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.QuoteAirdropFees(&_Kingdomly.CallOpts, amount)
}

// QuoteAirdropFees is a free data retrieval call binding the contract method 0xb80f8fb4.
//
// Solidity: function quoteAirdropFees(uint256 amount) view returns(uint256 totalAirdropCostWithFee)
func (_Kingdomly *KingdomlyCallerSession) QuoteAirdropFees(amount *big.Int) (*big.Int, error) {
	return _Kingdomly.Contract.QuoteAirdropFees(&_Kingdomly.CallOpts, amount)
}

// QuoteBatchMint is a free data retrieval call binding the contract method 0x22536c03.
//
// Solidity: function quoteBatchMint(uint256 mintId, uint256 amount) view returns(uint256 totalCostWithFee, uint256 feeAmount)
func (_Kingdomly *KingdomlyCaller) QuoteBatchMint(opts *bind.CallOpts, mintId *big.Int, amount *big.Int) (struct {
	TotalCostWithFee *big.Int
	FeeAmount        *big.Int
}, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "quoteBatchMint", mintId, amount)

	outstruct := new(struct {
		TotalCostWithFee *big.Int
		FeeAmount        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCostWithFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteBatchMint is a free data retrieval call binding the contract method 0x22536c03.
//
// Solidity: function quoteBatchMint(uint256 mintId, uint256 amount) view returns(uint256 totalCostWithFee, uint256 feeAmount)
func (_Kingdomly *KingdomlySession) QuoteBatchMint(mintId *big.Int, amount *big.Int) (struct {
	TotalCostWithFee *big.Int
	FeeAmount        *big.Int
}, error) {
	return _Kingdomly.Contract.QuoteBatchMint(&_Kingdomly.CallOpts, mintId, amount)
}

// QuoteBatchMint is a free data retrieval call binding the contract method 0x22536c03.
//
// Solidity: function quoteBatchMint(uint256 mintId, uint256 amount) view returns(uint256 totalCostWithFee, uint256 feeAmount)
func (_Kingdomly *KingdomlyCallerSession) QuoteBatchMint(mintId *big.Int, amount *big.Int) (struct {
	TotalCostWithFee *big.Int
	FeeAmount        *big.Int
}, error) {
	return _Kingdomly.Contract.QuoteBatchMint(&_Kingdomly.CallOpts, mintId, amount)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 amount)
func (_Kingdomly *KingdomlyCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 amount)
func (_Kingdomly *KingdomlySession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
}, error) {
	return _Kingdomly.Contract.RoyaltyInfo(&_Kingdomly.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 amount)
func (_Kingdomly *KingdomlyCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
}, error) {
	return _Kingdomly.Contract.RoyaltyInfo(&_Kingdomly.CallOpts, tokenId, salePrice)
}

// ScheduledMintLiveTimestamp is a free data retrieval call binding the contract method 0xf30874ea.
//
// Solidity: function scheduledMintLiveTimestamp() view returns(uint256)
func (_Kingdomly *KingdomlyCaller) ScheduledMintLiveTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "scheduledMintLiveTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScheduledMintLiveTimestamp is a free data retrieval call binding the contract method 0xf30874ea.
//
// Solidity: function scheduledMintLiveTimestamp() view returns(uint256)
func (_Kingdomly *KingdomlySession) ScheduledMintLiveTimestamp() (*big.Int, error) {
	return _Kingdomly.Contract.ScheduledMintLiveTimestamp(&_Kingdomly.CallOpts)
}

// ScheduledMintLiveTimestamp is a free data retrieval call binding the contract method 0xf30874ea.
//
// Solidity: function scheduledMintLiveTimestamp() view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) ScheduledMintLiveTimestamp() (*big.Int, error) {
	return _Kingdomly.Contract.ScheduledMintLiveTimestamp(&_Kingdomly.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kingdomly *KingdomlyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kingdomly *KingdomlySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kingdomly.Contract.SupportsInterface(&_Kingdomly.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kingdomly *KingdomlyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kingdomly.Contract.SupportsInterface(&_Kingdomly.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kingdomly *KingdomlyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kingdomly *KingdomlySession) Symbol() (string, error) {
	return _Kingdomly.Contract.Symbol(&_Kingdomly.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kingdomly *KingdomlyCallerSession) Symbol() (string, error) {
	return _Kingdomly.Contract.Symbol(&_Kingdomly.CallOpts)
}

// ThreeDollarsEth is a free data retrieval call binding the contract method 0xce55c66a.
//
// Solidity: function threeDollarsEth() view returns(uint256)
func (_Kingdomly *KingdomlyCaller) ThreeDollarsEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "threeDollarsEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThreeDollarsEth is a free data retrieval call binding the contract method 0xce55c66a.
//
// Solidity: function threeDollarsEth() view returns(uint256)
func (_Kingdomly *KingdomlySession) ThreeDollarsEth() (*big.Int, error) {
	return _Kingdomly.Contract.ThreeDollarsEth(&_Kingdomly.CallOpts)
}

// ThreeDollarsEth is a free data retrieval call binding the contract method 0xce55c66a.
//
// Solidity: function threeDollarsEth() view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) ThreeDollarsEth() (*big.Int, error) {
	return _Kingdomly.Contract.ThreeDollarsEth(&_Kingdomly.CallOpts)
}

// ThreeDollarsInCents is a free data retrieval call binding the contract method 0x17d791d0.
//
// Solidity: function threeDollarsInCents() view returns(uint256)
func (_Kingdomly *KingdomlyCaller) ThreeDollarsInCents(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "threeDollarsInCents")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThreeDollarsInCents is a free data retrieval call binding the contract method 0x17d791d0.
//
// Solidity: function threeDollarsInCents() view returns(uint256)
func (_Kingdomly *KingdomlySession) ThreeDollarsInCents() (*big.Int, error) {
	return _Kingdomly.Contract.ThreeDollarsInCents(&_Kingdomly.CallOpts)
}

// ThreeDollarsInCents is a free data retrieval call binding the contract method 0x17d791d0.
//
// Solidity: function threeDollarsInCents() view returns(uint256)
func (_Kingdomly *KingdomlyCallerSession) ThreeDollarsInCents() (*big.Int, error) {
	return _Kingdomly.Contract.ThreeDollarsInCents(&_Kingdomly.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kingdomly *KingdomlyCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kingdomly *KingdomlySession) TokenURI(tokenId *big.Int) (string, error) {
	return _Kingdomly.Contract.TokenURI(&_Kingdomly.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kingdomly *KingdomlyCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Kingdomly.Contract.TokenURI(&_Kingdomly.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Kingdomly *KingdomlyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kingdomly.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Kingdomly *KingdomlySession) TotalSupply() (*big.Int, error) {
	return _Kingdomly.Contract.TotalSupply(&_Kingdomly.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Kingdomly *KingdomlyCallerSession) TotalSupply() (*big.Int, error) {
	return _Kingdomly.Contract.TotalSupply(&_Kingdomly.CallOpts)
}

// AirdropNFTs is a paid mutator transaction binding the contract method 0xe213b5f6.
//
// Solidity: function airdropNFTs(address[] recipients, uint256[] amounts) payable returns(uint256 totalCharge)
func (_Kingdomly *KingdomlyTransactor) AirdropNFTs(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "airdropNFTs", recipients, amounts)
}

// AirdropNFTs is a paid mutator transaction binding the contract method 0xe213b5f6.
//
// Solidity: function airdropNFTs(address[] recipients, uint256[] amounts) payable returns(uint256 totalCharge)
func (_Kingdomly *KingdomlySession) AirdropNFTs(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.AirdropNFTs(&_Kingdomly.TransactOpts, recipients, amounts)
}

// AirdropNFTs is a paid mutator transaction binding the contract method 0xe213b5f6.
//
// Solidity: function airdropNFTs(address[] recipients, uint256[] amounts) payable returns(uint256 totalCharge)
func (_Kingdomly *KingdomlyTransactorSession) AirdropNFTs(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.AirdropNFTs(&_Kingdomly.TransactOpts, recipients, amounts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlyTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.Approve(&_Kingdomly.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlyTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.Approve(&_Kingdomly.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x0ed64eff.
//
// Solidity: function batchMint(uint256 amount, uint256 mintId) payable returns(uint256 totalCostWithFee)
func (_Kingdomly *KingdomlyTransactor) BatchMint(opts *bind.TransactOpts, amount *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "batchMint", amount, mintId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x0ed64eff.
//
// Solidity: function batchMint(uint256 amount, uint256 mintId) payable returns(uint256 totalCostWithFee)
func (_Kingdomly *KingdomlySession) BatchMint(amount *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.BatchMint(&_Kingdomly.TransactOpts, amount, mintId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x0ed64eff.
//
// Solidity: function batchMint(uint256 amount, uint256 mintId) payable returns(uint256 totalCostWithFee)
func (_Kingdomly *KingdomlyTransactorSession) BatchMint(amount *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.BatchMint(&_Kingdomly.TransactOpts, amount, mintId)
}

// ChangeMintStatus is a paid mutator transaction binding the contract method 0xa8ddf8f6.
//
// Solidity: function changeMintStatus(bool status) returns()
func (_Kingdomly *KingdomlyTransactor) ChangeMintStatus(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "changeMintStatus", status)
}

// ChangeMintStatus is a paid mutator transaction binding the contract method 0xa8ddf8f6.
//
// Solidity: function changeMintStatus(bool status) returns()
func (_Kingdomly *KingdomlySession) ChangeMintStatus(status bool) (*types.Transaction, error) {
	return _Kingdomly.Contract.ChangeMintStatus(&_Kingdomly.TransactOpts, status)
}

// ChangeMintStatus is a paid mutator transaction binding the contract method 0xa8ddf8f6.
//
// Solidity: function changeMintStatus(bool status) returns()
func (_Kingdomly *KingdomlyTransactorSession) ChangeMintStatus(status bool) (*types.Transaction, error) {
	return _Kingdomly.Contract.ChangeMintStatus(&_Kingdomly.TransactOpts, status)
}

// ChangeSalePrice is a paid mutator transaction binding the contract method 0xc82e474b.
//
// Solidity: function changeSalePrice(uint256 newmintPrice, uint256 mintId) returns()
func (_Kingdomly *KingdomlyTransactor) ChangeSalePrice(opts *bind.TransactOpts, newmintPrice *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "changeSalePrice", newmintPrice, mintId)
}

// ChangeSalePrice is a paid mutator transaction binding the contract method 0xc82e474b.
//
// Solidity: function changeSalePrice(uint256 newmintPrice, uint256 mintId) returns()
func (_Kingdomly *KingdomlySession) ChangeSalePrice(newmintPrice *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.ChangeSalePrice(&_Kingdomly.TransactOpts, newmintPrice, mintId)
}

// ChangeSalePrice is a paid mutator transaction binding the contract method 0xc82e474b.
//
// Solidity: function changeSalePrice(uint256 newmintPrice, uint256 mintId) returns()
func (_Kingdomly *KingdomlyTransactorSession) ChangeSalePrice(newmintPrice *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.ChangeSalePrice(&_Kingdomly.TransactOpts, newmintPrice, mintId)
}

// DelegatedMint is a paid mutator transaction binding the contract method 0xebf39a41.
//
// Solidity: function delegatedMint(uint256 amount, uint256 mintId, address vault) payable returns(uint256 totalCostWithFee)
func (_Kingdomly *KingdomlyTransactor) DelegatedMint(opts *bind.TransactOpts, amount *big.Int, mintId *big.Int, vault common.Address) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "delegatedMint", amount, mintId, vault)
}

// DelegatedMint is a paid mutator transaction binding the contract method 0xebf39a41.
//
// Solidity: function delegatedMint(uint256 amount, uint256 mintId, address vault) payable returns(uint256 totalCostWithFee)
func (_Kingdomly *KingdomlySession) DelegatedMint(amount *big.Int, mintId *big.Int, vault common.Address) (*types.Transaction, error) {
	return _Kingdomly.Contract.DelegatedMint(&_Kingdomly.TransactOpts, amount, mintId, vault)
}

// DelegatedMint is a paid mutator transaction binding the contract method 0xebf39a41.
//
// Solidity: function delegatedMint(uint256 amount, uint256 mintId, address vault) payable returns(uint256 totalCostWithFee)
func (_Kingdomly *KingdomlyTransactorSession) DelegatedMint(amount *big.Int, mintId *big.Int, vault common.Address) (*types.Transaction, error) {
	return _Kingdomly.Contract.DelegatedMint(&_Kingdomly.TransactOpts, amount, mintId, vault)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kingdomly *KingdomlyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kingdomly *KingdomlySession) RenounceOwnership() (*types.Transaction, error) {
	return _Kingdomly.Contract.RenounceOwnership(&_Kingdomly.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kingdomly *KingdomlyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kingdomly.Contract.RenounceOwnership(&_Kingdomly.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlyTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SafeTransferFrom(&_Kingdomly.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlyTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SafeTransferFrom(&_Kingdomly.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Kingdomly *KingdomlyTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Kingdomly *KingdomlySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Kingdomly.Contract.SafeTransferFrom0(&_Kingdomly.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Kingdomly *KingdomlyTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Kingdomly.Contract.SafeTransferFrom0(&_Kingdomly.TransactOpts, from, to, tokenId, _data)
}

// SchedulePresaleMintStart is a paid mutator transaction binding the contract method 0x80348242.
//
// Solidity: function schedulePresaleMintStart(uint256 startTimestamp, uint256 mintId) returns()
func (_Kingdomly *KingdomlyTransactor) SchedulePresaleMintStart(opts *bind.TransactOpts, startTimestamp *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "schedulePresaleMintStart", startTimestamp, mintId)
}

// SchedulePresaleMintStart is a paid mutator transaction binding the contract method 0x80348242.
//
// Solidity: function schedulePresaleMintStart(uint256 startTimestamp, uint256 mintId) returns()
func (_Kingdomly *KingdomlySession) SchedulePresaleMintStart(startTimestamp *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SchedulePresaleMintStart(&_Kingdomly.TransactOpts, startTimestamp, mintId)
}

// SchedulePresaleMintStart is a paid mutator transaction binding the contract method 0x80348242.
//
// Solidity: function schedulePresaleMintStart(uint256 startTimestamp, uint256 mintId) returns()
func (_Kingdomly *KingdomlyTransactorSession) SchedulePresaleMintStart(startTimestamp *big.Int, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SchedulePresaleMintStart(&_Kingdomly.TransactOpts, startTimestamp, mintId)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kingdomly *KingdomlyTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kingdomly *KingdomlySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetApprovalForAll(&_Kingdomly.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kingdomly *KingdomlyTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetApprovalForAll(&_Kingdomly.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Kingdomly *KingdomlyTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Kingdomly *KingdomlySession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetBaseURI(&_Kingdomly.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Kingdomly *KingdomlyTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetBaseURI(&_Kingdomly.TransactOpts, newBaseURI)
}

// SetMaxMintPerWallet is a paid mutator transaction binding the contract method 0x7bd4f071.
//
// Solidity: function setMaxMintPerWallet(uint256 newMaxMintPerWallet, uint256 mintGroupId) returns()
func (_Kingdomly *KingdomlyTransactor) SetMaxMintPerWallet(opts *bind.TransactOpts, newMaxMintPerWallet *big.Int, mintGroupId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "setMaxMintPerWallet", newMaxMintPerWallet, mintGroupId)
}

// SetMaxMintPerWallet is a paid mutator transaction binding the contract method 0x7bd4f071.
//
// Solidity: function setMaxMintPerWallet(uint256 newMaxMintPerWallet, uint256 mintGroupId) returns()
func (_Kingdomly *KingdomlySession) SetMaxMintPerWallet(newMaxMintPerWallet *big.Int, mintGroupId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetMaxMintPerWallet(&_Kingdomly.TransactOpts, newMaxMintPerWallet, mintGroupId)
}

// SetMaxMintPerWallet is a paid mutator transaction binding the contract method 0x7bd4f071.
//
// Solidity: function setMaxMintPerWallet(uint256 newMaxMintPerWallet, uint256 mintGroupId) returns()
func (_Kingdomly *KingdomlyTransactorSession) SetMaxMintPerWallet(newMaxMintPerWallet *big.Int, mintGroupId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetMaxMintPerWallet(&_Kingdomly.TransactOpts, newMaxMintPerWallet, mintGroupId)
}

// SetMintLiveTimestamp is a paid mutator transaction binding the contract method 0x31f72d77.
//
// Solidity: function setMintLiveTimestamp(uint256 timestamp) returns()
func (_Kingdomly *KingdomlyTransactor) SetMintLiveTimestamp(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "setMintLiveTimestamp", timestamp)
}

// SetMintLiveTimestamp is a paid mutator transaction binding the contract method 0x31f72d77.
//
// Solidity: function setMintLiveTimestamp(uint256 timestamp) returns()
func (_Kingdomly *KingdomlySession) SetMintLiveTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetMintLiveTimestamp(&_Kingdomly.TransactOpts, timestamp)
}

// SetMintLiveTimestamp is a paid mutator transaction binding the contract method 0x31f72d77.
//
// Solidity: function setMintLiveTimestamp(uint256 timestamp) returns()
func (_Kingdomly *KingdomlyTransactorSession) SetMintLiveTimestamp(timestamp *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetMintLiveTimestamp(&_Kingdomly.TransactOpts, timestamp)
}

// SetMintQuota is a paid mutator transaction binding the contract method 0x581636dd.
//
// Solidity: function setMintQuota(address[] addressToAdd, uint256 mintId, uint256[] _mintQuotas) returns()
func (_Kingdomly *KingdomlyTransactor) SetMintQuota(opts *bind.TransactOpts, addressToAdd []common.Address, mintId *big.Int, _mintQuotas []*big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "setMintQuota", addressToAdd, mintId, _mintQuotas)
}

// SetMintQuota is a paid mutator transaction binding the contract method 0x581636dd.
//
// Solidity: function setMintQuota(address[] addressToAdd, uint256 mintId, uint256[] _mintQuotas) returns()
func (_Kingdomly *KingdomlySession) SetMintQuota(addressToAdd []common.Address, mintId *big.Int, _mintQuotas []*big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetMintQuota(&_Kingdomly.TransactOpts, addressToAdd, mintId, _mintQuotas)
}

// SetMintQuota is a paid mutator transaction binding the contract method 0x581636dd.
//
// Solidity: function setMintQuota(address[] addressToAdd, uint256 mintId, uint256[] _mintQuotas) returns()
func (_Kingdomly *KingdomlyTransactorSession) SetMintQuota(addressToAdd []common.Address, mintId *big.Int, _mintQuotas []*big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetMintQuota(&_Kingdomly.TransactOpts, addressToAdd, mintId, _mintQuotas)
}

// SetNewKingdomlyFeeContract is a paid mutator transaction binding the contract method 0x1f466342.
//
// Solidity: function setNewKingdomlyFeeContract(address _kingdomlyFeeContract) returns()
func (_Kingdomly *KingdomlyTransactor) SetNewKingdomlyFeeContract(opts *bind.TransactOpts, _kingdomlyFeeContract common.Address) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "setNewKingdomlyFeeContract", _kingdomlyFeeContract)
}

// SetNewKingdomlyFeeContract is a paid mutator transaction binding the contract method 0x1f466342.
//
// Solidity: function setNewKingdomlyFeeContract(address _kingdomlyFeeContract) returns()
func (_Kingdomly *KingdomlySession) SetNewKingdomlyFeeContract(_kingdomlyFeeContract common.Address) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetNewKingdomlyFeeContract(&_Kingdomly.TransactOpts, _kingdomlyFeeContract)
}

// SetNewKingdomlyFeeContract is a paid mutator transaction binding the contract method 0x1f466342.
//
// Solidity: function setNewKingdomlyFeeContract(address _kingdomlyFeeContract) returns()
func (_Kingdomly *KingdomlyTransactorSession) SetNewKingdomlyFeeContract(_kingdomlyFeeContract common.Address) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetNewKingdomlyFeeContract(&_Kingdomly.TransactOpts, _kingdomlyFeeContract)
}

// SetNewMaxPerMintGroup is a paid mutator transaction binding the contract method 0x11f7acb9.
//
// Solidity: function setNewMaxPerMintGroup(uint256 mintId, uint256 newMax) returns()
func (_Kingdomly *KingdomlyTransactor) SetNewMaxPerMintGroup(opts *bind.TransactOpts, mintId *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "setNewMaxPerMintGroup", mintId, newMax)
}

// SetNewMaxPerMintGroup is a paid mutator transaction binding the contract method 0x11f7acb9.
//
// Solidity: function setNewMaxPerMintGroup(uint256 mintId, uint256 newMax) returns()
func (_Kingdomly *KingdomlySession) SetNewMaxPerMintGroup(mintId *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetNewMaxPerMintGroup(&_Kingdomly.TransactOpts, mintId, newMax)
}

// SetNewMaxPerMintGroup is a paid mutator transaction binding the contract method 0x11f7acb9.
//
// Solidity: function setNewMaxPerMintGroup(uint256 mintId, uint256 newMax) returns()
func (_Kingdomly *KingdomlyTransactorSession) SetNewMaxPerMintGroup(mintId *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.SetNewMaxPerMintGroup(&_Kingdomly.TransactOpts, mintId, newMax)
}

// StopOrStartpresaleMint is a paid mutator transaction binding the contract method 0xb3978a86.
//
// Solidity: function stopOrStartpresaleMint(bool presaleStatus, uint256 mintId) returns()
func (_Kingdomly *KingdomlyTransactor) StopOrStartpresaleMint(opts *bind.TransactOpts, presaleStatus bool, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "stopOrStartpresaleMint", presaleStatus, mintId)
}

// StopOrStartpresaleMint is a paid mutator transaction binding the contract method 0xb3978a86.
//
// Solidity: function stopOrStartpresaleMint(bool presaleStatus, uint256 mintId) returns()
func (_Kingdomly *KingdomlySession) StopOrStartpresaleMint(presaleStatus bool, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.StopOrStartpresaleMint(&_Kingdomly.TransactOpts, presaleStatus, mintId)
}

// StopOrStartpresaleMint is a paid mutator transaction binding the contract method 0xb3978a86.
//
// Solidity: function stopOrStartpresaleMint(bool presaleStatus, uint256 mintId) returns()
func (_Kingdomly *KingdomlyTransactorSession) StopOrStartpresaleMint(presaleStatus bool, mintId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.StopOrStartpresaleMint(&_Kingdomly.TransactOpts, presaleStatus, mintId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlyTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.TransferFrom(&_Kingdomly.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Kingdomly *KingdomlyTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kingdomly.Contract.TransferFrom(&_Kingdomly.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kingdomly *KingdomlyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kingdomly *KingdomlySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kingdomly.Contract.TransferOwnership(&_Kingdomly.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kingdomly *KingdomlyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kingdomly.Contract.TransferOwnership(&_Kingdomly.TransactOpts, newOwner)
}

// WithdrawFeeFunds is a paid mutator transaction binding the contract method 0xa75c3ad9.
//
// Solidity: function withdrawFeeFunds() returns()
func (_Kingdomly *KingdomlyTransactor) WithdrawFeeFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "withdrawFeeFunds")
}

// WithdrawFeeFunds is a paid mutator transaction binding the contract method 0xa75c3ad9.
//
// Solidity: function withdrawFeeFunds() returns()
func (_Kingdomly *KingdomlySession) WithdrawFeeFunds() (*types.Transaction, error) {
	return _Kingdomly.Contract.WithdrawFeeFunds(&_Kingdomly.TransactOpts)
}

// WithdrawFeeFunds is a paid mutator transaction binding the contract method 0xa75c3ad9.
//
// Solidity: function withdrawFeeFunds() returns()
func (_Kingdomly *KingdomlyTransactorSession) WithdrawFeeFunds() (*types.Transaction, error) {
	return _Kingdomly.Contract.WithdrawFeeFunds(&_Kingdomly.TransactOpts)
}

// WithdrawMintFunds is a paid mutator transaction binding the contract method 0x905d7b33.
//
// Solidity: function withdrawMintFunds() returns()
func (_Kingdomly *KingdomlyTransactor) WithdrawMintFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kingdomly.contract.Transact(opts, "withdrawMintFunds")
}

// WithdrawMintFunds is a paid mutator transaction binding the contract method 0x905d7b33.
//
// Solidity: function withdrawMintFunds() returns()
func (_Kingdomly *KingdomlySession) WithdrawMintFunds() (*types.Transaction, error) {
	return _Kingdomly.Contract.WithdrawMintFunds(&_Kingdomly.TransactOpts)
}

// WithdrawMintFunds is a paid mutator transaction binding the contract method 0x905d7b33.
//
// Solidity: function withdrawMintFunds() returns()
func (_Kingdomly *KingdomlyTransactorSession) WithdrawMintFunds() (*types.Transaction, error) {
	return _Kingdomly.Contract.WithdrawMintFunds(&_Kingdomly.TransactOpts)
}

// KingdomlyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Kingdomly contract.
type KingdomlyApprovalIterator struct {
	Event *KingdomlyApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyApproval represents a Approval event raised by the Kingdomly contract.
type KingdomlyApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Kingdomly *KingdomlyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KingdomlyApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyApprovalIterator{contract: _Kingdomly.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Kingdomly *KingdomlyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KingdomlyApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyApproval)
				if err := _Kingdomly.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Kingdomly *KingdomlyFilterer) ParseApproval(log types.Log) (*KingdomlyApproval, error) {
	event := new(KingdomlyApproval)
	if err := _Kingdomly.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Kingdomly contract.
type KingdomlyApprovalForAllIterator struct {
	Event *KingdomlyApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyApprovalForAll represents a ApprovalForAll event raised by the Kingdomly contract.
type KingdomlyApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Kingdomly *KingdomlyFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KingdomlyApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyApprovalForAllIterator{contract: _Kingdomly.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Kingdomly *KingdomlyFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KingdomlyApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyApprovalForAll)
				if err := _Kingdomly.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Kingdomly *KingdomlyFilterer) ParseApprovalForAll(log types.Log) (*KingdomlyApprovalForAll, error) {
	event := new(KingdomlyApprovalForAll)
	if err := _Kingdomly.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Kingdomly contract.
type KingdomlyBatchMetadataUpdateIterator struct {
	Event *KingdomlyBatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyBatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyBatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Kingdomly contract.
type KingdomlyBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 indexed fromTokenId, uint256 indexed toTokenId)
func (_Kingdomly *KingdomlyFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts, fromTokenId []*big.Int, toTokenId []*big.Int) (*KingdomlyBatchMetadataUpdateIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}
	var toTokenIdRule []interface{}
	for _, toTokenIdItem := range toTokenId {
		toTokenIdRule = append(toTokenIdRule, toTokenIdItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "BatchMetadataUpdate", fromTokenIdRule, toTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyBatchMetadataUpdateIterator{contract: _Kingdomly.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 indexed fromTokenId, uint256 indexed toTokenId)
func (_Kingdomly *KingdomlyFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *KingdomlyBatchMetadataUpdate, fromTokenId []*big.Int, toTokenId []*big.Int) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}
	var toTokenIdRule []interface{}
	for _, toTokenIdItem := range toTokenId {
		toTokenIdRule = append(toTokenIdRule, toTokenIdItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "BatchMetadataUpdate", fromTokenIdRule, toTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyBatchMetadataUpdate)
				if err := _Kingdomly.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 indexed fromTokenId, uint256 indexed toTokenId)
func (_Kingdomly *KingdomlyFilterer) ParseBatchMetadataUpdate(log types.Log) (*KingdomlyBatchMetadataUpdate, error) {
	event := new(KingdomlyBatchMetadataUpdate)
	if err := _Kingdomly.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Kingdomly contract.
type KingdomlyConsecutiveTransferIterator struct {
	Event *KingdomlyConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Kingdomly contract.
type KingdomlyConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Kingdomly *KingdomlyFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*KingdomlyConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyConsecutiveTransferIterator{contract: _Kingdomly.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Kingdomly *KingdomlyFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *KingdomlyConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyConsecutiveTransfer)
				if err := _Kingdomly.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Kingdomly *KingdomlyFilterer) ParseConsecutiveTransfer(log types.Log) (*KingdomlyConsecutiveTransfer, error) {
	event := new(KingdomlyConsecutiveTransfer)
	if err := _Kingdomly.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyKingdomlyFeeContractChangedIterator is returned from FilterKingdomlyFeeContractChanged and is used to iterate over the raw logs and unpacked data for KingdomlyFeeContractChanged events raised by the Kingdomly contract.
type KingdomlyKingdomlyFeeContractChangedIterator struct {
	Event *KingdomlyKingdomlyFeeContractChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyKingdomlyFeeContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyKingdomlyFeeContractChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyKingdomlyFeeContractChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyKingdomlyFeeContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyKingdomlyFeeContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyKingdomlyFeeContractChanged represents a KingdomlyFeeContractChanged event raised by the Kingdomly contract.
type KingdomlyKingdomlyFeeContractChanged struct {
	FeeContractAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterKingdomlyFeeContractChanged is a free log retrieval operation binding the contract event 0x65b193217dd691927510cfa45296799f4dc5a6b0d113a7f1863661cd57b1587f.
//
// Solidity: event KingdomlyFeeContractChanged(address feeContractAddress)
func (_Kingdomly *KingdomlyFilterer) FilterKingdomlyFeeContractChanged(opts *bind.FilterOpts) (*KingdomlyKingdomlyFeeContractChangedIterator, error) {

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "KingdomlyFeeContractChanged")
	if err != nil {
		return nil, err
	}
	return &KingdomlyKingdomlyFeeContractChangedIterator{contract: _Kingdomly.contract, event: "KingdomlyFeeContractChanged", logs: logs, sub: sub}, nil
}

// WatchKingdomlyFeeContractChanged is a free log subscription operation binding the contract event 0x65b193217dd691927510cfa45296799f4dc5a6b0d113a7f1863661cd57b1587f.
//
// Solidity: event KingdomlyFeeContractChanged(address feeContractAddress)
func (_Kingdomly *KingdomlyFilterer) WatchKingdomlyFeeContractChanged(opts *bind.WatchOpts, sink chan<- *KingdomlyKingdomlyFeeContractChanged) (event.Subscription, error) {

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "KingdomlyFeeContractChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyKingdomlyFeeContractChanged)
				if err := _Kingdomly.contract.UnpackLog(event, "KingdomlyFeeContractChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseKingdomlyFeeContractChanged is a log parse operation binding the contract event 0x65b193217dd691927510cfa45296799f4dc5a6b0d113a7f1863661cd57b1587f.
//
// Solidity: event KingdomlyFeeContractChanged(address feeContractAddress)
func (_Kingdomly *KingdomlyFilterer) ParseKingdomlyFeeContractChanged(log types.Log) (*KingdomlyKingdomlyFeeContractChanged, error) {
	event := new(KingdomlyKingdomlyFeeContractChanged)
	if err := _Kingdomly.contract.UnpackLog(event, "KingdomlyFeeContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyMaxMintPerWalletChangedIterator is returned from FilterMaxMintPerWalletChanged and is used to iterate over the raw logs and unpacked data for MaxMintPerWalletChanged events raised by the Kingdomly contract.
type KingdomlyMaxMintPerWalletChangedIterator struct {
	Event *KingdomlyMaxMintPerWalletChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyMaxMintPerWalletChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyMaxMintPerWalletChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyMaxMintPerWalletChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyMaxMintPerWalletChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyMaxMintPerWalletChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyMaxMintPerWalletChanged represents a MaxMintPerWalletChanged event raised by the Kingdomly contract.
type KingdomlyMaxMintPerWalletChanged struct {
	NewMaxMintPerWallet *big.Int
	MintGroupId         *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMaxMintPerWalletChanged is a free log retrieval operation binding the contract event 0xd6255c9c7b77c2a54f193e4634719645ef0fbdc4816638350b099b15a4ebca7f.
//
// Solidity: event MaxMintPerWalletChanged(uint256 newMaxMintPerWallet, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) FilterMaxMintPerWalletChanged(opts *bind.FilterOpts) (*KingdomlyMaxMintPerWalletChangedIterator, error) {

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "MaxMintPerWalletChanged")
	if err != nil {
		return nil, err
	}
	return &KingdomlyMaxMintPerWalletChangedIterator{contract: _Kingdomly.contract, event: "MaxMintPerWalletChanged", logs: logs, sub: sub}, nil
}

// WatchMaxMintPerWalletChanged is a free log subscription operation binding the contract event 0xd6255c9c7b77c2a54f193e4634719645ef0fbdc4816638350b099b15a4ebca7f.
//
// Solidity: event MaxMintPerWalletChanged(uint256 newMaxMintPerWallet, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) WatchMaxMintPerWalletChanged(opts *bind.WatchOpts, sink chan<- *KingdomlyMaxMintPerWalletChanged) (event.Subscription, error) {

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "MaxMintPerWalletChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyMaxMintPerWalletChanged)
				if err := _Kingdomly.contract.UnpackLog(event, "MaxMintPerWalletChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxMintPerWalletChanged is a log parse operation binding the contract event 0xd6255c9c7b77c2a54f193e4634719645ef0fbdc4816638350b099b15a4ebca7f.
//
// Solidity: event MaxMintPerWalletChanged(uint256 newMaxMintPerWallet, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) ParseMaxMintPerWalletChanged(log types.Log) (*KingdomlyMaxMintPerWalletChanged, error) {
	event := new(KingdomlyMaxMintPerWalletChanged)
	if err := _Kingdomly.contract.UnpackLog(event, "MaxMintPerWalletChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Kingdomly contract.
type KingdomlyOwnershipTransferredIterator struct {
	Event *KingdomlyOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyOwnershipTransferred represents a OwnershipTransferred event raised by the Kingdomly contract.
type KingdomlyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kingdomly *KingdomlyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KingdomlyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyOwnershipTransferredIterator{contract: _Kingdomly.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kingdomly *KingdomlyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KingdomlyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyOwnershipTransferred)
				if err := _Kingdomly.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kingdomly *KingdomlyFilterer) ParseOwnershipTransferred(log types.Log) (*KingdomlyOwnershipTransferred, error) {
	event := new(KingdomlyOwnershipTransferred)
	if err := _Kingdomly.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyPreSaleMintScheduledStartTimestampChangedIterator is returned from FilterPreSaleMintScheduledStartTimestampChanged and is used to iterate over the raw logs and unpacked data for PreSaleMintScheduledStartTimestampChanged events raised by the Kingdomly contract.
type KingdomlyPreSaleMintScheduledStartTimestampChangedIterator struct {
	Event *KingdomlyPreSaleMintScheduledStartTimestampChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyPreSaleMintScheduledStartTimestampChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyPreSaleMintScheduledStartTimestampChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyPreSaleMintScheduledStartTimestampChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyPreSaleMintScheduledStartTimestampChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyPreSaleMintScheduledStartTimestampChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyPreSaleMintScheduledStartTimestampChanged represents a PreSaleMintScheduledStartTimestampChanged event raised by the Kingdomly contract.
type KingdomlyPreSaleMintScheduledStartTimestampChanged struct {
	Timestamp   *big.Int
	MintGroupId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPreSaleMintScheduledStartTimestampChanged is a free log retrieval operation binding the contract event 0xb024af7f651e94348c4e66cfe71f68a3f246eff857f95f105766bc15f4ea84fc.
//
// Solidity: event PreSaleMintScheduledStartTimestampChanged(uint256 timestamp, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) FilterPreSaleMintScheduledStartTimestampChanged(opts *bind.FilterOpts) (*KingdomlyPreSaleMintScheduledStartTimestampChangedIterator, error) {

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "PreSaleMintScheduledStartTimestampChanged")
	if err != nil {
		return nil, err
	}
	return &KingdomlyPreSaleMintScheduledStartTimestampChangedIterator{contract: _Kingdomly.contract, event: "PreSaleMintScheduledStartTimestampChanged", logs: logs, sub: sub}, nil
}

// WatchPreSaleMintScheduledStartTimestampChanged is a free log subscription operation binding the contract event 0xb024af7f651e94348c4e66cfe71f68a3f246eff857f95f105766bc15f4ea84fc.
//
// Solidity: event PreSaleMintScheduledStartTimestampChanged(uint256 timestamp, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) WatchPreSaleMintScheduledStartTimestampChanged(opts *bind.WatchOpts, sink chan<- *KingdomlyPreSaleMintScheduledStartTimestampChanged) (event.Subscription, error) {

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "PreSaleMintScheduledStartTimestampChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyPreSaleMintScheduledStartTimestampChanged)
				if err := _Kingdomly.contract.UnpackLog(event, "PreSaleMintScheduledStartTimestampChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreSaleMintScheduledStartTimestampChanged is a log parse operation binding the contract event 0xb024af7f651e94348c4e66cfe71f68a3f246eff857f95f105766bc15f4ea84fc.
//
// Solidity: event PreSaleMintScheduledStartTimestampChanged(uint256 timestamp, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) ParsePreSaleMintScheduledStartTimestampChanged(log types.Log) (*KingdomlyPreSaleMintScheduledStartTimestampChanged, error) {
	event := new(KingdomlyPreSaleMintScheduledStartTimestampChanged)
	if err := _Kingdomly.contract.UnpackLog(event, "PreSaleMintScheduledStartTimestampChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyPreSaleMintStatusChangedIterator is returned from FilterPreSaleMintStatusChanged and is used to iterate over the raw logs and unpacked data for PreSaleMintStatusChanged events raised by the Kingdomly contract.
type KingdomlyPreSaleMintStatusChangedIterator struct {
	Event *KingdomlyPreSaleMintStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyPreSaleMintStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyPreSaleMintStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyPreSaleMintStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyPreSaleMintStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyPreSaleMintStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyPreSaleMintStatusChanged represents a PreSaleMintStatusChanged event raised by the Kingdomly contract.
type KingdomlyPreSaleMintStatusChanged struct {
	Status      bool
	MintGroupId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPreSaleMintStatusChanged is a free log retrieval operation binding the contract event 0xd88af4ad8d3188d34ada8f4e850b4b6b3b2610392d5c8ffffa63b28ae831da54.
//
// Solidity: event PreSaleMintStatusChanged(bool status, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) FilterPreSaleMintStatusChanged(opts *bind.FilterOpts) (*KingdomlyPreSaleMintStatusChangedIterator, error) {

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "PreSaleMintStatusChanged")
	if err != nil {
		return nil, err
	}
	return &KingdomlyPreSaleMintStatusChangedIterator{contract: _Kingdomly.contract, event: "PreSaleMintStatusChanged", logs: logs, sub: sub}, nil
}

// WatchPreSaleMintStatusChanged is a free log subscription operation binding the contract event 0xd88af4ad8d3188d34ada8f4e850b4b6b3b2610392d5c8ffffa63b28ae831da54.
//
// Solidity: event PreSaleMintStatusChanged(bool status, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) WatchPreSaleMintStatusChanged(opts *bind.WatchOpts, sink chan<- *KingdomlyPreSaleMintStatusChanged) (event.Subscription, error) {

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "PreSaleMintStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyPreSaleMintStatusChanged)
				if err := _Kingdomly.contract.UnpackLog(event, "PreSaleMintStatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreSaleMintStatusChanged is a log parse operation binding the contract event 0xd88af4ad8d3188d34ada8f4e850b4b6b3b2610392d5c8ffffa63b28ae831da54.
//
// Solidity: event PreSaleMintStatusChanged(bool status, uint256 mintGroupId)
func (_Kingdomly *KingdomlyFilterer) ParsePreSaleMintStatusChanged(log types.Log) (*KingdomlyPreSaleMintStatusChanged, error) {
	event := new(KingdomlyPreSaleMintStatusChanged)
	if err := _Kingdomly.contract.UnpackLog(event, "PreSaleMintStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlySalePriceChangedIterator is returned from FilterSalePriceChanged and is used to iterate over the raw logs and unpacked data for SalePriceChanged events raised by the Kingdomly contract.
type KingdomlySalePriceChangedIterator struct {
	Event *KingdomlySalePriceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlySalePriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlySalePriceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlySalePriceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlySalePriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlySalePriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlySalePriceChanged represents a SalePriceChanged event raised by the Kingdomly contract.
type KingdomlySalePriceChanged struct {
	MintId   *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSalePriceChanged is a free log retrieval operation binding the contract event 0xa7e52343431f792020e7cb8411a08014688ca11782fd5709fa2531b3d74ba457.
//
// Solidity: event SalePriceChanged(uint256 indexed mintId, uint256 newPrice)
func (_Kingdomly *KingdomlyFilterer) FilterSalePriceChanged(opts *bind.FilterOpts, mintId []*big.Int) (*KingdomlySalePriceChangedIterator, error) {

	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "SalePriceChanged", mintIdRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlySalePriceChangedIterator{contract: _Kingdomly.contract, event: "SalePriceChanged", logs: logs, sub: sub}, nil
}

// WatchSalePriceChanged is a free log subscription operation binding the contract event 0xa7e52343431f792020e7cb8411a08014688ca11782fd5709fa2531b3d74ba457.
//
// Solidity: event SalePriceChanged(uint256 indexed mintId, uint256 newPrice)
func (_Kingdomly *KingdomlyFilterer) WatchSalePriceChanged(opts *bind.WatchOpts, sink chan<- *KingdomlySalePriceChanged, mintId []*big.Int) (event.Subscription, error) {

	var mintIdRule []interface{}
	for _, mintIdItem := range mintId {
		mintIdRule = append(mintIdRule, mintIdItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "SalePriceChanged", mintIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlySalePriceChanged)
				if err := _Kingdomly.contract.UnpackLog(event, "SalePriceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSalePriceChanged is a log parse operation binding the contract event 0xa7e52343431f792020e7cb8411a08014688ca11782fd5709fa2531b3d74ba457.
//
// Solidity: event SalePriceChanged(uint256 indexed mintId, uint256 newPrice)
func (_Kingdomly *KingdomlyFilterer) ParseSalePriceChanged(log types.Log) (*KingdomlySalePriceChanged, error) {
	event := new(KingdomlySalePriceChanged)
	if err := _Kingdomly.contract.UnpackLog(event, "SalePriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyTokensDelegateMintedIterator is returned from FilterTokensDelegateMinted and is used to iterate over the raw logs and unpacked data for TokensDelegateMinted events raised by the Kingdomly contract.
type KingdomlyTokensDelegateMintedIterator struct {
	Event *KingdomlyTokensDelegateMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyTokensDelegateMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyTokensDelegateMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyTokensDelegateMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyTokensDelegateMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyTokensDelegateMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyTokensDelegateMinted represents a TokensDelegateMinted event raised by the Kingdomly contract.
type KingdomlyTokensDelegateMinted struct {
	Vault     common.Address
	HotWallet common.Address
	Amount    *big.Int
	MintId    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensDelegateMinted is a free log retrieval operation binding the contract event 0xce55cbb6a167cf85969795eda1f8d0a2e0152274849017210ca560bd3c963f7f.
//
// Solidity: event TokensDelegateMinted(address indexed vault, address indexed hotWallet, uint256 amount, uint256 mintId)
func (_Kingdomly *KingdomlyFilterer) FilterTokensDelegateMinted(opts *bind.FilterOpts, vault []common.Address, hotWallet []common.Address) (*KingdomlyTokensDelegateMintedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var hotWalletRule []interface{}
	for _, hotWalletItem := range hotWallet {
		hotWalletRule = append(hotWalletRule, hotWalletItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "TokensDelegateMinted", vaultRule, hotWalletRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyTokensDelegateMintedIterator{contract: _Kingdomly.contract, event: "TokensDelegateMinted", logs: logs, sub: sub}, nil
}

// WatchTokensDelegateMinted is a free log subscription operation binding the contract event 0xce55cbb6a167cf85969795eda1f8d0a2e0152274849017210ca560bd3c963f7f.
//
// Solidity: event TokensDelegateMinted(address indexed vault, address indexed hotWallet, uint256 amount, uint256 mintId)
func (_Kingdomly *KingdomlyFilterer) WatchTokensDelegateMinted(opts *bind.WatchOpts, sink chan<- *KingdomlyTokensDelegateMinted, vault []common.Address, hotWallet []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var hotWalletRule []interface{}
	for _, hotWalletItem := range hotWallet {
		hotWalletRule = append(hotWalletRule, hotWalletItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "TokensDelegateMinted", vaultRule, hotWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyTokensDelegateMinted)
				if err := _Kingdomly.contract.UnpackLog(event, "TokensDelegateMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensDelegateMinted is a log parse operation binding the contract event 0xce55cbb6a167cf85969795eda1f8d0a2e0152274849017210ca560bd3c963f7f.
//
// Solidity: event TokensDelegateMinted(address indexed vault, address indexed hotWallet, uint256 amount, uint256 mintId)
func (_Kingdomly *KingdomlyFilterer) ParseTokensDelegateMinted(log types.Log) (*KingdomlyTokensDelegateMinted, error) {
	event := new(KingdomlyTokensDelegateMinted)
	if err := _Kingdomly.contract.UnpackLog(event, "TokensDelegateMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyTokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the Kingdomly contract.
type KingdomlyTokensMintedIterator struct {
	Event *KingdomlyTokensMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyTokensMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyTokensMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyTokensMinted represents a TokensMinted event raised by the Kingdomly contract.
type KingdomlyTokensMinted struct {
	Recipient common.Address
	Amount    *big.Int
	MintId    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x2e8ac5177a616f2aec08c3048f5021e4e9743ece034e8d83ba5caf76688bb475.
//
// Solidity: event TokensMinted(address indexed recipient, uint256 amount, uint256 mintId)
func (_Kingdomly *KingdomlyFilterer) FilterTokensMinted(opts *bind.FilterOpts, recipient []common.Address) (*KingdomlyTokensMintedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "TokensMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyTokensMintedIterator{contract: _Kingdomly.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x2e8ac5177a616f2aec08c3048f5021e4e9743ece034e8d83ba5caf76688bb475.
//
// Solidity: event TokensMinted(address indexed recipient, uint256 amount, uint256 mintId)
func (_Kingdomly *KingdomlyFilterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *KingdomlyTokensMinted, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "TokensMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyTokensMinted)
				if err := _Kingdomly.contract.UnpackLog(event, "TokensMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensMinted is a log parse operation binding the contract event 0x2e8ac5177a616f2aec08c3048f5021e4e9743ece034e8d83ba5caf76688bb475.
//
// Solidity: event TokensMinted(address indexed recipient, uint256 amount, uint256 mintId)
func (_Kingdomly *KingdomlyFilterer) ParseTokensMinted(log types.Log) (*KingdomlyTokensMinted, error) {
	event := new(KingdomlyTokensMinted)
	if err := _Kingdomly.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KingdomlyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Kingdomly contract.
type KingdomlyTransferIterator struct {
	Event *KingdomlyTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KingdomlyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KingdomlyTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KingdomlyTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KingdomlyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KingdomlyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KingdomlyTransfer represents a Transfer event raised by the Kingdomly contract.
type KingdomlyTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Kingdomly *KingdomlyFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KingdomlyTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Kingdomly.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KingdomlyTransferIterator{contract: _Kingdomly.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Kingdomly *KingdomlyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KingdomlyTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Kingdomly.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KingdomlyTransfer)
				if err := _Kingdomly.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Kingdomly *KingdomlyFilterer) ParseTransfer(log types.Log) (*KingdomlyTransfer, error) {
	event := new(KingdomlyTransfer)
	if err := _Kingdomly.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
