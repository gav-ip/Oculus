// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// CollectionSecurityPolicy is an auto generated low-level Go binding around an user-defined struct.
type CollectionSecurityPolicy struct {
	TransferSecurityLevel        uint8
	OperatorWhitelistId          *big.Int
	PermittedContractReceiversId *big.Int
}

// IERC721ATokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721ATokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// IERC721MMintStageInfo is an auto generated low-level Go binding around an user-defined struct.
type IERC721MMintStageInfo struct {
	Price                *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds uint64
	EndTimeUnixSeconds   uint64
}

// MagicedenMetaData contains all meta data concerning the Magiceden contract.
var MagicedenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"collectionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"collectionSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURISuffix\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestampExpirySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyFeeNumerator\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotIncreaseMaxMintableSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdatePermanentBaseURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignerNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorTokenBase__InvalidTransferValidatorContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorTokenBase__SetTransferValidatorFirst\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossmintAddressNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CrossmintOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalWalletLimitOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStageTimeGap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStageArgsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSupplyLeft\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ShouldNotMintToBurnAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StageSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletGlobalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletStageLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongMintCurrency\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"DefaultRoyaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"PermanentBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeStage\",\"type\":\"uint256\"}],\"name\":\"SetActiveStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"SetBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"SetCosigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"crossmintAddress\",\"type\":\"address\"}],\"name\":\"SetCrossmintAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"SetGlobalWalletLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"SetMaxMintableSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"}],\"name\":\"SetMintCurrency\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"SetMintable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"SetTimestampExpirySeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"TokenRoyaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"TransferValidatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"name\":\"UpdateStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_OPERATOR_WHITELIST_ID\",\"outputs\":[{\"internalType\":\"uint120\",\"name\":\"\",\"type\":\"uint120\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_TRANSFER_SECURITY_LEVEL\",\"outputs\":[{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_TRANSFER_VALIDATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"assertValidCosign\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"crossmint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"explicitOwnershipsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721A.TokenOwnership[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getActiveStageFromTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getCosignDigest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getCosignNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalWalletLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMintableSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberStages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermittedContractReceivers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecurityPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"transferSecurityLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"operatorWhitelistId\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"permittedContractReceiversId\",\"type\":\"uint120\"}],\"internalType\":\"structCollectionSecurityPolicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStageInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structIERC721M.MintStageInfo\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransferValidator\",\"outputs\":[{\"internalType\":\"contractICreatorTokenTransferValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"isContractReceiverPermitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"isTransferAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintWithLimit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ownerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setBaseURIPermanent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"setCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossmintAddress\",\"type\":\"address\"}],\"name\":\"setCrossmintAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"setGlobalWalletLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxMintableSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structIERC721M.MintStageInfo[]\",\"name\":\"newStages\",\"type\":\"tuple[]\"}],\"name\":\"setStages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setTimestampExpirySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"operatorWhitelistId\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"permittedContractReceiversAllowlistId\",\"type\":\"uint120\"}],\"name\":\"setToCustomSecurityPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumTransferSecurityLevels\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"uint120\",\"name\":\"operatorWhitelistId\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"permittedContractReceiversAllowlistId\",\"type\":\"uint120\"}],\"name\":\"setToCustomValidatorAndSecurityPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setToDefaultSecurityPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setTokenRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"name\":\"setTokenURISuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transferValidator_\",\"type\":\"address\"}],\"name\":\"setTransferValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"totalMintedByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint64\"}],\"name\":\"updateStage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MagicedenABI is the input ABI used to generate the binding from.
// Deprecated: Use MagicedenMetaData.ABI instead.
var MagicedenABI = MagicedenMetaData.ABI

// Magiceden is an auto generated Go binding around an Ethereum contract.
type Magiceden struct {
	MagicedenCaller     // Read-only binding to the contract
	MagicedenTransactor // Write-only binding to the contract
	MagicedenFilterer   // Log filterer for contract events
}

// MagicedenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MagicedenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicedenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MagicedenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicedenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MagicedenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicedenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MagicedenSession struct {
	Contract     *Magiceden        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MagicedenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MagicedenCallerSession struct {
	Contract *MagicedenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MagicedenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MagicedenTransactorSession struct {
	Contract     *MagicedenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MagicedenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MagicedenRaw struct {
	Contract *Magiceden // Generic contract binding to access the raw methods on
}

// MagicedenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MagicedenCallerRaw struct {
	Contract *MagicedenCaller // Generic read-only contract binding to access the raw methods on
}

// MagicedenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MagicedenTransactorRaw struct {
	Contract *MagicedenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMagiceden creates a new instance of Magiceden, bound to a specific deployed contract.
func NewMagiceden(address common.Address, backend bind.ContractBackend) (*Magiceden, error) {
	contract, err := bindMagiceden(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Magiceden{MagicedenCaller: MagicedenCaller{contract: contract}, MagicedenTransactor: MagicedenTransactor{contract: contract}, MagicedenFilterer: MagicedenFilterer{contract: contract}}, nil
}

// NewMagicedenCaller creates a new read-only instance of Magiceden, bound to a specific deployed contract.
func NewMagicedenCaller(address common.Address, caller bind.ContractCaller) (*MagicedenCaller, error) {
	contract, err := bindMagiceden(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MagicedenCaller{contract: contract}, nil
}

// NewMagicedenTransactor creates a new write-only instance of Magiceden, bound to a specific deployed contract.
func NewMagicedenTransactor(address common.Address, transactor bind.ContractTransactor) (*MagicedenTransactor, error) {
	contract, err := bindMagiceden(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MagicedenTransactor{contract: contract}, nil
}

// NewMagicedenFilterer creates a new log filterer instance of Magiceden, bound to a specific deployed contract.
func NewMagicedenFilterer(address common.Address, filterer bind.ContractFilterer) (*MagicedenFilterer, error) {
	contract, err := bindMagiceden(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MagicedenFilterer{contract: contract}, nil
}

// bindMagiceden binds a generic wrapper to an already deployed contract.
func bindMagiceden(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MagicedenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Magiceden *MagicedenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Magiceden.Contract.MagicedenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Magiceden *MagicedenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.Contract.MagicedenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Magiceden *MagicedenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Magiceden.Contract.MagicedenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Magiceden *MagicedenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Magiceden.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Magiceden *MagicedenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Magiceden *MagicedenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Magiceden.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTOPERATORWHITELISTID is a free data retrieval call binding the contract method 0x5d4c1d46.
//
// Solidity: function DEFAULT_OPERATOR_WHITELIST_ID() view returns(uint120)
func (_Magiceden *MagicedenCaller) DEFAULTOPERATORWHITELISTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "DEFAULT_OPERATOR_WHITELIST_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTOPERATORWHITELISTID is a free data retrieval call binding the contract method 0x5d4c1d46.
//
// Solidity: function DEFAULT_OPERATOR_WHITELIST_ID() view returns(uint120)
func (_Magiceden *MagicedenSession) DEFAULTOPERATORWHITELISTID() (*big.Int, error) {
	return _Magiceden.Contract.DEFAULTOPERATORWHITELISTID(&_Magiceden.CallOpts)
}

// DEFAULTOPERATORWHITELISTID is a free data retrieval call binding the contract method 0x5d4c1d46.
//
// Solidity: function DEFAULT_OPERATOR_WHITELIST_ID() view returns(uint120)
func (_Magiceden *MagicedenCallerSession) DEFAULTOPERATORWHITELISTID() (*big.Int, error) {
	return _Magiceden.Contract.DEFAULTOPERATORWHITELISTID(&_Magiceden.CallOpts)
}

// DEFAULTTRANSFERSECURITYLEVEL is a free data retrieval call binding the contract method 0x1c33b328.
//
// Solidity: function DEFAULT_TRANSFER_SECURITY_LEVEL() view returns(uint8)
func (_Magiceden *MagicedenCaller) DEFAULTTRANSFERSECURITYLEVEL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "DEFAULT_TRANSFER_SECURITY_LEVEL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DEFAULTTRANSFERSECURITYLEVEL is a free data retrieval call binding the contract method 0x1c33b328.
//
// Solidity: function DEFAULT_TRANSFER_SECURITY_LEVEL() view returns(uint8)
func (_Magiceden *MagicedenSession) DEFAULTTRANSFERSECURITYLEVEL() (uint8, error) {
	return _Magiceden.Contract.DEFAULTTRANSFERSECURITYLEVEL(&_Magiceden.CallOpts)
}

// DEFAULTTRANSFERSECURITYLEVEL is a free data retrieval call binding the contract method 0x1c33b328.
//
// Solidity: function DEFAULT_TRANSFER_SECURITY_LEVEL() view returns(uint8)
func (_Magiceden *MagicedenCallerSession) DEFAULTTRANSFERSECURITYLEVEL() (uint8, error) {
	return _Magiceden.Contract.DEFAULTTRANSFERSECURITYLEVEL(&_Magiceden.CallOpts)
}

// DEFAULTTRANSFERVALIDATOR is a free data retrieval call binding the contract method 0x01463546.
//
// Solidity: function DEFAULT_TRANSFER_VALIDATOR() view returns(address)
func (_Magiceden *MagicedenCaller) DEFAULTTRANSFERVALIDATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "DEFAULT_TRANSFER_VALIDATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTTRANSFERVALIDATOR is a free data retrieval call binding the contract method 0x01463546.
//
// Solidity: function DEFAULT_TRANSFER_VALIDATOR() view returns(address)
func (_Magiceden *MagicedenSession) DEFAULTTRANSFERVALIDATOR() (common.Address, error) {
	return _Magiceden.Contract.DEFAULTTRANSFERVALIDATOR(&_Magiceden.CallOpts)
}

// DEFAULTTRANSFERVALIDATOR is a free data retrieval call binding the contract method 0x01463546.
//
// Solidity: function DEFAULT_TRANSFER_VALIDATOR() view returns(address)
func (_Magiceden *MagicedenCallerSession) DEFAULTTRANSFERVALIDATOR() (common.Address, error) {
	return _Magiceden.Contract.DEFAULTTRANSFERVALIDATOR(&_Magiceden.CallOpts)
}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Magiceden *MagicedenCaller) AssertValidCosign(opts *bind.CallOpts, minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "assertValidCosign", minter, qty, timestamp, signature)

	if err != nil {
		return err
	}

	return err

}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Magiceden *MagicedenSession) AssertValidCosign(minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	return _Magiceden.Contract.AssertValidCosign(&_Magiceden.CallOpts, minter, qty, timestamp, signature)
}

// AssertValidCosign is a free data retrieval call binding the contract method 0xb50248e7.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint64 timestamp, bytes signature) view returns()
func (_Magiceden *MagicedenCallerSession) AssertValidCosign(minter common.Address, qty uint32, timestamp uint64, signature []byte) error {
	return _Magiceden.Contract.AssertValidCosign(&_Magiceden.CallOpts, minter, qty, timestamp, signature)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Magiceden *MagicedenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Magiceden *MagicedenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Magiceden.Contract.BalanceOf(&_Magiceden.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Magiceden *MagicedenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Magiceden.Contract.BalanceOf(&_Magiceden.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Magiceden *MagicedenCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Magiceden *MagicedenSession) ContractURI() (string, error) {
	return _Magiceden.Contract.ContractURI(&_Magiceden.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Magiceden *MagicedenCallerSession) ContractURI() (string, error) {
	return _Magiceden.Contract.ContractURI(&_Magiceden.CallOpts)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Magiceden *MagicedenCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721ATokenOwnership)).(*IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Magiceden *MagicedenSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipOf(&_Magiceden.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24))
func (_Magiceden *MagicedenCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721ATokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipOf(&_Magiceden.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Magiceden *MagicedenCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721ATokenOwnership)).(*[]IERC721ATokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Magiceden *MagicedenSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipsOf(&_Magiceden.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Magiceden *MagicedenCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721ATokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipsOf(&_Magiceden.CallOpts, tokenIds)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Magiceden *MagicedenCaller) GetActiveStageFromTimestamp(opts *bind.CallOpts, timestamp uint64) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getActiveStageFromTimestamp", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Magiceden *MagicedenSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Magiceden.Contract.GetActiveStageFromTimestamp(&_Magiceden.CallOpts, timestamp)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x67808a34.
//
// Solidity: function getActiveStageFromTimestamp(uint64 timestamp) view returns(uint256)
func (_Magiceden *MagicedenCallerSession) GetActiveStageFromTimestamp(timestamp uint64) (*big.Int, error) {
	return _Magiceden.Contract.GetActiveStageFromTimestamp(&_Magiceden.CallOpts, timestamp)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Magiceden *MagicedenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Magiceden *MagicedenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Magiceden.Contract.GetApproved(&_Magiceden.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Magiceden *MagicedenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Magiceden.Contract.GetApproved(&_Magiceden.CallOpts, tokenId)
}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Magiceden *MagicedenCaller) GetCosignDigest(opts *bind.CallOpts, minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getCosignDigest", minter, qty, timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Magiceden *MagicedenSession) GetCosignDigest(minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	return _Magiceden.Contract.GetCosignDigest(&_Magiceden.CallOpts, minter, qty, timestamp)
}

// GetCosignDigest is a free data retrieval call binding the contract method 0x1ce03eed.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, uint64 timestamp) view returns(bytes32)
func (_Magiceden *MagicedenCallerSession) GetCosignDigest(minter common.Address, qty uint32, timestamp uint64) ([32]byte, error) {
	return _Magiceden.Contract.GetCosignDigest(&_Magiceden.CallOpts, minter, qty, timestamp)
}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Magiceden *MagicedenCaller) GetCosignNonce(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getCosignNonce", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Magiceden *MagicedenSession) GetCosignNonce(minter common.Address) (*big.Int, error) {
	return _Magiceden.Contract.GetCosignNonce(&_Magiceden.CallOpts, minter)
}

// GetCosignNonce is a free data retrieval call binding the contract method 0xa06c492f.
//
// Solidity: function getCosignNonce(address minter) view returns(uint256)
func (_Magiceden *MagicedenCallerSession) GetCosignNonce(minter common.Address) (*big.Int, error) {
	return _Magiceden.Contract.GetCosignNonce(&_Magiceden.CallOpts, minter)
}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Magiceden *MagicedenCaller) GetGlobalWalletLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getGlobalWalletLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Magiceden *MagicedenSession) GetGlobalWalletLimit() (*big.Int, error) {
	return _Magiceden.Contract.GetGlobalWalletLimit(&_Magiceden.CallOpts)
}

// GetGlobalWalletLimit is a free data retrieval call binding the contract method 0xefdaa2ec.
//
// Solidity: function getGlobalWalletLimit() view returns(uint256)
func (_Magiceden *MagicedenCallerSession) GetGlobalWalletLimit() (*big.Int, error) {
	return _Magiceden.Contract.GetGlobalWalletLimit(&_Magiceden.CallOpts)
}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Magiceden *MagicedenCaller) GetMaxMintableSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getMaxMintableSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Magiceden *MagicedenSession) GetMaxMintableSupply() (*big.Int, error) {
	return _Magiceden.Contract.GetMaxMintableSupply(&_Magiceden.CallOpts)
}

// GetMaxMintableSupply is a free data retrieval call binding the contract method 0x4b1c53b4.
//
// Solidity: function getMaxMintableSupply() view returns(uint256)
func (_Magiceden *MagicedenCallerSession) GetMaxMintableSupply() (*big.Int, error) {
	return _Magiceden.Contract.GetMaxMintableSupply(&_Magiceden.CallOpts)
}

// GetMintCurrency is a free data retrieval call binding the contract method 0x424aa884.
//
// Solidity: function getMintCurrency() view returns(address)
func (_Magiceden *MagicedenCaller) GetMintCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getMintCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMintCurrency is a free data retrieval call binding the contract method 0x424aa884.
//
// Solidity: function getMintCurrency() view returns(address)
func (_Magiceden *MagicedenSession) GetMintCurrency() (common.Address, error) {
	return _Magiceden.Contract.GetMintCurrency(&_Magiceden.CallOpts)
}

// GetMintCurrency is a free data retrieval call binding the contract method 0x424aa884.
//
// Solidity: function getMintCurrency() view returns(address)
func (_Magiceden *MagicedenCallerSession) GetMintCurrency() (common.Address, error) {
	return _Magiceden.Contract.GetMintCurrency(&_Magiceden.CallOpts)
}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Magiceden *MagicedenCaller) GetMintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getMintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Magiceden *MagicedenSession) GetMintable() (bool, error) {
	return _Magiceden.Contract.GetMintable(&_Magiceden.CallOpts)
}

// GetMintable is a free data retrieval call binding the contract method 0xf698bceb.
//
// Solidity: function getMintable() view returns(bool)
func (_Magiceden *MagicedenCallerSession) GetMintable() (bool, error) {
	return _Magiceden.Contract.GetMintable(&_Magiceden.CallOpts)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Magiceden *MagicedenCaller) GetNumberStages(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getNumberStages")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Magiceden *MagicedenSession) GetNumberStages() (*big.Int, error) {
	return _Magiceden.Contract.GetNumberStages(&_Magiceden.CallOpts)
}

// GetNumberStages is a free data retrieval call binding the contract method 0x70da24ee.
//
// Solidity: function getNumberStages() view returns(uint256)
func (_Magiceden *MagicedenCallerSession) GetNumberStages() (*big.Int, error) {
	return _Magiceden.Contract.GetNumberStages(&_Magiceden.CallOpts)
}

// GetPermittedContractReceivers is a free data retrieval call binding the contract method 0xd007af5c.
//
// Solidity: function getPermittedContractReceivers() view returns(address[])
func (_Magiceden *MagicedenCaller) GetPermittedContractReceivers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getPermittedContractReceivers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPermittedContractReceivers is a free data retrieval call binding the contract method 0xd007af5c.
//
// Solidity: function getPermittedContractReceivers() view returns(address[])
func (_Magiceden *MagicedenSession) GetPermittedContractReceivers() ([]common.Address, error) {
	return _Magiceden.Contract.GetPermittedContractReceivers(&_Magiceden.CallOpts)
}

// GetPermittedContractReceivers is a free data retrieval call binding the contract method 0xd007af5c.
//
// Solidity: function getPermittedContractReceivers() view returns(address[])
func (_Magiceden *MagicedenCallerSession) GetPermittedContractReceivers() ([]common.Address, error) {
	return _Magiceden.Contract.GetPermittedContractReceivers(&_Magiceden.CallOpts)
}

// GetSecurityPolicy is a free data retrieval call binding the contract method 0xbe537f43.
//
// Solidity: function getSecurityPolicy() view returns((uint8,uint120,uint120))
func (_Magiceden *MagicedenCaller) GetSecurityPolicy(opts *bind.CallOpts) (CollectionSecurityPolicy, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getSecurityPolicy")

	if err != nil {
		return *new(CollectionSecurityPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(CollectionSecurityPolicy)).(*CollectionSecurityPolicy)

	return out0, err

}

// GetSecurityPolicy is a free data retrieval call binding the contract method 0xbe537f43.
//
// Solidity: function getSecurityPolicy() view returns((uint8,uint120,uint120))
func (_Magiceden *MagicedenSession) GetSecurityPolicy() (CollectionSecurityPolicy, error) {
	return _Magiceden.Contract.GetSecurityPolicy(&_Magiceden.CallOpts)
}

// GetSecurityPolicy is a free data retrieval call binding the contract method 0xbe537f43.
//
// Solidity: function getSecurityPolicy() view returns((uint8,uint120,uint120))
func (_Magiceden *MagicedenCallerSession) GetSecurityPolicy() (CollectionSecurityPolicy, error) {
	return _Magiceden.Contract.GetSecurityPolicy(&_Magiceden.CallOpts)
}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Magiceden *MagicedenCaller) GetStageInfo(opts *bind.CallOpts, index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getStageInfo", index)

	if err != nil {
		return *new(IERC721MMintStageInfo), *new(uint32), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721MMintStageInfo)).(*IERC721MMintStageInfo)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Magiceden *MagicedenSession) GetStageInfo(index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	return _Magiceden.Contract.GetStageInfo(&_Magiceden.CallOpts, index)
}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint32,bytes32,uint24,uint64,uint64), uint32, uint256)
func (_Magiceden *MagicedenCallerSession) GetStageInfo(index *big.Int) (IERC721MMintStageInfo, uint32, *big.Int, error) {
	return _Magiceden.Contract.GetStageInfo(&_Magiceden.CallOpts, index)
}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address)
func (_Magiceden *MagicedenCaller) GetTransferValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getTransferValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address)
func (_Magiceden *MagicedenSession) GetTransferValidator() (common.Address, error) {
	return _Magiceden.Contract.GetTransferValidator(&_Magiceden.CallOpts)
}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address)
func (_Magiceden *MagicedenCallerSession) GetTransferValidator() (common.Address, error) {
	return _Magiceden.Contract.GetTransferValidator(&_Magiceden.CallOpts)
}

// GetWhitelistedOperators is a free data retrieval call binding the contract method 0x495c8bf9.
//
// Solidity: function getWhitelistedOperators() view returns(address[])
func (_Magiceden *MagicedenCaller) GetWhitelistedOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getWhitelistedOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedOperators is a free data retrieval call binding the contract method 0x495c8bf9.
//
// Solidity: function getWhitelistedOperators() view returns(address[])
func (_Magiceden *MagicedenSession) GetWhitelistedOperators() ([]common.Address, error) {
	return _Magiceden.Contract.GetWhitelistedOperators(&_Magiceden.CallOpts)
}

// GetWhitelistedOperators is a free data retrieval call binding the contract method 0x495c8bf9.
//
// Solidity: function getWhitelistedOperators() view returns(address[])
func (_Magiceden *MagicedenCallerSession) GetWhitelistedOperators() ([]common.Address, error) {
	return _Magiceden.Contract.GetWhitelistedOperators(&_Magiceden.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Magiceden *MagicedenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Magiceden *MagicedenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Magiceden.Contract.IsApprovedForAll(&_Magiceden.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Magiceden *MagicedenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Magiceden.Contract.IsApprovedForAll(&_Magiceden.CallOpts, owner, operator)
}

// IsContractReceiverPermitted is a free data retrieval call binding the contract method 0x9d645a44.
//
// Solidity: function isContractReceiverPermitted(address receiver) view returns(bool)
func (_Magiceden *MagicedenCaller) IsContractReceiverPermitted(opts *bind.CallOpts, receiver common.Address) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "isContractReceiverPermitted", receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractReceiverPermitted is a free data retrieval call binding the contract method 0x9d645a44.
//
// Solidity: function isContractReceiverPermitted(address receiver) view returns(bool)
func (_Magiceden *MagicedenSession) IsContractReceiverPermitted(receiver common.Address) (bool, error) {
	return _Magiceden.Contract.IsContractReceiverPermitted(&_Magiceden.CallOpts, receiver)
}

// IsContractReceiverPermitted is a free data retrieval call binding the contract method 0x9d645a44.
//
// Solidity: function isContractReceiverPermitted(address receiver) view returns(bool)
func (_Magiceden *MagicedenCallerSession) IsContractReceiverPermitted(receiver common.Address) (bool, error) {
	return _Magiceden.Contract.IsContractReceiverPermitted(&_Magiceden.CallOpts, receiver)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_Magiceden *MagicedenCaller) IsOperatorWhitelisted(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "isOperatorWhitelisted", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_Magiceden *MagicedenSession) IsOperatorWhitelisted(operator common.Address) (bool, error) {
	return _Magiceden.Contract.IsOperatorWhitelisted(&_Magiceden.CallOpts, operator)
}

// IsOperatorWhitelisted is a free data retrieval call binding the contract method 0x2e8da829.
//
// Solidity: function isOperatorWhitelisted(address operator) view returns(bool)
func (_Magiceden *MagicedenCallerSession) IsOperatorWhitelisted(operator common.Address) (bool, error) {
	return _Magiceden.Contract.IsOperatorWhitelisted(&_Magiceden.CallOpts, operator)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x1b25b077.
//
// Solidity: function isTransferAllowed(address caller, address from, address to) view returns(bool)
func (_Magiceden *MagicedenCaller) IsTransferAllowed(opts *bind.CallOpts, caller common.Address, from common.Address, to common.Address) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "isTransferAllowed", caller, from, to)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x1b25b077.
//
// Solidity: function isTransferAllowed(address caller, address from, address to) view returns(bool)
func (_Magiceden *MagicedenSession) IsTransferAllowed(caller common.Address, from common.Address, to common.Address) (bool, error) {
	return _Magiceden.Contract.IsTransferAllowed(&_Magiceden.CallOpts, caller, from, to)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x1b25b077.
//
// Solidity: function isTransferAllowed(address caller, address from, address to) view returns(bool)
func (_Magiceden *MagicedenCallerSession) IsTransferAllowed(caller common.Address, from common.Address, to common.Address) (bool, error) {
	return _Magiceden.Contract.IsTransferAllowed(&_Magiceden.CallOpts, caller, from, to)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Magiceden *MagicedenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Magiceden *MagicedenSession) Name() (string, error) {
	return _Magiceden.Contract.Name(&_Magiceden.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Magiceden *MagicedenCallerSession) Name() (string, error) {
	return _Magiceden.Contract.Name(&_Magiceden.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Magiceden *MagicedenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Magiceden *MagicedenSession) Owner() (common.Address, error) {
	return _Magiceden.Contract.Owner(&_Magiceden.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Magiceden *MagicedenCallerSession) Owner() (common.Address, error) {
	return _Magiceden.Contract.Owner(&_Magiceden.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Magiceden *MagicedenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Magiceden *MagicedenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Magiceden.Contract.OwnerOf(&_Magiceden.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Magiceden *MagicedenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Magiceden.Contract.OwnerOf(&_Magiceden.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Magiceden *MagicedenCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Magiceden *MagicedenSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Magiceden.Contract.RoyaltyInfo(&_Magiceden.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_Magiceden *MagicedenCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Magiceden.Contract.RoyaltyInfo(&_Magiceden.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Magiceden *MagicedenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Magiceden *MagicedenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Magiceden.Contract.SupportsInterface(&_Magiceden.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Magiceden *MagicedenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Magiceden.Contract.SupportsInterface(&_Magiceden.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Magiceden *MagicedenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Magiceden *MagicedenSession) Symbol() (string, error) {
	return _Magiceden.Contract.Symbol(&_Magiceden.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Magiceden *MagicedenCallerSession) Symbol() (string, error) {
	return _Magiceden.Contract.Symbol(&_Magiceden.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Magiceden *MagicedenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Magiceden *MagicedenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Magiceden.Contract.TokenURI(&_Magiceden.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Magiceden *MagicedenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Magiceden.Contract.TokenURI(&_Magiceden.CallOpts, tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Magiceden *MagicedenCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "tokensOfOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Magiceden *MagicedenSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Magiceden.Contract.TokensOfOwner(&_Magiceden.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_Magiceden *MagicedenCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _Magiceden.Contract.TokensOfOwner(&_Magiceden.CallOpts, owner)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Magiceden *MagicedenCaller) TokensOfOwnerIn(opts *bind.CallOpts, owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "tokensOfOwnerIn", owner, start, stop)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Magiceden *MagicedenSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Magiceden.Contract.TokensOfOwnerIn(&_Magiceden.CallOpts, owner, start, stop)
}

// TokensOfOwnerIn is a free data retrieval call binding the contract method 0x99a2557a.
//
// Solidity: function tokensOfOwnerIn(address owner, uint256 start, uint256 stop) view returns(uint256[])
func (_Magiceden *MagicedenCallerSession) TokensOfOwnerIn(owner common.Address, start *big.Int, stop *big.Int) ([]*big.Int, error) {
	return _Magiceden.Contract.TokensOfOwnerIn(&_Magiceden.CallOpts, owner, start, stop)
}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Magiceden *MagicedenCaller) TotalMintedByAddress(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "totalMintedByAddress", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Magiceden *MagicedenSession) TotalMintedByAddress(a common.Address) (*big.Int, error) {
	return _Magiceden.Contract.TotalMintedByAddress(&_Magiceden.CallOpts, a)
}

// TotalMintedByAddress is a free data retrieval call binding the contract method 0x97cf84fc.
//
// Solidity: function totalMintedByAddress(address a) view returns(uint256)
func (_Magiceden *MagicedenCallerSession) TotalMintedByAddress(a common.Address) (*big.Int, error) {
	return _Magiceden.Contract.TotalMintedByAddress(&_Magiceden.CallOpts, a)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Magiceden *MagicedenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Magiceden *MagicedenSession) TotalSupply() (*big.Int, error) {
	return _Magiceden.Contract.TotalSupply(&_Magiceden.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Magiceden *MagicedenCallerSession) TotalSupply() (*big.Int, error) {
	return _Magiceden.Contract.TotalSupply(&_Magiceden.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.Approve(&_Magiceden.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.Approve(&_Magiceden.TransactOpts, to, tokenId)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactor) Crossmint(opts *bind.TransactOpts, qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "crossmint", qty, to, proof, timestamp, signature)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenSession) Crossmint(qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.Crossmint(&_Magiceden.TransactOpts, qty, to, proof, timestamp, signature)
}

// Crossmint is a paid mutator transaction binding the contract method 0x62acbd9a.
//
// Solidity: function crossmint(uint32 qty, address to, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactorSession) Crossmint(qty uint32, to common.Address, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.Crossmint(&_Magiceden.TransactOpts, qty, to, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactor) Mint(opts *bind.TransactOpts, qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "mint", qty, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenSession) Mint(qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.Mint(&_Magiceden.TransactOpts, qty, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xefb6b11f.
//
// Solidity: function mint(uint32 qty, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactorSession) Mint(qty uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.Mint(&_Magiceden.TransactOpts, qty, proof, timestamp, signature)
}

// MintWithLimit is a paid mutator transaction binding the contract method 0x3d6375b2.
//
// Solidity: function mintWithLimit(uint32 qty, uint32 limit, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactor) MintWithLimit(opts *bind.TransactOpts, qty uint32, limit uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "mintWithLimit", qty, limit, proof, timestamp, signature)
}

// MintWithLimit is a paid mutator transaction binding the contract method 0x3d6375b2.
//
// Solidity: function mintWithLimit(uint32 qty, uint32 limit, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenSession) MintWithLimit(qty uint32, limit uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.MintWithLimit(&_Magiceden.TransactOpts, qty, limit, proof, timestamp, signature)
}

// MintWithLimit is a paid mutator transaction binding the contract method 0x3d6375b2.
//
// Solidity: function mintWithLimit(uint32 qty, uint32 limit, bytes32[] proof, uint64 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactorSession) MintWithLimit(qty uint32, limit uint32, proof [][32]byte, timestamp uint64, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.MintWithLimit(&_Magiceden.TransactOpts, qty, limit, proof, timestamp, signature)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Magiceden *MagicedenTransactor) OwnerMint(opts *bind.TransactOpts, qty uint32, to common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "ownerMint", qty, to)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Magiceden *MagicedenSession) OwnerMint(qty uint32, to common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.OwnerMint(&_Magiceden.TransactOpts, qty, to)
}

// OwnerMint is a paid mutator transaction binding the contract method 0xaac5ab1f.
//
// Solidity: function ownerMint(uint32 qty, address to) returns()
func (_Magiceden *MagicedenTransactorSession) OwnerMint(qty uint32, to common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.OwnerMint(&_Magiceden.TransactOpts, qty, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Magiceden *MagicedenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Magiceden *MagicedenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Magiceden.Contract.RenounceOwnership(&_Magiceden.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Magiceden *MagicedenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Magiceden.Contract.RenounceOwnership(&_Magiceden.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SafeTransferFrom(&_Magiceden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SafeTransferFrom(&_Magiceden.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Magiceden *MagicedenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Magiceden *MagicedenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.SafeTransferFrom0(&_Magiceden.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Magiceden *MagicedenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.SafeTransferFrom0(&_Magiceden.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Magiceden *MagicedenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Magiceden *MagicedenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Magiceden.Contract.SetApprovalForAll(&_Magiceden.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Magiceden *MagicedenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Magiceden.Contract.SetApprovalForAll(&_Magiceden.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Magiceden *MagicedenTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Magiceden *MagicedenSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Magiceden.Contract.SetBaseURI(&_Magiceden.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_Magiceden *MagicedenTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _Magiceden.Contract.SetBaseURI(&_Magiceden.TransactOpts, baseURI)
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x1053a815.
//
// Solidity: function setBaseURIPermanent() returns()
func (_Magiceden *MagicedenTransactor) SetBaseURIPermanent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setBaseURIPermanent")
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x1053a815.
//
// Solidity: function setBaseURIPermanent() returns()
func (_Magiceden *MagicedenSession) SetBaseURIPermanent() (*types.Transaction, error) {
	return _Magiceden.Contract.SetBaseURIPermanent(&_Magiceden.TransactOpts)
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x1053a815.
//
// Solidity: function setBaseURIPermanent() returns()
func (_Magiceden *MagicedenTransactorSession) SetBaseURIPermanent() (*types.Transaction, error) {
	return _Magiceden.Contract.SetBaseURIPermanent(&_Magiceden.TransactOpts)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Magiceden *MagicedenTransactor) SetContractURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setContractURI", uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Magiceden *MagicedenSession) SetContractURI(uri string) (*types.Transaction, error) {
	return _Magiceden.Contract.SetContractURI(&_Magiceden.TransactOpts, uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string uri) returns()
func (_Magiceden *MagicedenTransactorSession) SetContractURI(uri string) (*types.Transaction, error) {
	return _Magiceden.Contract.SetContractURI(&_Magiceden.TransactOpts, uri)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Magiceden *MagicedenTransactor) SetCosigner(opts *bind.TransactOpts, cosigner common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setCosigner", cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Magiceden *MagicedenSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.SetCosigner(&_Magiceden.TransactOpts, cosigner)
}

// SetCosigner is a paid mutator transaction binding the contract method 0x02045138.
//
// Solidity: function setCosigner(address cosigner) returns()
func (_Magiceden *MagicedenTransactorSession) SetCosigner(cosigner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.SetCosigner(&_Magiceden.TransactOpts, cosigner)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Magiceden *MagicedenTransactor) SetCrossmintAddress(opts *bind.TransactOpts, crossmintAddress common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setCrossmintAddress", crossmintAddress)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Magiceden *MagicedenSession) SetCrossmintAddress(crossmintAddress common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.SetCrossmintAddress(&_Magiceden.TransactOpts, crossmintAddress)
}

// SetCrossmintAddress is a paid mutator transaction binding the contract method 0x99755624.
//
// Solidity: function setCrossmintAddress(address crossmintAddress) returns()
func (_Magiceden *MagicedenTransactorSession) SetCrossmintAddress(crossmintAddress common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.SetCrossmintAddress(&_Magiceden.TransactOpts, crossmintAddress)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Magiceden *MagicedenTransactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setDefaultRoyalty", receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Magiceden *MagicedenSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetDefaultRoyalty(&_Magiceden.TransactOpts, receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Magiceden *MagicedenTransactorSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetDefaultRoyalty(&_Magiceden.TransactOpts, receiver, feeNumerator)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Magiceden *MagicedenTransactor) SetGlobalWalletLimit(opts *bind.TransactOpts, globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setGlobalWalletLimit", globalWalletLimit)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Magiceden *MagicedenSession) SetGlobalWalletLimit(globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetGlobalWalletLimit(&_Magiceden.TransactOpts, globalWalletLimit)
}

// SetGlobalWalletLimit is a paid mutator transaction binding the contract method 0x372992e4.
//
// Solidity: function setGlobalWalletLimit(uint256 globalWalletLimit) returns()
func (_Magiceden *MagicedenTransactorSession) SetGlobalWalletLimit(globalWalletLimit *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetGlobalWalletLimit(&_Magiceden.TransactOpts, globalWalletLimit)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Magiceden *MagicedenTransactor) SetMaxMintableSupply(opts *bind.TransactOpts, maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setMaxMintableSupply", maxMintableSupply)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Magiceden *MagicedenSession) SetMaxMintableSupply(maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetMaxMintableSupply(&_Magiceden.TransactOpts, maxMintableSupply)
}

// SetMaxMintableSupply is a paid mutator transaction binding the contract method 0xf8d09696.
//
// Solidity: function setMaxMintableSupply(uint256 maxMintableSupply) returns()
func (_Magiceden *MagicedenTransactorSession) SetMaxMintableSupply(maxMintableSupply *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetMaxMintableSupply(&_Magiceden.TransactOpts, maxMintableSupply)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Magiceden *MagicedenTransactor) SetMintable(opts *bind.TransactOpts, mintable bool) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setMintable", mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Magiceden *MagicedenSession) SetMintable(mintable bool) (*types.Transaction, error) {
	return _Magiceden.Contract.SetMintable(&_Magiceden.TransactOpts, mintable)
}

// SetMintable is a paid mutator transaction binding the contract method 0x285d70d4.
//
// Solidity: function setMintable(bool mintable) returns()
func (_Magiceden *MagicedenTransactorSession) SetMintable(mintable bool) (*types.Transaction, error) {
	return _Magiceden.Contract.SetMintable(&_Magiceden.TransactOpts, mintable)
}

// SetStages is a paid mutator transaction binding the contract method 0x8dcdb09d.
//
// Solidity: function setStages((uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Magiceden *MagicedenTransactor) SetStages(opts *bind.TransactOpts, newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setStages", newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0x8dcdb09d.
//
// Solidity: function setStages((uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Magiceden *MagicedenSession) SetStages(newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Magiceden.Contract.SetStages(&_Magiceden.TransactOpts, newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0x8dcdb09d.
//
// Solidity: function setStages((uint80,uint32,bytes32,uint24,uint64,uint64)[] newStages) returns()
func (_Magiceden *MagicedenTransactorSession) SetStages(newStages []IERC721MMintStageInfo) (*types.Transaction, error) {
	return _Magiceden.Contract.SetStages(&_Magiceden.TransactOpts, newStages)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Magiceden *MagicedenTransactor) SetTimestampExpirySeconds(opts *bind.TransactOpts, expiry uint64) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setTimestampExpirySeconds", expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Magiceden *MagicedenSession) SetTimestampExpirySeconds(expiry uint64) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTimestampExpirySeconds(&_Magiceden.TransactOpts, expiry)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0xce2b0ec0.
//
// Solidity: function setTimestampExpirySeconds(uint64 expiry) returns()
func (_Magiceden *MagicedenTransactorSession) SetTimestampExpirySeconds(expiry uint64) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTimestampExpirySeconds(&_Magiceden.TransactOpts, expiry)
}

// SetToCustomSecurityPolicy is a paid mutator transaction binding the contract method 0x61347162.
//
// Solidity: function setToCustomSecurityPolicy(uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Magiceden *MagicedenTransactor) SetToCustomSecurityPolicy(opts *bind.TransactOpts, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setToCustomSecurityPolicy", level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomSecurityPolicy is a paid mutator transaction binding the contract method 0x61347162.
//
// Solidity: function setToCustomSecurityPolicy(uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Magiceden *MagicedenSession) SetToCustomSecurityPolicy(level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetToCustomSecurityPolicy(&_Magiceden.TransactOpts, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomSecurityPolicy is a paid mutator transaction binding the contract method 0x61347162.
//
// Solidity: function setToCustomSecurityPolicy(uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Magiceden *MagicedenTransactorSession) SetToCustomSecurityPolicy(level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetToCustomSecurityPolicy(&_Magiceden.TransactOpts, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomValidatorAndSecurityPolicy is a paid mutator transaction binding the contract method 0xfd762d92.
//
// Solidity: function setToCustomValidatorAndSecurityPolicy(address validator, uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Magiceden *MagicedenTransactor) SetToCustomValidatorAndSecurityPolicy(opts *bind.TransactOpts, validator common.Address, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setToCustomValidatorAndSecurityPolicy", validator, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomValidatorAndSecurityPolicy is a paid mutator transaction binding the contract method 0xfd762d92.
//
// Solidity: function setToCustomValidatorAndSecurityPolicy(address validator, uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Magiceden *MagicedenSession) SetToCustomValidatorAndSecurityPolicy(validator common.Address, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetToCustomValidatorAndSecurityPolicy(&_Magiceden.TransactOpts, validator, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToCustomValidatorAndSecurityPolicy is a paid mutator transaction binding the contract method 0xfd762d92.
//
// Solidity: function setToCustomValidatorAndSecurityPolicy(address validator, uint8 level, uint120 operatorWhitelistId, uint120 permittedContractReceiversAllowlistId) returns()
func (_Magiceden *MagicedenTransactorSession) SetToCustomValidatorAndSecurityPolicy(validator common.Address, level uint8, operatorWhitelistId *big.Int, permittedContractReceiversAllowlistId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetToCustomValidatorAndSecurityPolicy(&_Magiceden.TransactOpts, validator, level, operatorWhitelistId, permittedContractReceiversAllowlistId)
}

// SetToDefaultSecurityPolicy is a paid mutator transaction binding the contract method 0x6c3b8699.
//
// Solidity: function setToDefaultSecurityPolicy() returns()
func (_Magiceden *MagicedenTransactor) SetToDefaultSecurityPolicy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setToDefaultSecurityPolicy")
}

// SetToDefaultSecurityPolicy is a paid mutator transaction binding the contract method 0x6c3b8699.
//
// Solidity: function setToDefaultSecurityPolicy() returns()
func (_Magiceden *MagicedenSession) SetToDefaultSecurityPolicy() (*types.Transaction, error) {
	return _Magiceden.Contract.SetToDefaultSecurityPolicy(&_Magiceden.TransactOpts)
}

// SetToDefaultSecurityPolicy is a paid mutator transaction binding the contract method 0x6c3b8699.
//
// Solidity: function setToDefaultSecurityPolicy() returns()
func (_Magiceden *MagicedenTransactorSession) SetToDefaultSecurityPolicy() (*types.Transaction, error) {
	return _Magiceden.Contract.SetToDefaultSecurityPolicy(&_Magiceden.TransactOpts)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Magiceden *MagicedenTransactor) SetTokenRoyalty(opts *bind.TransactOpts, tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setTokenRoyalty", tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Magiceden *MagicedenSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTokenRoyalty(&_Magiceden.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Magiceden *MagicedenTransactorSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTokenRoyalty(&_Magiceden.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Magiceden *MagicedenTransactor) SetTokenURISuffix(opts *bind.TransactOpts, suffix string) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setTokenURISuffix", suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Magiceden *MagicedenSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTokenURISuffix(&_Magiceden.TransactOpts, suffix)
}

// SetTokenURISuffix is a paid mutator transaction binding the contract method 0xa9852bfb.
//
// Solidity: function setTokenURISuffix(string suffix) returns()
func (_Magiceden *MagicedenTransactorSession) SetTokenURISuffix(suffix string) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTokenURISuffix(&_Magiceden.TransactOpts, suffix)
}

// SetTransferValidator is a paid mutator transaction binding the contract method 0xa9fc664e.
//
// Solidity: function setTransferValidator(address transferValidator_) returns()
func (_Magiceden *MagicedenTransactor) SetTransferValidator(opts *bind.TransactOpts, transferValidator_ common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setTransferValidator", transferValidator_)
}

// SetTransferValidator is a paid mutator transaction binding the contract method 0xa9fc664e.
//
// Solidity: function setTransferValidator(address transferValidator_) returns()
func (_Magiceden *MagicedenSession) SetTransferValidator(transferValidator_ common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTransferValidator(&_Magiceden.TransactOpts, transferValidator_)
}

// SetTransferValidator is a paid mutator transaction binding the contract method 0xa9fc664e.
//
// Solidity: function setTransferValidator(address transferValidator_) returns()
func (_Magiceden *MagicedenTransactorSession) SetTransferValidator(transferValidator_ common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTransferValidator(&_Magiceden.TransactOpts, transferValidator_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.TransferFrom(&_Magiceden.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Magiceden *MagicedenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.TransferFrom(&_Magiceden.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Magiceden *MagicedenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Magiceden *MagicedenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.TransferOwnership(&_Magiceden.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Magiceden *MagicedenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.TransferOwnership(&_Magiceden.TransactOpts, newOwner)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x73e1607e.
//
// Solidity: function updateStage(uint256 index, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds) returns()
func (_Magiceden *MagicedenTransactor) UpdateStage(opts *bind.TransactOpts, index *big.Int, price *big.Int, walletLimit uint32, merkleRoot [32]byte, maxStageSupply *big.Int, startTimeUnixSeconds uint64, endTimeUnixSeconds uint64) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "updateStage", index, price, walletLimit, merkleRoot, maxStageSupply, startTimeUnixSeconds, endTimeUnixSeconds)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x73e1607e.
//
// Solidity: function updateStage(uint256 index, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds) returns()
func (_Magiceden *MagicedenSession) UpdateStage(index *big.Int, price *big.Int, walletLimit uint32, merkleRoot [32]byte, maxStageSupply *big.Int, startTimeUnixSeconds uint64, endTimeUnixSeconds uint64) (*types.Transaction, error) {
	return _Magiceden.Contract.UpdateStage(&_Magiceden.TransactOpts, index, price, walletLimit, merkleRoot, maxStageSupply, startTimeUnixSeconds, endTimeUnixSeconds)
}

// UpdateStage is a paid mutator transaction binding the contract method 0x73e1607e.
//
// Solidity: function updateStage(uint256 index, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds) returns()
func (_Magiceden *MagicedenTransactorSession) UpdateStage(index *big.Int, price *big.Int, walletLimit uint32, merkleRoot [32]byte, maxStageSupply *big.Int, startTimeUnixSeconds uint64, endTimeUnixSeconds uint64) (*types.Transaction, error) {
	return _Magiceden.Contract.UpdateStage(&_Magiceden.TransactOpts, index, price, walletLimit, merkleRoot, maxStageSupply, startTimeUnixSeconds, endTimeUnixSeconds)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Magiceden *MagicedenTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Magiceden *MagicedenSession) Withdraw() (*types.Transaction, error) {
	return _Magiceden.Contract.Withdraw(&_Magiceden.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Magiceden *MagicedenTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Magiceden.Contract.Withdraw(&_Magiceden.TransactOpts)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2ed6d5e8.
//
// Solidity: function withdrawERC20() returns()
func (_Magiceden *MagicedenTransactor) WithdrawERC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "withdrawERC20")
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2ed6d5e8.
//
// Solidity: function withdrawERC20() returns()
func (_Magiceden *MagicedenSession) WithdrawERC20() (*types.Transaction, error) {
	return _Magiceden.Contract.WithdrawERC20(&_Magiceden.TransactOpts)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2ed6d5e8.
//
// Solidity: function withdrawERC20() returns()
func (_Magiceden *MagicedenTransactorSession) WithdrawERC20() (*types.Transaction, error) {
	return _Magiceden.Contract.WithdrawERC20(&_Magiceden.TransactOpts)
}

// MagicedenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Magiceden contract.
type MagicedenApprovalIterator struct {
	Event *MagicedenApproval // Event containing the contract specifics and raw log

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
func (it *MagicedenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenApproval)
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
		it.Event = new(MagicedenApproval)
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
func (it *MagicedenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenApproval represents a Approval event raised by the Magiceden contract.
type MagicedenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Magiceden *MagicedenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MagicedenApprovalIterator, error) {

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

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenApprovalIterator{contract: _Magiceden.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Magiceden *MagicedenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MagicedenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenApproval)
				if err := _Magiceden.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Magiceden *MagicedenFilterer) ParseApproval(log types.Log) (*MagicedenApproval, error) {
	event := new(MagicedenApproval)
	if err := _Magiceden.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Magiceden contract.
type MagicedenApprovalForAllIterator struct {
	Event *MagicedenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MagicedenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenApprovalForAll)
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
		it.Event = new(MagicedenApprovalForAll)
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
func (it *MagicedenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenApprovalForAll represents a ApprovalForAll event raised by the Magiceden contract.
type MagicedenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Magiceden *MagicedenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MagicedenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenApprovalForAllIterator{contract: _Magiceden.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Magiceden *MagicedenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MagicedenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenApprovalForAll)
				if err := _Magiceden.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Magiceden *MagicedenFilterer) ParseApprovalForAll(log types.Log) (*MagicedenApprovalForAll, error) {
	event := new(MagicedenApprovalForAll)
	if err := _Magiceden.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Magiceden contract.
type MagicedenConsecutiveTransferIterator struct {
	Event *MagicedenConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *MagicedenConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenConsecutiveTransfer)
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
		it.Event = new(MagicedenConsecutiveTransfer)
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
func (it *MagicedenConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Magiceden contract.
type MagicedenConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Magiceden *MagicedenFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*MagicedenConsecutiveTransferIterator, error) {

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

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenConsecutiveTransferIterator{contract: _Magiceden.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Magiceden *MagicedenFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *MagicedenConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenConsecutiveTransfer)
				if err := _Magiceden.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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
func (_Magiceden *MagicedenFilterer) ParseConsecutiveTransfer(log types.Log) (*MagicedenConsecutiveTransfer, error) {
	event := new(MagicedenConsecutiveTransfer)
	if err := _Magiceden.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenDefaultRoyaltySetIterator is returned from FilterDefaultRoyaltySet and is used to iterate over the raw logs and unpacked data for DefaultRoyaltySet events raised by the Magiceden contract.
type MagicedenDefaultRoyaltySetIterator struct {
	Event *MagicedenDefaultRoyaltySet // Event containing the contract specifics and raw log

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
func (it *MagicedenDefaultRoyaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenDefaultRoyaltySet)
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
		it.Event = new(MagicedenDefaultRoyaltySet)
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
func (it *MagicedenDefaultRoyaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenDefaultRoyaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenDefaultRoyaltySet represents a DefaultRoyaltySet event raised by the Magiceden contract.
type MagicedenDefaultRoyaltySet struct {
	Receiver     common.Address
	FeeNumerator *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyaltySet is a free log retrieval operation binding the contract event 0x8a8bae378cb731c5c40b632330c6836c2f916f48edb967699c86736f9a6a76ef.
//
// Solidity: event DefaultRoyaltySet(address indexed receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) FilterDefaultRoyaltySet(opts *bind.FilterOpts, receiver []common.Address) (*MagicedenDefaultRoyaltySetIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "DefaultRoyaltySet", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenDefaultRoyaltySetIterator{contract: _Magiceden.contract, event: "DefaultRoyaltySet", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyaltySet is a free log subscription operation binding the contract event 0x8a8bae378cb731c5c40b632330c6836c2f916f48edb967699c86736f9a6a76ef.
//
// Solidity: event DefaultRoyaltySet(address indexed receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) WatchDefaultRoyaltySet(opts *bind.WatchOpts, sink chan<- *MagicedenDefaultRoyaltySet, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "DefaultRoyaltySet", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenDefaultRoyaltySet)
				if err := _Magiceden.contract.UnpackLog(event, "DefaultRoyaltySet", log); err != nil {
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

// ParseDefaultRoyaltySet is a log parse operation binding the contract event 0x8a8bae378cb731c5c40b632330c6836c2f916f48edb967699c86736f9a6a76ef.
//
// Solidity: event DefaultRoyaltySet(address indexed receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) ParseDefaultRoyaltySet(log types.Log) (*MagicedenDefaultRoyaltySet, error) {
	event := new(MagicedenDefaultRoyaltySet)
	if err := _Magiceden.contract.UnpackLog(event, "DefaultRoyaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Magiceden contract.
type MagicedenOwnershipTransferredIterator struct {
	Event *MagicedenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MagicedenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenOwnershipTransferred)
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
		it.Event = new(MagicedenOwnershipTransferred)
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
func (it *MagicedenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenOwnershipTransferred represents a OwnershipTransferred event raised by the Magiceden contract.
type MagicedenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Magiceden *MagicedenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MagicedenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenOwnershipTransferredIterator{contract: _Magiceden.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Magiceden *MagicedenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MagicedenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenOwnershipTransferred)
				if err := _Magiceden.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Magiceden *MagicedenFilterer) ParseOwnershipTransferred(log types.Log) (*MagicedenOwnershipTransferred, error) {
	event := new(MagicedenOwnershipTransferred)
	if err := _Magiceden.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenPermanentBaseURIIterator is returned from FilterPermanentBaseURI and is used to iterate over the raw logs and unpacked data for PermanentBaseURI events raised by the Magiceden contract.
type MagicedenPermanentBaseURIIterator struct {
	Event *MagicedenPermanentBaseURI // Event containing the contract specifics and raw log

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
func (it *MagicedenPermanentBaseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenPermanentBaseURI)
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
		it.Event = new(MagicedenPermanentBaseURI)
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
func (it *MagicedenPermanentBaseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenPermanentBaseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenPermanentBaseURI represents a PermanentBaseURI event raised by the Magiceden contract.
type MagicedenPermanentBaseURI struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPermanentBaseURI is a free log retrieval operation binding the contract event 0xc6a6c2b165e62c9d37fc51a18ed76e5be22304bc1d337877c98f31c23e40b0f5.
//
// Solidity: event PermanentBaseURI(string baseURI)
func (_Magiceden *MagicedenFilterer) FilterPermanentBaseURI(opts *bind.FilterOpts) (*MagicedenPermanentBaseURIIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "PermanentBaseURI")
	if err != nil {
		return nil, err
	}
	return &MagicedenPermanentBaseURIIterator{contract: _Magiceden.contract, event: "PermanentBaseURI", logs: logs, sub: sub}, nil
}

// WatchPermanentBaseURI is a free log subscription operation binding the contract event 0xc6a6c2b165e62c9d37fc51a18ed76e5be22304bc1d337877c98f31c23e40b0f5.
//
// Solidity: event PermanentBaseURI(string baseURI)
func (_Magiceden *MagicedenFilterer) WatchPermanentBaseURI(opts *bind.WatchOpts, sink chan<- *MagicedenPermanentBaseURI) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "PermanentBaseURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenPermanentBaseURI)
				if err := _Magiceden.contract.UnpackLog(event, "PermanentBaseURI", log); err != nil {
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

// ParsePermanentBaseURI is a log parse operation binding the contract event 0xc6a6c2b165e62c9d37fc51a18ed76e5be22304bc1d337877c98f31c23e40b0f5.
//
// Solidity: event PermanentBaseURI(string baseURI)
func (_Magiceden *MagicedenFilterer) ParsePermanentBaseURI(log types.Log) (*MagicedenPermanentBaseURI, error) {
	event := new(MagicedenPermanentBaseURI)
	if err := _Magiceden.contract.UnpackLog(event, "PermanentBaseURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetActiveStageIterator is returned from FilterSetActiveStage and is used to iterate over the raw logs and unpacked data for SetActiveStage events raised by the Magiceden contract.
type MagicedenSetActiveStageIterator struct {
	Event *MagicedenSetActiveStage // Event containing the contract specifics and raw log

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
func (it *MagicedenSetActiveStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetActiveStage)
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
		it.Event = new(MagicedenSetActiveStage)
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
func (it *MagicedenSetActiveStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetActiveStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetActiveStage represents a SetActiveStage event raised by the Magiceden contract.
type MagicedenSetActiveStage struct {
	ActiveStage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetActiveStage is a free log retrieval operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Magiceden *MagicedenFilterer) FilterSetActiveStage(opts *bind.FilterOpts) (*MagicedenSetActiveStageIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetActiveStage")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetActiveStageIterator{contract: _Magiceden.contract, event: "SetActiveStage", logs: logs, sub: sub}, nil
}

// WatchSetActiveStage is a free log subscription operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Magiceden *MagicedenFilterer) WatchSetActiveStage(opts *bind.WatchOpts, sink chan<- *MagicedenSetActiveStage) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetActiveStage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetActiveStage)
				if err := _Magiceden.contract.UnpackLog(event, "SetActiveStage", log); err != nil {
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

// ParseSetActiveStage is a log parse operation binding the contract event 0x160d6de2c069c3adf7f4c1252236d0b325c0e3eb963ddb10c67a81aadf9a3058.
//
// Solidity: event SetActiveStage(uint256 activeStage)
func (_Magiceden *MagicedenFilterer) ParseSetActiveStage(log types.Log) (*MagicedenSetActiveStage, error) {
	event := new(MagicedenSetActiveStage)
	if err := _Magiceden.contract.UnpackLog(event, "SetActiveStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetBaseURIIterator is returned from FilterSetBaseURI and is used to iterate over the raw logs and unpacked data for SetBaseURI events raised by the Magiceden contract.
type MagicedenSetBaseURIIterator struct {
	Event *MagicedenSetBaseURI // Event containing the contract specifics and raw log

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
func (it *MagicedenSetBaseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetBaseURI)
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
		it.Event = new(MagicedenSetBaseURI)
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
func (it *MagicedenSetBaseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetBaseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetBaseURI represents a SetBaseURI event raised by the Magiceden contract.
type MagicedenSetBaseURI struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBaseURI is a free log retrieval operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Magiceden *MagicedenFilterer) FilterSetBaseURI(opts *bind.FilterOpts) (*MagicedenSetBaseURIIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetBaseURI")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetBaseURIIterator{contract: _Magiceden.contract, event: "SetBaseURI", logs: logs, sub: sub}, nil
}

// WatchSetBaseURI is a free log subscription operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Magiceden *MagicedenFilterer) WatchSetBaseURI(opts *bind.WatchOpts, sink chan<- *MagicedenSetBaseURI) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetBaseURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetBaseURI)
				if err := _Magiceden.contract.UnpackLog(event, "SetBaseURI", log); err != nil {
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

// ParseSetBaseURI is a log parse operation binding the contract event 0x23c8c9488efebfd474e85a7956de6f39b17c7ab88502d42a623db2d8e382bbaa.
//
// Solidity: event SetBaseURI(string baseURI)
func (_Magiceden *MagicedenFilterer) ParseSetBaseURI(log types.Log) (*MagicedenSetBaseURI, error) {
	event := new(MagicedenSetBaseURI)
	if err := _Magiceden.contract.UnpackLog(event, "SetBaseURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetCosignerIterator is returned from FilterSetCosigner and is used to iterate over the raw logs and unpacked data for SetCosigner events raised by the Magiceden contract.
type MagicedenSetCosignerIterator struct {
	Event *MagicedenSetCosigner // Event containing the contract specifics and raw log

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
func (it *MagicedenSetCosignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetCosigner)
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
		it.Event = new(MagicedenSetCosigner)
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
func (it *MagicedenSetCosignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetCosignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetCosigner represents a SetCosigner event raised by the Magiceden contract.
type MagicedenSetCosigner struct {
	Cosigner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetCosigner is a free log retrieval operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Magiceden *MagicedenFilterer) FilterSetCosigner(opts *bind.FilterOpts) (*MagicedenSetCosignerIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetCosigner")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetCosignerIterator{contract: _Magiceden.contract, event: "SetCosigner", logs: logs, sub: sub}, nil
}

// WatchSetCosigner is a free log subscription operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Magiceden *MagicedenFilterer) WatchSetCosigner(opts *bind.WatchOpts, sink chan<- *MagicedenSetCosigner) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetCosigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetCosigner)
				if err := _Magiceden.contract.UnpackLog(event, "SetCosigner", log); err != nil {
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

// ParseSetCosigner is a log parse operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address cosigner)
func (_Magiceden *MagicedenFilterer) ParseSetCosigner(log types.Log) (*MagicedenSetCosigner, error) {
	event := new(MagicedenSetCosigner)
	if err := _Magiceden.contract.UnpackLog(event, "SetCosigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetCrossmintAddressIterator is returned from FilterSetCrossmintAddress and is used to iterate over the raw logs and unpacked data for SetCrossmintAddress events raised by the Magiceden contract.
type MagicedenSetCrossmintAddressIterator struct {
	Event *MagicedenSetCrossmintAddress // Event containing the contract specifics and raw log

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
func (it *MagicedenSetCrossmintAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetCrossmintAddress)
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
		it.Event = new(MagicedenSetCrossmintAddress)
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
func (it *MagicedenSetCrossmintAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetCrossmintAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetCrossmintAddress represents a SetCrossmintAddress event raised by the Magiceden contract.
type MagicedenSetCrossmintAddress struct {
	CrossmintAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetCrossmintAddress is a free log retrieval operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Magiceden *MagicedenFilterer) FilterSetCrossmintAddress(opts *bind.FilterOpts) (*MagicedenSetCrossmintAddressIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetCrossmintAddress")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetCrossmintAddressIterator{contract: _Magiceden.contract, event: "SetCrossmintAddress", logs: logs, sub: sub}, nil
}

// WatchSetCrossmintAddress is a free log subscription operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Magiceden *MagicedenFilterer) WatchSetCrossmintAddress(opts *bind.WatchOpts, sink chan<- *MagicedenSetCrossmintAddress) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetCrossmintAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetCrossmintAddress)
				if err := _Magiceden.contract.UnpackLog(event, "SetCrossmintAddress", log); err != nil {
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

// ParseSetCrossmintAddress is a log parse operation binding the contract event 0xf477d93c015f2a73c2ccc5ed37078d12123b80fc5d12e0014c60b913bc1a1ec4.
//
// Solidity: event SetCrossmintAddress(address crossmintAddress)
func (_Magiceden *MagicedenFilterer) ParseSetCrossmintAddress(log types.Log) (*MagicedenSetCrossmintAddress, error) {
	event := new(MagicedenSetCrossmintAddress)
	if err := _Magiceden.contract.UnpackLog(event, "SetCrossmintAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetGlobalWalletLimitIterator is returned from FilterSetGlobalWalletLimit and is used to iterate over the raw logs and unpacked data for SetGlobalWalletLimit events raised by the Magiceden contract.
type MagicedenSetGlobalWalletLimitIterator struct {
	Event *MagicedenSetGlobalWalletLimit // Event containing the contract specifics and raw log

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
func (it *MagicedenSetGlobalWalletLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetGlobalWalletLimit)
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
		it.Event = new(MagicedenSetGlobalWalletLimit)
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
func (it *MagicedenSetGlobalWalletLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetGlobalWalletLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetGlobalWalletLimit represents a SetGlobalWalletLimit event raised by the Magiceden contract.
type MagicedenSetGlobalWalletLimit struct {
	GlobalWalletLimit *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalWalletLimit is a free log retrieval operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Magiceden *MagicedenFilterer) FilterSetGlobalWalletLimit(opts *bind.FilterOpts) (*MagicedenSetGlobalWalletLimitIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetGlobalWalletLimit")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetGlobalWalletLimitIterator{contract: _Magiceden.contract, event: "SetGlobalWalletLimit", logs: logs, sub: sub}, nil
}

// WatchSetGlobalWalletLimit is a free log subscription operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Magiceden *MagicedenFilterer) WatchSetGlobalWalletLimit(opts *bind.WatchOpts, sink chan<- *MagicedenSetGlobalWalletLimit) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetGlobalWalletLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetGlobalWalletLimit)
				if err := _Magiceden.contract.UnpackLog(event, "SetGlobalWalletLimit", log); err != nil {
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

// ParseSetGlobalWalletLimit is a log parse operation binding the contract event 0x5307de8ad7d34d5ddfd5171435c143bdc645493980f453eb5d7cdb3e494a1b35.
//
// Solidity: event SetGlobalWalletLimit(uint256 globalWalletLimit)
func (_Magiceden *MagicedenFilterer) ParseSetGlobalWalletLimit(log types.Log) (*MagicedenSetGlobalWalletLimit, error) {
	event := new(MagicedenSetGlobalWalletLimit)
	if err := _Magiceden.contract.UnpackLog(event, "SetGlobalWalletLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetMaxMintableSupplyIterator is returned from FilterSetMaxMintableSupply and is used to iterate over the raw logs and unpacked data for SetMaxMintableSupply events raised by the Magiceden contract.
type MagicedenSetMaxMintableSupplyIterator struct {
	Event *MagicedenSetMaxMintableSupply // Event containing the contract specifics and raw log

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
func (it *MagicedenSetMaxMintableSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetMaxMintableSupply)
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
		it.Event = new(MagicedenSetMaxMintableSupply)
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
func (it *MagicedenSetMaxMintableSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetMaxMintableSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetMaxMintableSupply represents a SetMaxMintableSupply event raised by the Magiceden contract.
type MagicedenSetMaxMintableSupply struct {
	MaxMintableSupply *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMaxMintableSupply is a free log retrieval operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Magiceden *MagicedenFilterer) FilterSetMaxMintableSupply(opts *bind.FilterOpts) (*MagicedenSetMaxMintableSupplyIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetMaxMintableSupply")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetMaxMintableSupplyIterator{contract: _Magiceden.contract, event: "SetMaxMintableSupply", logs: logs, sub: sub}, nil
}

// WatchSetMaxMintableSupply is a free log subscription operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Magiceden *MagicedenFilterer) WatchSetMaxMintableSupply(opts *bind.WatchOpts, sink chan<- *MagicedenSetMaxMintableSupply) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetMaxMintableSupply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetMaxMintableSupply)
				if err := _Magiceden.contract.UnpackLog(event, "SetMaxMintableSupply", log); err != nil {
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

// ParseSetMaxMintableSupply is a log parse operation binding the contract event 0xc7bbc2b288fc13314546ea4aa51f6bcf71b7ba4740beeb3d32e9acef57b6668a.
//
// Solidity: event SetMaxMintableSupply(uint256 maxMintableSupply)
func (_Magiceden *MagicedenFilterer) ParseSetMaxMintableSupply(log types.Log) (*MagicedenSetMaxMintableSupply, error) {
	event := new(MagicedenSetMaxMintableSupply)
	if err := _Magiceden.contract.UnpackLog(event, "SetMaxMintableSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetMintCurrencyIterator is returned from FilterSetMintCurrency and is used to iterate over the raw logs and unpacked data for SetMintCurrency events raised by the Magiceden contract.
type MagicedenSetMintCurrencyIterator struct {
	Event *MagicedenSetMintCurrency // Event containing the contract specifics and raw log

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
func (it *MagicedenSetMintCurrencyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetMintCurrency)
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
		it.Event = new(MagicedenSetMintCurrency)
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
func (it *MagicedenSetMintCurrencyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetMintCurrencyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetMintCurrency represents a SetMintCurrency event raised by the Magiceden contract.
type MagicedenSetMintCurrency struct {
	MintCurrency common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetMintCurrency is a free log retrieval operation binding the contract event 0x20e0479aeca91ed524bd7f643c7d75dc55805d7107e589c5450d47eb0233f267.
//
// Solidity: event SetMintCurrency(address mintCurrency)
func (_Magiceden *MagicedenFilterer) FilterSetMintCurrency(opts *bind.FilterOpts) (*MagicedenSetMintCurrencyIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetMintCurrency")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetMintCurrencyIterator{contract: _Magiceden.contract, event: "SetMintCurrency", logs: logs, sub: sub}, nil
}

// WatchSetMintCurrency is a free log subscription operation binding the contract event 0x20e0479aeca91ed524bd7f643c7d75dc55805d7107e589c5450d47eb0233f267.
//
// Solidity: event SetMintCurrency(address mintCurrency)
func (_Magiceden *MagicedenFilterer) WatchSetMintCurrency(opts *bind.WatchOpts, sink chan<- *MagicedenSetMintCurrency) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetMintCurrency")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetMintCurrency)
				if err := _Magiceden.contract.UnpackLog(event, "SetMintCurrency", log); err != nil {
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

// ParseSetMintCurrency is a log parse operation binding the contract event 0x20e0479aeca91ed524bd7f643c7d75dc55805d7107e589c5450d47eb0233f267.
//
// Solidity: event SetMintCurrency(address mintCurrency)
func (_Magiceden *MagicedenFilterer) ParseSetMintCurrency(log types.Log) (*MagicedenSetMintCurrency, error) {
	event := new(MagicedenSetMintCurrency)
	if err := _Magiceden.contract.UnpackLog(event, "SetMintCurrency", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetMintableIterator is returned from FilterSetMintable and is used to iterate over the raw logs and unpacked data for SetMintable events raised by the Magiceden contract.
type MagicedenSetMintableIterator struct {
	Event *MagicedenSetMintable // Event containing the contract specifics and raw log

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
func (it *MagicedenSetMintableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetMintable)
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
		it.Event = new(MagicedenSetMintable)
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
func (it *MagicedenSetMintableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetMintableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetMintable represents a SetMintable event raised by the Magiceden contract.
type MagicedenSetMintable struct {
	Mintable bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMintable is a free log retrieval operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Magiceden *MagicedenFilterer) FilterSetMintable(opts *bind.FilterOpts) (*MagicedenSetMintableIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetMintable")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetMintableIterator{contract: _Magiceden.contract, event: "SetMintable", logs: logs, sub: sub}, nil
}

// WatchSetMintable is a free log subscription operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Magiceden *MagicedenFilterer) WatchSetMintable(opts *bind.WatchOpts, sink chan<- *MagicedenSetMintable) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetMintable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetMintable)
				if err := _Magiceden.contract.UnpackLog(event, "SetMintable", log); err != nil {
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

// ParseSetMintable is a log parse operation binding the contract event 0xe717a2bfc51e250b028aaac5eb448e76f4df26b9609956782bff49097bb792cf.
//
// Solidity: event SetMintable(bool mintable)
func (_Magiceden *MagicedenFilterer) ParseSetMintable(log types.Log) (*MagicedenSetMintable, error) {
	event := new(MagicedenSetMintable)
	if err := _Magiceden.contract.UnpackLog(event, "SetMintable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetTimestampExpirySecondsIterator is returned from FilterSetTimestampExpirySeconds and is used to iterate over the raw logs and unpacked data for SetTimestampExpirySeconds events raised by the Magiceden contract.
type MagicedenSetTimestampExpirySecondsIterator struct {
	Event *MagicedenSetTimestampExpirySeconds // Event containing the contract specifics and raw log

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
func (it *MagicedenSetTimestampExpirySecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetTimestampExpirySeconds)
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
		it.Event = new(MagicedenSetTimestampExpirySeconds)
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
func (it *MagicedenSetTimestampExpirySecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetTimestampExpirySecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetTimestampExpirySeconds represents a SetTimestampExpirySeconds event raised by the Magiceden contract.
type MagicedenSetTimestampExpirySeconds struct {
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTimestampExpirySeconds is a free log retrieval operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Magiceden *MagicedenFilterer) FilterSetTimestampExpirySeconds(opts *bind.FilterOpts) (*MagicedenSetTimestampExpirySecondsIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetTimestampExpirySeconds")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetTimestampExpirySecondsIterator{contract: _Magiceden.contract, event: "SetTimestampExpirySeconds", logs: logs, sub: sub}, nil
}

// WatchSetTimestampExpirySeconds is a free log subscription operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Magiceden *MagicedenFilterer) WatchSetTimestampExpirySeconds(opts *bind.WatchOpts, sink chan<- *MagicedenSetTimestampExpirySeconds) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetTimestampExpirySeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetTimestampExpirySeconds)
				if err := _Magiceden.contract.UnpackLog(event, "SetTimestampExpirySeconds", log); err != nil {
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

// ParseSetTimestampExpirySeconds is a log parse operation binding the contract event 0x41b9126ccd8cb4505310c40a376055b5ef246bd4c9214de02af31ef4f26b1b5f.
//
// Solidity: event SetTimestampExpirySeconds(uint64 expiry)
func (_Magiceden *MagicedenFilterer) ParseSetTimestampExpirySeconds(log types.Log) (*MagicedenSetTimestampExpirySeconds, error) {
	event := new(MagicedenSetTimestampExpirySeconds)
	if err := _Magiceden.contract.UnpackLog(event, "SetTimestampExpirySeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenTokenRoyaltySetIterator is returned from FilterTokenRoyaltySet and is used to iterate over the raw logs and unpacked data for TokenRoyaltySet events raised by the Magiceden contract.
type MagicedenTokenRoyaltySetIterator struct {
	Event *MagicedenTokenRoyaltySet // Event containing the contract specifics and raw log

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
func (it *MagicedenTokenRoyaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenTokenRoyaltySet)
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
		it.Event = new(MagicedenTokenRoyaltySet)
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
func (it *MagicedenTokenRoyaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenTokenRoyaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenTokenRoyaltySet represents a TokenRoyaltySet event raised by the Magiceden contract.
type MagicedenTokenRoyaltySet struct {
	TokenId      *big.Int
	Receiver     common.Address
	FeeNumerator *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRoyaltySet is a free log retrieval operation binding the contract event 0x7f5b076c952c0ec86e5425963c1326dd0f03a3595c19f81d765e8ff559a6e33c.
//
// Solidity: event TokenRoyaltySet(uint256 indexed tokenId, address indexed receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) FilterTokenRoyaltySet(opts *bind.FilterOpts, tokenId []*big.Int, receiver []common.Address) (*MagicedenTokenRoyaltySetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "TokenRoyaltySet", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenTokenRoyaltySetIterator{contract: _Magiceden.contract, event: "TokenRoyaltySet", logs: logs, sub: sub}, nil
}

// WatchTokenRoyaltySet is a free log subscription operation binding the contract event 0x7f5b076c952c0ec86e5425963c1326dd0f03a3595c19f81d765e8ff559a6e33c.
//
// Solidity: event TokenRoyaltySet(uint256 indexed tokenId, address indexed receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) WatchTokenRoyaltySet(opts *bind.WatchOpts, sink chan<- *MagicedenTokenRoyaltySet, tokenId []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "TokenRoyaltySet", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenTokenRoyaltySet)
				if err := _Magiceden.contract.UnpackLog(event, "TokenRoyaltySet", log); err != nil {
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

// ParseTokenRoyaltySet is a log parse operation binding the contract event 0x7f5b076c952c0ec86e5425963c1326dd0f03a3595c19f81d765e8ff559a6e33c.
//
// Solidity: event TokenRoyaltySet(uint256 indexed tokenId, address indexed receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) ParseTokenRoyaltySet(log types.Log) (*MagicedenTokenRoyaltySet, error) {
	event := new(MagicedenTokenRoyaltySet)
	if err := _Magiceden.contract.UnpackLog(event, "TokenRoyaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Magiceden contract.
type MagicedenTransferIterator struct {
	Event *MagicedenTransfer // Event containing the contract specifics and raw log

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
func (it *MagicedenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenTransfer)
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
		it.Event = new(MagicedenTransfer)
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
func (it *MagicedenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenTransfer represents a Transfer event raised by the Magiceden contract.
type MagicedenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Magiceden *MagicedenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MagicedenTransferIterator, error) {

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

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenTransferIterator{contract: _Magiceden.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Magiceden *MagicedenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MagicedenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenTransfer)
				if err := _Magiceden.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Magiceden *MagicedenFilterer) ParseTransfer(log types.Log) (*MagicedenTransfer, error) {
	event := new(MagicedenTransfer)
	if err := _Magiceden.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenTransferValidatorUpdatedIterator is returned from FilterTransferValidatorUpdated and is used to iterate over the raw logs and unpacked data for TransferValidatorUpdated events raised by the Magiceden contract.
type MagicedenTransferValidatorUpdatedIterator struct {
	Event *MagicedenTransferValidatorUpdated // Event containing the contract specifics and raw log

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
func (it *MagicedenTransferValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenTransferValidatorUpdated)
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
		it.Event = new(MagicedenTransferValidatorUpdated)
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
func (it *MagicedenTransferValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenTransferValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenTransferValidatorUpdated represents a TransferValidatorUpdated event raised by the Magiceden contract.
type MagicedenTransferValidatorUpdated struct {
	OldValidator common.Address
	NewValidator common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransferValidatorUpdated is a free log retrieval operation binding the contract event 0xcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac.
//
// Solidity: event TransferValidatorUpdated(address oldValidator, address newValidator)
func (_Magiceden *MagicedenFilterer) FilterTransferValidatorUpdated(opts *bind.FilterOpts) (*MagicedenTransferValidatorUpdatedIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "TransferValidatorUpdated")
	if err != nil {
		return nil, err
	}
	return &MagicedenTransferValidatorUpdatedIterator{contract: _Magiceden.contract, event: "TransferValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchTransferValidatorUpdated is a free log subscription operation binding the contract event 0xcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac.
//
// Solidity: event TransferValidatorUpdated(address oldValidator, address newValidator)
func (_Magiceden *MagicedenFilterer) WatchTransferValidatorUpdated(opts *bind.WatchOpts, sink chan<- *MagicedenTransferValidatorUpdated) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "TransferValidatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenTransferValidatorUpdated)
				if err := _Magiceden.contract.UnpackLog(event, "TransferValidatorUpdated", log); err != nil {
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

// ParseTransferValidatorUpdated is a log parse operation binding the contract event 0xcc5dc080ff977b3c3a211fa63ab74f90f658f5ba9d3236e92c8f59570f442aac.
//
// Solidity: event TransferValidatorUpdated(address oldValidator, address newValidator)
func (_Magiceden *MagicedenFilterer) ParseTransferValidatorUpdated(log types.Log) (*MagicedenTransferValidatorUpdated, error) {
	event := new(MagicedenTransferValidatorUpdated)
	if err := _Magiceden.contract.UnpackLog(event, "TransferValidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenUpdateStageIterator is returned from FilterUpdateStage and is used to iterate over the raw logs and unpacked data for UpdateStage events raised by the Magiceden contract.
type MagicedenUpdateStageIterator struct {
	Event *MagicedenUpdateStage // Event containing the contract specifics and raw log

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
func (it *MagicedenUpdateStageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenUpdateStage)
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
		it.Event = new(MagicedenUpdateStage)
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
func (it *MagicedenUpdateStageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenUpdateStageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenUpdateStage represents a UpdateStage event raised by the Magiceden contract.
type MagicedenUpdateStage struct {
	Stage                *big.Int
	Price                *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds uint64
	EndTimeUnixSeconds   uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateStage is a free log retrieval operation binding the contract event 0xb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Magiceden *MagicedenFilterer) FilterUpdateStage(opts *bind.FilterOpts) (*MagicedenUpdateStageIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "UpdateStage")
	if err != nil {
		return nil, err
	}
	return &MagicedenUpdateStageIterator{contract: _Magiceden.contract, event: "UpdateStage", logs: logs, sub: sub}, nil
}

// WatchUpdateStage is a free log subscription operation binding the contract event 0xb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Magiceden *MagicedenFilterer) WatchUpdateStage(opts *bind.WatchOpts, sink chan<- *MagicedenUpdateStage) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "UpdateStage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenUpdateStage)
				if err := _Magiceden.contract.UnpackLog(event, "UpdateStage", log); err != nil {
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

// ParseUpdateStage is a log parse operation binding the contract event 0xb3268648542a1bb1b2dd12e3b14aeb5a3ab22c592de96bdd3e842154a5b394fa.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint64 startTimeUnixSeconds, uint64 endTimeUnixSeconds)
func (_Magiceden *MagicedenFilterer) ParseUpdateStage(log types.Log) (*MagicedenUpdateStage, error) {
	event := new(MagicedenUpdateStage)
	if err := _Magiceden.contract.UnpackLog(event, "UpdateStage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Magiceden contract.
type MagicedenWithdrawIterator struct {
	Event *MagicedenWithdraw // Event containing the contract specifics and raw log

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
func (it *MagicedenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenWithdraw)
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
		it.Event = new(MagicedenWithdraw)
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
func (it *MagicedenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenWithdraw represents a Withdraw event raised by the Magiceden contract.
type MagicedenWithdraw struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Magiceden *MagicedenFilterer) FilterWithdraw(opts *bind.FilterOpts) (*MagicedenWithdrawIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &MagicedenWithdrawIterator{contract: _Magiceden.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Magiceden *MagicedenFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MagicedenWithdraw) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenWithdraw)
				if err := _Magiceden.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: event Withdraw(uint256 value)
func (_Magiceden *MagicedenFilterer) ParseWithdraw(log types.Log) (*MagicedenWithdraw, error) {
	event := new(MagicedenWithdraw)
	if err := _Magiceden.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the Magiceden contract.
type MagicedenWithdrawERC20Iterator struct {
	Event *MagicedenWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *MagicedenWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenWithdrawERC20)
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
		it.Event = new(MagicedenWithdrawERC20)
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
func (it *MagicedenWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenWithdrawERC20 represents a WithdrawERC20 event raised by the Magiceden contract.
type MagicedenWithdrawERC20 struct {
	MintCurrency common.Address
	Value        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xbe7426aee8a34d0263892b55ce65ce81d8f4c806eb4719e59015ea49feb92d22.
//
// Solidity: event WithdrawERC20(address mintCurrency, uint256 value)
func (_Magiceden *MagicedenFilterer) FilterWithdrawERC20(opts *bind.FilterOpts) (*MagicedenWithdrawERC20Iterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "WithdrawERC20")
	if err != nil {
		return nil, err
	}
	return &MagicedenWithdrawERC20Iterator{contract: _Magiceden.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xbe7426aee8a34d0263892b55ce65ce81d8f4c806eb4719e59015ea49feb92d22.
//
// Solidity: event WithdrawERC20(address mintCurrency, uint256 value)
func (_Magiceden *MagicedenFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *MagicedenWithdrawERC20) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "WithdrawERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenWithdrawERC20)
				if err := _Magiceden.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xbe7426aee8a34d0263892b55ce65ce81d8f4c806eb4719e59015ea49feb92d22.
//
// Solidity: event WithdrawERC20(address mintCurrency, uint256 value)
func (_Magiceden *MagicedenFilterer) ParseWithdrawERC20(log types.Log) (*MagicedenWithdrawERC20, error) {
	event := new(MagicedenWithdrawERC20)
	if err := _Magiceden.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
