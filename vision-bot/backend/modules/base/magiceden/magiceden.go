// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base

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

// IERC721AUpgradeableTokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type IERC721AUpgradeableTokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
	Burned         bool
	ExtraData      *big.Int
}

// MintStageInfo is an auto generated low-level Go binding around an user-defined struct.
type MintStageInfo struct {
	Price                *big.Int
	MintFee              *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds *big.Int
	EndTimeUnixSeconds   *big.Int
}

// MagicedenMetaData contains all meta data concerning the Magiceden contract.
var MagicedenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotIncreaseMaxMintableSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignerNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorTokenBase__InvalidTransferValidatorContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalWalletLimitOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitialOwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStageTimeGap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQueryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStageArgsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewSupplyLessThanTotalSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSupplyLeft\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCompatibleWithSpotMints\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTransferable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoyaltyReceiverIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialMintExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequentialUpToTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ShouldNotMintToBurnAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SpotMintTokenIdTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StageSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletGlobalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletStageLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongMintCurrency\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"AuthorizedMinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"AuthorizedMinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"autoApproved\",\"type\":\"bool\"}],\"name\":\"AutomaticApprovalOfTransferValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"DefaultRoyaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeStage\",\"type\":\"uint256\"}],\"name\":\"SetActiveStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"SetBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"SetContractURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"SetCosigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"SetDefaultRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"SetGlobalWalletLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"SetMaxMintableSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"}],\"name\":\"SetMintCurrency\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"SetMintable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestampExpirySeconds\",\"type\":\"uint256\"}],\"name\":\"SetTimestampExpirySeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"name\":\"SetTokenURISuffix\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"TokenRoyaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"TransferValidatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint80\",\"name\":\"mintFee\",\"type\":\"uint80\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint256\"}],\"name\":\"UpdateStage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_TRANSFER_VALIDATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"__ERC721ACQueryableInitializable_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"addAuthorizedMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"cosignNonce\",\"type\":\"uint256\"}],\"name\":\"assertValidCosign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"limit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"authorizedMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoApproveTransfersFromValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractNameAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"explicitOwnershipOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721AUpgradeable.TokenOwnership\",\"name\":\"ownership\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"explicitOwnershipsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"burned\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"extraData\",\"type\":\"uint24\"}],\"internalType\":\"structIERC721AUpgradeable.TokenOwnership[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getActiveStageFromTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"waiveMintFee\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cosignNonce\",\"type\":\"uint256\"}],\"name\":\"getCosignDigest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getCosignNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalWalletLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxMintableSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberStages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStageInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"mintFee\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structMintStageInfo\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimestampExpirySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransferValidationFunction\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"functionSignature\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"isViewFunction\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransferValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"isAuthorizedMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"limit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"qty\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ownerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"removeAuthorizedMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"}],\"name\":\"setAutomaticApprovalOfTransfersFromValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"setCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"}],\"name\":\"setGlobalWalletLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxMintableSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"}],\"name\":\"setMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"mintFee\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structMintStageInfo[]\",\"name\":\"newStages\",\"type\":\"tuple[]\"}],\"name\":\"setStages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestampExpirySeconds\",\"type\":\"uint256\"}],\"name\":\"setTimestampExpirySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"suffix\",\"type\":\"string\"}],\"name\":\"setTokenURISuffix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transferValidator_\",\"type\":\"address\"}],\"name\":\"setTransferValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxMintableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"globalWalletLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"price\",\"type\":\"uint80\"},{\"internalType\":\"uint80\",\"name\":\"mintFee\",\"type\":\"uint80\"},{\"internalType\":\"uint32\",\"name\":\"walletLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"maxStageSupply\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"startTimeUnixSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimeUnixSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structMintStageInfo[]\",\"name\":\"initialStages\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"royaltyFeeNumerator\",\"type\":\"uint96\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stop\",\"type\":\"uint256\"}],\"name\":\"tokensOfOwnerIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"totalMintedByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AssertValidCosign is a free data retrieval call binding the contract method 0x8febc54b.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint256 timestamp, bytes signature, uint256 cosignNonce) view returns(bool)
func (_Magiceden *MagicedenCaller) AssertValidCosign(opts *bind.CallOpts, minter common.Address, qty uint32, timestamp *big.Int, signature []byte, cosignNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "assertValidCosign", minter, qty, timestamp, signature, cosignNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssertValidCosign is a free data retrieval call binding the contract method 0x8febc54b.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint256 timestamp, bytes signature, uint256 cosignNonce) view returns(bool)
func (_Magiceden *MagicedenSession) AssertValidCosign(minter common.Address, qty uint32, timestamp *big.Int, signature []byte, cosignNonce *big.Int) (bool, error) {
	return _Magiceden.Contract.AssertValidCosign(&_Magiceden.CallOpts, minter, qty, timestamp, signature, cosignNonce)
}

// AssertValidCosign is a free data retrieval call binding the contract method 0x8febc54b.
//
// Solidity: function assertValidCosign(address minter, uint32 qty, uint256 timestamp, bytes signature, uint256 cosignNonce) view returns(bool)
func (_Magiceden *MagicedenCallerSession) AssertValidCosign(minter common.Address, qty uint32, timestamp *big.Int, signature []byte, cosignNonce *big.Int) (bool, error) {
	return _Magiceden.Contract.AssertValidCosign(&_Magiceden.CallOpts, minter, qty, timestamp, signature, cosignNonce)
}

// AutoApproveTransfersFromValidator is a free data retrieval call binding the contract method 0x6221d13c.
//
// Solidity: function autoApproveTransfersFromValidator() view returns(bool)
func (_Magiceden *MagicedenCaller) AutoApproveTransfersFromValidator(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "autoApproveTransfersFromValidator")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutoApproveTransfersFromValidator is a free data retrieval call binding the contract method 0x6221d13c.
//
// Solidity: function autoApproveTransfersFromValidator() view returns(bool)
func (_Magiceden *MagicedenSession) AutoApproveTransfersFromValidator() (bool, error) {
	return _Magiceden.Contract.AutoApproveTransfersFromValidator(&_Magiceden.CallOpts)
}

// AutoApproveTransfersFromValidator is a free data retrieval call binding the contract method 0x6221d13c.
//
// Solidity: function autoApproveTransfersFromValidator() view returns(bool)
func (_Magiceden *MagicedenCallerSession) AutoApproveTransfersFromValidator() (bool, error) {
	return _Magiceden.Contract.AutoApproveTransfersFromValidator(&_Magiceden.CallOpts)
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

// ContractNameAndVersion is a free data retrieval call binding the contract method 0xc3db27c1.
//
// Solidity: function contractNameAndVersion() pure returns(string, string)
func (_Magiceden *MagicedenCaller) ContractNameAndVersion(opts *bind.CallOpts) (string, string, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "contractNameAndVersion")

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// ContractNameAndVersion is a free data retrieval call binding the contract method 0xc3db27c1.
//
// Solidity: function contractNameAndVersion() pure returns(string, string)
func (_Magiceden *MagicedenSession) ContractNameAndVersion() (string, string, error) {
	return _Magiceden.Contract.ContractNameAndVersion(&_Magiceden.CallOpts)
}

// ContractNameAndVersion is a free data retrieval call binding the contract method 0xc3db27c1.
//
// Solidity: function contractNameAndVersion() pure returns(string, string)
func (_Magiceden *MagicedenCallerSession) ContractNameAndVersion() (string, string, error) {
	return _Magiceden.Contract.ContractNameAndVersion(&_Magiceden.CallOpts)
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
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Magiceden *MagicedenCaller) ExplicitOwnershipOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "explicitOwnershipOf", tokenId)

	if err != nil {
		return *new(IERC721AUpgradeableTokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721AUpgradeableTokenOwnership)).(*IERC721AUpgradeableTokenOwnership)

	return out0, err

}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Magiceden *MagicedenSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipOf(&_Magiceden.CallOpts, tokenId)
}

// ExplicitOwnershipOf is a free data retrieval call binding the contract method 0xc23dc68f.
//
// Solidity: function explicitOwnershipOf(uint256 tokenId) view returns((address,uint64,bool,uint24) ownership)
func (_Magiceden *MagicedenCallerSession) ExplicitOwnershipOf(tokenId *big.Int) (IERC721AUpgradeableTokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipOf(&_Magiceden.CallOpts, tokenId)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Magiceden *MagicedenCaller) ExplicitOwnershipsOf(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "explicitOwnershipsOf", tokenIds)

	if err != nil {
		return *new([]IERC721AUpgradeableTokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC721AUpgradeableTokenOwnership)).(*[]IERC721AUpgradeableTokenOwnership)

	return out0, err

}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Magiceden *MagicedenSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipsOf(&_Magiceden.CallOpts, tokenIds)
}

// ExplicitOwnershipsOf is a free data retrieval call binding the contract method 0x5bbb2177.
//
// Solidity: function explicitOwnershipsOf(uint256[] tokenIds) view returns((address,uint64,bool,uint24)[])
func (_Magiceden *MagicedenCallerSession) ExplicitOwnershipsOf(tokenIds []*big.Int) ([]IERC721AUpgradeableTokenOwnership, error) {
	return _Magiceden.Contract.ExplicitOwnershipsOf(&_Magiceden.CallOpts, tokenIds)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x72bbedb8.
//
// Solidity: function getActiveStageFromTimestamp(uint256 timestamp) view returns(uint256)
func (_Magiceden *MagicedenCaller) GetActiveStageFromTimestamp(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getActiveStageFromTimestamp", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x72bbedb8.
//
// Solidity: function getActiveStageFromTimestamp(uint256 timestamp) view returns(uint256)
func (_Magiceden *MagicedenSession) GetActiveStageFromTimestamp(timestamp *big.Int) (*big.Int, error) {
	return _Magiceden.Contract.GetActiveStageFromTimestamp(&_Magiceden.CallOpts, timestamp)
}

// GetActiveStageFromTimestamp is a free data retrieval call binding the contract method 0x72bbedb8.
//
// Solidity: function getActiveStageFromTimestamp(uint256 timestamp) view returns(uint256)
func (_Magiceden *MagicedenCallerSession) GetActiveStageFromTimestamp(timestamp *big.Int) (*big.Int, error) {
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

// GetCosignDigest is a free data retrieval call binding the contract method 0xff1b4ba9.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, bool waiveMintFee, uint256 timestamp, uint256 cosignNonce) view returns(bytes32)
func (_Magiceden *MagicedenCaller) GetCosignDigest(opts *bind.CallOpts, minter common.Address, qty uint32, waiveMintFee bool, timestamp *big.Int, cosignNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getCosignDigest", minter, qty, waiveMintFee, timestamp, cosignNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCosignDigest is a free data retrieval call binding the contract method 0xff1b4ba9.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, bool waiveMintFee, uint256 timestamp, uint256 cosignNonce) view returns(bytes32)
func (_Magiceden *MagicedenSession) GetCosignDigest(minter common.Address, qty uint32, waiveMintFee bool, timestamp *big.Int, cosignNonce *big.Int) ([32]byte, error) {
	return _Magiceden.Contract.GetCosignDigest(&_Magiceden.CallOpts, minter, qty, waiveMintFee, timestamp, cosignNonce)
}

// GetCosignDigest is a free data retrieval call binding the contract method 0xff1b4ba9.
//
// Solidity: function getCosignDigest(address minter, uint32 qty, bool waiveMintFee, uint256 timestamp, uint256 cosignNonce) view returns(bytes32)
func (_Magiceden *MagicedenCallerSession) GetCosignDigest(minter common.Address, qty uint32, waiveMintFee bool, timestamp *big.Int, cosignNonce *big.Int) ([32]byte, error) {
	return _Magiceden.Contract.GetCosignDigest(&_Magiceden.CallOpts, minter, qty, waiveMintFee, timestamp, cosignNonce)
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

// GetCosigner is a free data retrieval call binding the contract method 0x33bbbf06.
//
// Solidity: function getCosigner() view returns(address)
func (_Magiceden *MagicedenCaller) GetCosigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getCosigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCosigner is a free data retrieval call binding the contract method 0x33bbbf06.
//
// Solidity: function getCosigner() view returns(address)
func (_Magiceden *MagicedenSession) GetCosigner() (common.Address, error) {
	return _Magiceden.Contract.GetCosigner(&_Magiceden.CallOpts)
}

// GetCosigner is a free data retrieval call binding the contract method 0x33bbbf06.
//
// Solidity: function getCosigner() view returns(address)
func (_Magiceden *MagicedenCallerSession) GetCosigner() (common.Address, error) {
	return _Magiceden.Contract.GetCosigner(&_Magiceden.CallOpts)
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

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint80,uint32,bytes32,uint24,uint256,uint256), uint32, uint256)
func (_Magiceden *MagicedenCaller) GetStageInfo(opts *bind.CallOpts, index *big.Int) (MintStageInfo, uint32, *big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getStageInfo", index)

	if err != nil {
		return *new(MintStageInfo), *new(uint32), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(MintStageInfo)).(*MintStageInfo)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint80,uint32,bytes32,uint24,uint256,uint256), uint32, uint256)
func (_Magiceden *MagicedenSession) GetStageInfo(index *big.Int) (MintStageInfo, uint32, *big.Int, error) {
	return _Magiceden.Contract.GetStageInfo(&_Magiceden.CallOpts, index)
}

// GetStageInfo is a free data retrieval call binding the contract method 0xa3759f60.
//
// Solidity: function getStageInfo(uint256 index) view returns((uint80,uint80,uint32,bytes32,uint24,uint256,uint256), uint32, uint256)
func (_Magiceden *MagicedenCallerSession) GetStageInfo(index *big.Int) (MintStageInfo, uint32, *big.Int, error) {
	return _Magiceden.Contract.GetStageInfo(&_Magiceden.CallOpts, index)
}

// GetTimestampExpirySeconds is a free data retrieval call binding the contract method 0x4ae0402f.
//
// Solidity: function getTimestampExpirySeconds() view returns(uint256)
func (_Magiceden *MagicedenCaller) GetTimestampExpirySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getTimestampExpirySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampExpirySeconds is a free data retrieval call binding the contract method 0x4ae0402f.
//
// Solidity: function getTimestampExpirySeconds() view returns(uint256)
func (_Magiceden *MagicedenSession) GetTimestampExpirySeconds() (*big.Int, error) {
	return _Magiceden.Contract.GetTimestampExpirySeconds(&_Magiceden.CallOpts)
}

// GetTimestampExpirySeconds is a free data retrieval call binding the contract method 0x4ae0402f.
//
// Solidity: function getTimestampExpirySeconds() view returns(uint256)
func (_Magiceden *MagicedenCallerSession) GetTimestampExpirySeconds() (*big.Int, error) {
	return _Magiceden.Contract.GetTimestampExpirySeconds(&_Magiceden.CallOpts)
}

// GetTransferValidationFunction is a free data retrieval call binding the contract method 0x0d705df6.
//
// Solidity: function getTransferValidationFunction() pure returns(bytes4 functionSignature, bool isViewFunction)
func (_Magiceden *MagicedenCaller) GetTransferValidationFunction(opts *bind.CallOpts) (struct {
	FunctionSignature [4]byte
	IsViewFunction    bool
}, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "getTransferValidationFunction")

	outstruct := new(struct {
		FunctionSignature [4]byte
		IsViewFunction    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FunctionSignature = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.IsViewFunction = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetTransferValidationFunction is a free data retrieval call binding the contract method 0x0d705df6.
//
// Solidity: function getTransferValidationFunction() pure returns(bytes4 functionSignature, bool isViewFunction)
func (_Magiceden *MagicedenSession) GetTransferValidationFunction() (struct {
	FunctionSignature [4]byte
	IsViewFunction    bool
}, error) {
	return _Magiceden.Contract.GetTransferValidationFunction(&_Magiceden.CallOpts)
}

// GetTransferValidationFunction is a free data retrieval call binding the contract method 0x0d705df6.
//
// Solidity: function getTransferValidationFunction() pure returns(bytes4 functionSignature, bool isViewFunction)
func (_Magiceden *MagicedenCallerSession) GetTransferValidationFunction() (struct {
	FunctionSignature [4]byte
	IsViewFunction    bool
}, error) {
	return _Magiceden.Contract.GetTransferValidationFunction(&_Magiceden.CallOpts)
}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address validator)
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
// Solidity: function getTransferValidator() view returns(address validator)
func (_Magiceden *MagicedenSession) GetTransferValidator() (common.Address, error) {
	return _Magiceden.Contract.GetTransferValidator(&_Magiceden.CallOpts)
}

// GetTransferValidator is a free data retrieval call binding the contract method 0x098144d4.
//
// Solidity: function getTransferValidator() view returns(address validator)
func (_Magiceden *MagicedenCallerSession) GetTransferValidator() (common.Address, error) {
	return _Magiceden.Contract.GetTransferValidator(&_Magiceden.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool isApproved)
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
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool isApproved)
func (_Magiceden *MagicedenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Magiceden.Contract.IsApprovedForAll(&_Magiceden.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool isApproved)
func (_Magiceden *MagicedenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Magiceden.Contract.IsApprovedForAll(&_Magiceden.CallOpts, owner, operator)
}

// IsAuthorizedMinter is a free data retrieval call binding the contract method 0x842392c2.
//
// Solidity: function isAuthorizedMinter(address minter) view returns(bool)
func (_Magiceden *MagicedenCaller) IsAuthorizedMinter(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "isAuthorizedMinter", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedMinter is a free data retrieval call binding the contract method 0x842392c2.
//
// Solidity: function isAuthorizedMinter(address minter) view returns(bool)
func (_Magiceden *MagicedenSession) IsAuthorizedMinter(minter common.Address) (bool, error) {
	return _Magiceden.Contract.IsAuthorizedMinter(&_Magiceden.CallOpts, minter)
}

// IsAuthorizedMinter is a free data retrieval call binding the contract method 0x842392c2.
//
// Solidity: function isAuthorizedMinter(address minter) view returns(bool)
func (_Magiceden *MagicedenCallerSession) IsAuthorizedMinter(minter common.Address) (bool, error) {
	return _Magiceden.Contract.IsAuthorizedMinter(&_Magiceden.CallOpts, minter)
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
// Solidity: function owner() view returns(address result)
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
// Solidity: function owner() view returns(address result)
func (_Magiceden *MagicedenSession) Owner() (common.Address, error) {
	return _Magiceden.Contract.Owner(&_Magiceden.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
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

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Magiceden *MagicedenCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Magiceden *MagicedenSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Magiceden.Contract.OwnershipHandoverExpiresAt(&_Magiceden.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Magiceden *MagicedenCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Magiceden.Contract.OwnershipHandoverExpiresAt(&_Magiceden.CallOpts, pendingOwner)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Magiceden *MagicedenCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Magiceden.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Magiceden *MagicedenSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Magiceden.Contract.RoyaltyInfo(&_Magiceden.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Magiceden *MagicedenCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
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
// Solidity: function totalSupply() view returns(uint256 result)
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
// Solidity: function totalSupply() view returns(uint256 result)
func (_Magiceden *MagicedenSession) TotalSupply() (*big.Int, error) {
	return _Magiceden.Contract.TotalSupply(&_Magiceden.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_Magiceden *MagicedenCallerSession) TotalSupply() (*big.Int, error) {
	return _Magiceden.Contract.TotalSupply(&_Magiceden.CallOpts)
}

// ERC721ACQueryableInitializableInit is a paid mutator transaction binding the contract method 0xf73134d0.
//
// Solidity: function __ERC721ACQueryableInitializable_init(string name_, string symbol_) returns()
func (_Magiceden *MagicedenTransactor) ERC721ACQueryableInitializableInit(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "__ERC721ACQueryableInitializable_init", name_, symbol_)
}

// ERC721ACQueryableInitializableInit is a paid mutator transaction binding the contract method 0xf73134d0.
//
// Solidity: function __ERC721ACQueryableInitializable_init(string name_, string symbol_) returns()
func (_Magiceden *MagicedenSession) ERC721ACQueryableInitializableInit(name_ string, symbol_ string) (*types.Transaction, error) {
	return _Magiceden.Contract.ERC721ACQueryableInitializableInit(&_Magiceden.TransactOpts, name_, symbol_)
}

// ERC721ACQueryableInitializableInit is a paid mutator transaction binding the contract method 0xf73134d0.
//
// Solidity: function __ERC721ACQueryableInitializable_init(string name_, string symbol_) returns()
func (_Magiceden *MagicedenTransactorSession) ERC721ACQueryableInitializableInit(name_ string, symbol_ string) (*types.Transaction, error) {
	return _Magiceden.Contract.ERC721ACQueryableInitializableInit(&_Magiceden.TransactOpts, name_, symbol_)
}

// AddAuthorizedMinter is a paid mutator transaction binding the contract method 0x5f710f5c.
//
// Solidity: function addAuthorizedMinter(address minter) returns()
func (_Magiceden *MagicedenTransactor) AddAuthorizedMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "addAuthorizedMinter", minter)
}

// AddAuthorizedMinter is a paid mutator transaction binding the contract method 0x5f710f5c.
//
// Solidity: function addAuthorizedMinter(address minter) returns()
func (_Magiceden *MagicedenSession) AddAuthorizedMinter(minter common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.AddAuthorizedMinter(&_Magiceden.TransactOpts, minter)
}

// AddAuthorizedMinter is a paid mutator transaction binding the contract method 0x5f710f5c.
//
// Solidity: function addAuthorizedMinter(address minter) returns()
func (_Magiceden *MagicedenTransactorSession) AddAuthorizedMinter(minter common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.AddAuthorizedMinter(&_Magiceden.TransactOpts, minter)
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

// AuthorizedMint is a paid mutator transaction binding the contract method 0xb38c9f4c.
//
// Solidity: function authorizedMint(uint32 qty, address to, uint32 limit, bytes32[] proof, uint256 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactor) AuthorizedMint(opts *bind.TransactOpts, qty uint32, to common.Address, limit uint32, proof [][32]byte, timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "authorizedMint", qty, to, limit, proof, timestamp, signature)
}

// AuthorizedMint is a paid mutator transaction binding the contract method 0xb38c9f4c.
//
// Solidity: function authorizedMint(uint32 qty, address to, uint32 limit, bytes32[] proof, uint256 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenSession) AuthorizedMint(qty uint32, to common.Address, limit uint32, proof [][32]byte, timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.AuthorizedMint(&_Magiceden.TransactOpts, qty, to, limit, proof, timestamp, signature)
}

// AuthorizedMint is a paid mutator transaction binding the contract method 0xb38c9f4c.
//
// Solidity: function authorizedMint(uint32 qty, address to, uint32 limit, bytes32[] proof, uint256 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactorSession) AuthorizedMint(qty uint32, to common.Address, limit uint32, proof [][32]byte, timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.AuthorizedMint(&_Magiceden.TransactOpts, qty, to, limit, proof, timestamp, signature)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Magiceden *MagicedenTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Magiceden *MagicedenSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Magiceden.Contract.CancelOwnershipHandover(&_Magiceden.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Magiceden *MagicedenTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Magiceden.Contract.CancelOwnershipHandover(&_Magiceden.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Magiceden *MagicedenTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Magiceden *MagicedenSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.CompleteOwnershipHandover(&_Magiceden.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Magiceden *MagicedenTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.CompleteOwnershipHandover(&_Magiceden.TransactOpts, pendingOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address initialOwner) returns()
func (_Magiceden *MagicedenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, initialOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "initialize", name, symbol, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address initialOwner) returns()
func (_Magiceden *MagicedenSession) Initialize(name string, symbol string, initialOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.Initialize(&_Magiceden.TransactOpts, name, symbol, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address initialOwner) returns()
func (_Magiceden *MagicedenTransactorSession) Initialize(name string, symbol string, initialOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.Initialize(&_Magiceden.TransactOpts, name, symbol, initialOwner)
}

// Mint is a paid mutator transaction binding the contract method 0xb971b4c4.
//
// Solidity: function mint(uint32 qty, uint32 limit, bytes32[] proof, uint256 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactor) Mint(opts *bind.TransactOpts, qty uint32, limit uint32, proof [][32]byte, timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "mint", qty, limit, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xb971b4c4.
//
// Solidity: function mint(uint32 qty, uint32 limit, bytes32[] proof, uint256 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenSession) Mint(qty uint32, limit uint32, proof [][32]byte, timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.Mint(&_Magiceden.TransactOpts, qty, limit, proof, timestamp, signature)
}

// Mint is a paid mutator transaction binding the contract method 0xb971b4c4.
//
// Solidity: function mint(uint32 qty, uint32 limit, bytes32[] proof, uint256 timestamp, bytes signature) payable returns()
func (_Magiceden *MagicedenTransactorSession) Mint(qty uint32, limit uint32, proof [][32]byte, timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Magiceden.Contract.Mint(&_Magiceden.TransactOpts, qty, limit, proof, timestamp, signature)
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

// RemoveAuthorizedMinter is a paid mutator transaction binding the contract method 0x475ae039.
//
// Solidity: function removeAuthorizedMinter(address minter) returns()
func (_Magiceden *MagicedenTransactor) RemoveAuthorizedMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "removeAuthorizedMinter", minter)
}

// RemoveAuthorizedMinter is a paid mutator transaction binding the contract method 0x475ae039.
//
// Solidity: function removeAuthorizedMinter(address minter) returns()
func (_Magiceden *MagicedenSession) RemoveAuthorizedMinter(minter common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.RemoveAuthorizedMinter(&_Magiceden.TransactOpts, minter)
}

// RemoveAuthorizedMinter is a paid mutator transaction binding the contract method 0x475ae039.
//
// Solidity: function removeAuthorizedMinter(address minter) returns()
func (_Magiceden *MagicedenTransactorSession) RemoveAuthorizedMinter(minter common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.RemoveAuthorizedMinter(&_Magiceden.TransactOpts, minter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Magiceden *MagicedenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Magiceden *MagicedenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Magiceden.Contract.RenounceOwnership(&_Magiceden.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Magiceden *MagicedenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Magiceden.Contract.RenounceOwnership(&_Magiceden.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Magiceden *MagicedenTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Magiceden *MagicedenSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Magiceden.Contract.RequestOwnershipHandover(&_Magiceden.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Magiceden *MagicedenTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Magiceden.Contract.RequestOwnershipHandover(&_Magiceden.TransactOpts)
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

// SetAutomaticApprovalOfTransfersFromValidator is a paid mutator transaction binding the contract method 0x9e05d240.
//
// Solidity: function setAutomaticApprovalOfTransfersFromValidator(bool autoApprove) returns()
func (_Magiceden *MagicedenTransactor) SetAutomaticApprovalOfTransfersFromValidator(opts *bind.TransactOpts, autoApprove bool) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setAutomaticApprovalOfTransfersFromValidator", autoApprove)
}

// SetAutomaticApprovalOfTransfersFromValidator is a paid mutator transaction binding the contract method 0x9e05d240.
//
// Solidity: function setAutomaticApprovalOfTransfersFromValidator(bool autoApprove) returns()
func (_Magiceden *MagicedenSession) SetAutomaticApprovalOfTransfersFromValidator(autoApprove bool) (*types.Transaction, error) {
	return _Magiceden.Contract.SetAutomaticApprovalOfTransfersFromValidator(&_Magiceden.TransactOpts, autoApprove)
}

// SetAutomaticApprovalOfTransfersFromValidator is a paid mutator transaction binding the contract method 0x9e05d240.
//
// Solidity: function setAutomaticApprovalOfTransfersFromValidator(bool autoApprove) returns()
func (_Magiceden *MagicedenTransactorSession) SetAutomaticApprovalOfTransfersFromValidator(autoApprove bool) (*types.Transaction, error) {
	return _Magiceden.Contract.SetAutomaticApprovalOfTransfersFromValidator(&_Magiceden.TransactOpts, autoApprove)
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

// SetStages is a paid mutator transaction binding the contract method 0x359f7183.
//
// Solidity: function setStages((uint80,uint80,uint32,bytes32,uint24,uint256,uint256)[] newStages) returns()
func (_Magiceden *MagicedenTransactor) SetStages(opts *bind.TransactOpts, newStages []MintStageInfo) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setStages", newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0x359f7183.
//
// Solidity: function setStages((uint80,uint80,uint32,bytes32,uint24,uint256,uint256)[] newStages) returns()
func (_Magiceden *MagicedenSession) SetStages(newStages []MintStageInfo) (*types.Transaction, error) {
	return _Magiceden.Contract.SetStages(&_Magiceden.TransactOpts, newStages)
}

// SetStages is a paid mutator transaction binding the contract method 0x359f7183.
//
// Solidity: function setStages((uint80,uint80,uint32,bytes32,uint24,uint256,uint256)[] newStages) returns()
func (_Magiceden *MagicedenTransactorSession) SetStages(newStages []MintStageInfo) (*types.Transaction, error) {
	return _Magiceden.Contract.SetStages(&_Magiceden.TransactOpts, newStages)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0x16dc705c.
//
// Solidity: function setTimestampExpirySeconds(uint256 timestampExpirySeconds) returns()
func (_Magiceden *MagicedenTransactor) SetTimestampExpirySeconds(opts *bind.TransactOpts, timestampExpirySeconds *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setTimestampExpirySeconds", timestampExpirySeconds)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0x16dc705c.
//
// Solidity: function setTimestampExpirySeconds(uint256 timestampExpirySeconds) returns()
func (_Magiceden *MagicedenSession) SetTimestampExpirySeconds(timestampExpirySeconds *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTimestampExpirySeconds(&_Magiceden.TransactOpts, timestampExpirySeconds)
}

// SetTimestampExpirySeconds is a paid mutator transaction binding the contract method 0x16dc705c.
//
// Solidity: function setTimestampExpirySeconds(uint256 timestampExpirySeconds) returns()
func (_Magiceden *MagicedenTransactorSession) SetTimestampExpirySeconds(timestampExpirySeconds *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.SetTimestampExpirySeconds(&_Magiceden.TransactOpts, timestampExpirySeconds)
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

// Setup is a paid mutator transaction binding the contract method 0xac0f3d86.
//
// Solidity: function setup(uint256 maxMintableSupply, uint256 globalWalletLimit, address mintCurrency, address fundReceiver, (uint80,uint80,uint32,bytes32,uint24,uint256,uint256)[] initialStages, address royaltyReceiver, uint96 royaltyFeeNumerator) returns()
func (_Magiceden *MagicedenTransactor) Setup(opts *bind.TransactOpts, maxMintableSupply *big.Int, globalWalletLimit *big.Int, mintCurrency common.Address, fundReceiver common.Address, initialStages []MintStageInfo, royaltyReceiver common.Address, royaltyFeeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "setup", maxMintableSupply, globalWalletLimit, mintCurrency, fundReceiver, initialStages, royaltyReceiver, royaltyFeeNumerator)
}

// Setup is a paid mutator transaction binding the contract method 0xac0f3d86.
//
// Solidity: function setup(uint256 maxMintableSupply, uint256 globalWalletLimit, address mintCurrency, address fundReceiver, (uint80,uint80,uint32,bytes32,uint24,uint256,uint256)[] initialStages, address royaltyReceiver, uint96 royaltyFeeNumerator) returns()
func (_Magiceden *MagicedenSession) Setup(maxMintableSupply *big.Int, globalWalletLimit *big.Int, mintCurrency common.Address, fundReceiver common.Address, initialStages []MintStageInfo, royaltyReceiver common.Address, royaltyFeeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.Setup(&_Magiceden.TransactOpts, maxMintableSupply, globalWalletLimit, mintCurrency, fundReceiver, initialStages, royaltyReceiver, royaltyFeeNumerator)
}

// Setup is a paid mutator transaction binding the contract method 0xac0f3d86.
//
// Solidity: function setup(uint256 maxMintableSupply, uint256 globalWalletLimit, address mintCurrency, address fundReceiver, (uint80,uint80,uint32,bytes32,uint24,uint256,uint256)[] initialStages, address royaltyReceiver, uint96 royaltyFeeNumerator) returns()
func (_Magiceden *MagicedenTransactorSession) Setup(maxMintableSupply *big.Int, globalWalletLimit *big.Int, mintCurrency common.Address, fundReceiver common.Address, initialStages []MintStageInfo, royaltyReceiver common.Address, royaltyFeeNumerator *big.Int) (*types.Transaction, error) {
	return _Magiceden.Contract.Setup(&_Magiceden.TransactOpts, maxMintableSupply, globalWalletLimit, mintCurrency, fundReceiver, initialStages, royaltyReceiver, royaltyFeeNumerator)
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
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Magiceden *MagicedenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Magiceden *MagicedenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.TransferOwnership(&_Magiceden.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Magiceden *MagicedenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Magiceden.Contract.TransferOwnership(&_Magiceden.TransactOpts, newOwner)
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

// MagicedenAuthorizedMinterAddedIterator is returned from FilterAuthorizedMinterAdded and is used to iterate over the raw logs and unpacked data for AuthorizedMinterAdded events raised by the Magiceden contract.
type MagicedenAuthorizedMinterAddedIterator struct {
	Event *MagicedenAuthorizedMinterAdded // Event containing the contract specifics and raw log

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
func (it *MagicedenAuthorizedMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenAuthorizedMinterAdded)
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
		it.Event = new(MagicedenAuthorizedMinterAdded)
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
func (it *MagicedenAuthorizedMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenAuthorizedMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenAuthorizedMinterAdded represents a AuthorizedMinterAdded event raised by the Magiceden contract.
type MagicedenAuthorizedMinterAdded struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedMinterAdded is a free log retrieval operation binding the contract event 0xe6be4d6cc04eb0219337b22db08c688969a9ec8e34d9a0a2ba38a114e050f1ae.
//
// Solidity: event AuthorizedMinterAdded(address indexed minter)
func (_Magiceden *MagicedenFilterer) FilterAuthorizedMinterAdded(opts *bind.FilterOpts, minter []common.Address) (*MagicedenAuthorizedMinterAddedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "AuthorizedMinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenAuthorizedMinterAddedIterator{contract: _Magiceden.contract, event: "AuthorizedMinterAdded", logs: logs, sub: sub}, nil
}

// WatchAuthorizedMinterAdded is a free log subscription operation binding the contract event 0xe6be4d6cc04eb0219337b22db08c688969a9ec8e34d9a0a2ba38a114e050f1ae.
//
// Solidity: event AuthorizedMinterAdded(address indexed minter)
func (_Magiceden *MagicedenFilterer) WatchAuthorizedMinterAdded(opts *bind.WatchOpts, sink chan<- *MagicedenAuthorizedMinterAdded, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "AuthorizedMinterAdded", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenAuthorizedMinterAdded)
				if err := _Magiceden.contract.UnpackLog(event, "AuthorizedMinterAdded", log); err != nil {
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

// ParseAuthorizedMinterAdded is a log parse operation binding the contract event 0xe6be4d6cc04eb0219337b22db08c688969a9ec8e34d9a0a2ba38a114e050f1ae.
//
// Solidity: event AuthorizedMinterAdded(address indexed minter)
func (_Magiceden *MagicedenFilterer) ParseAuthorizedMinterAdded(log types.Log) (*MagicedenAuthorizedMinterAdded, error) {
	event := new(MagicedenAuthorizedMinterAdded)
	if err := _Magiceden.contract.UnpackLog(event, "AuthorizedMinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenAuthorizedMinterRemovedIterator is returned from FilterAuthorizedMinterRemoved and is used to iterate over the raw logs and unpacked data for AuthorizedMinterRemoved events raised by the Magiceden contract.
type MagicedenAuthorizedMinterRemovedIterator struct {
	Event *MagicedenAuthorizedMinterRemoved // Event containing the contract specifics and raw log

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
func (it *MagicedenAuthorizedMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenAuthorizedMinterRemoved)
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
		it.Event = new(MagicedenAuthorizedMinterRemoved)
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
func (it *MagicedenAuthorizedMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenAuthorizedMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenAuthorizedMinterRemoved represents a AuthorizedMinterRemoved event raised by the Magiceden contract.
type MagicedenAuthorizedMinterRemoved struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedMinterRemoved is a free log retrieval operation binding the contract event 0xc6711413797b8a562634e98c95d50e7619d39702ed5b82ce335dc93546c3a88c.
//
// Solidity: event AuthorizedMinterRemoved(address indexed minter)
func (_Magiceden *MagicedenFilterer) FilterAuthorizedMinterRemoved(opts *bind.FilterOpts, minter []common.Address) (*MagicedenAuthorizedMinterRemovedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "AuthorizedMinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenAuthorizedMinterRemovedIterator{contract: _Magiceden.contract, event: "AuthorizedMinterRemoved", logs: logs, sub: sub}, nil
}

// WatchAuthorizedMinterRemoved is a free log subscription operation binding the contract event 0xc6711413797b8a562634e98c95d50e7619d39702ed5b82ce335dc93546c3a88c.
//
// Solidity: event AuthorizedMinterRemoved(address indexed minter)
func (_Magiceden *MagicedenFilterer) WatchAuthorizedMinterRemoved(opts *bind.WatchOpts, sink chan<- *MagicedenAuthorizedMinterRemoved, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "AuthorizedMinterRemoved", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenAuthorizedMinterRemoved)
				if err := _Magiceden.contract.UnpackLog(event, "AuthorizedMinterRemoved", log); err != nil {
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

// ParseAuthorizedMinterRemoved is a log parse operation binding the contract event 0xc6711413797b8a562634e98c95d50e7619d39702ed5b82ce335dc93546c3a88c.
//
// Solidity: event AuthorizedMinterRemoved(address indexed minter)
func (_Magiceden *MagicedenFilterer) ParseAuthorizedMinterRemoved(log types.Log) (*MagicedenAuthorizedMinterRemoved, error) {
	event := new(MagicedenAuthorizedMinterRemoved)
	if err := _Magiceden.contract.UnpackLog(event, "AuthorizedMinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenAutomaticApprovalOfTransferValidatorSetIterator is returned from FilterAutomaticApprovalOfTransferValidatorSet and is used to iterate over the raw logs and unpacked data for AutomaticApprovalOfTransferValidatorSet events raised by the Magiceden contract.
type MagicedenAutomaticApprovalOfTransferValidatorSetIterator struct {
	Event *MagicedenAutomaticApprovalOfTransferValidatorSet // Event containing the contract specifics and raw log

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
func (it *MagicedenAutomaticApprovalOfTransferValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenAutomaticApprovalOfTransferValidatorSet)
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
		it.Event = new(MagicedenAutomaticApprovalOfTransferValidatorSet)
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
func (it *MagicedenAutomaticApprovalOfTransferValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenAutomaticApprovalOfTransferValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenAutomaticApprovalOfTransferValidatorSet represents a AutomaticApprovalOfTransferValidatorSet event raised by the Magiceden contract.
type MagicedenAutomaticApprovalOfTransferValidatorSet struct {
	AutoApproved bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAutomaticApprovalOfTransferValidatorSet is a free log retrieval operation binding the contract event 0x6787c7f9a80aa0f5ceddab2c54f1f5169c0b88e75dd5e19d5e858a64144c7dbc.
//
// Solidity: event AutomaticApprovalOfTransferValidatorSet(bool autoApproved)
func (_Magiceden *MagicedenFilterer) FilterAutomaticApprovalOfTransferValidatorSet(opts *bind.FilterOpts) (*MagicedenAutomaticApprovalOfTransferValidatorSetIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "AutomaticApprovalOfTransferValidatorSet")
	if err != nil {
		return nil, err
	}
	return &MagicedenAutomaticApprovalOfTransferValidatorSetIterator{contract: _Magiceden.contract, event: "AutomaticApprovalOfTransferValidatorSet", logs: logs, sub: sub}, nil
}

// WatchAutomaticApprovalOfTransferValidatorSet is a free log subscription operation binding the contract event 0x6787c7f9a80aa0f5ceddab2c54f1f5169c0b88e75dd5e19d5e858a64144c7dbc.
//
// Solidity: event AutomaticApprovalOfTransferValidatorSet(bool autoApproved)
func (_Magiceden *MagicedenFilterer) WatchAutomaticApprovalOfTransferValidatorSet(opts *bind.WatchOpts, sink chan<- *MagicedenAutomaticApprovalOfTransferValidatorSet) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "AutomaticApprovalOfTransferValidatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenAutomaticApprovalOfTransferValidatorSet)
				if err := _Magiceden.contract.UnpackLog(event, "AutomaticApprovalOfTransferValidatorSet", log); err != nil {
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

// ParseAutomaticApprovalOfTransferValidatorSet is a log parse operation binding the contract event 0x6787c7f9a80aa0f5ceddab2c54f1f5169c0b88e75dd5e19d5e858a64144c7dbc.
//
// Solidity: event AutomaticApprovalOfTransferValidatorSet(bool autoApproved)
func (_Magiceden *MagicedenFilterer) ParseAutomaticApprovalOfTransferValidatorSet(log types.Log) (*MagicedenAutomaticApprovalOfTransferValidatorSet, error) {
	event := new(MagicedenAutomaticApprovalOfTransferValidatorSet)
	if err := _Magiceden.contract.UnpackLog(event, "AutomaticApprovalOfTransferValidatorSet", log); err != nil {
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

// MagicedenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Magiceden contract.
type MagicedenInitializedIterator struct {
	Event *MagicedenInitialized // Event containing the contract specifics and raw log

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
func (it *MagicedenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenInitialized)
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
		it.Event = new(MagicedenInitialized)
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
func (it *MagicedenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenInitialized represents a Initialized event raised by the Magiceden contract.
type MagicedenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Magiceden *MagicedenFilterer) FilterInitialized(opts *bind.FilterOpts) (*MagicedenInitializedIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MagicedenInitializedIterator{contract: _Magiceden.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Magiceden *MagicedenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MagicedenInitialized) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenInitialized)
				if err := _Magiceden.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Magiceden *MagicedenFilterer) ParseInitialized(log types.Log) (*MagicedenInitialized, error) {
	event := new(MagicedenInitialized)
	if err := _Magiceden.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Magiceden contract.
type MagicedenOwnershipHandoverCanceledIterator struct {
	Event *MagicedenOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *MagicedenOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenOwnershipHandoverCanceled)
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
		it.Event = new(MagicedenOwnershipHandoverCanceled)
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
func (it *MagicedenOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Magiceden contract.
type MagicedenOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Magiceden *MagicedenFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*MagicedenOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenOwnershipHandoverCanceledIterator{contract: _Magiceden.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Magiceden *MagicedenFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *MagicedenOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenOwnershipHandoverCanceled)
				if err := _Magiceden.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Magiceden *MagicedenFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*MagicedenOwnershipHandoverCanceled, error) {
	event := new(MagicedenOwnershipHandoverCanceled)
	if err := _Magiceden.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Magiceden contract.
type MagicedenOwnershipHandoverRequestedIterator struct {
	Event *MagicedenOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *MagicedenOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenOwnershipHandoverRequested)
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
		it.Event = new(MagicedenOwnershipHandoverRequested)
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
func (it *MagicedenOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Magiceden contract.
type MagicedenOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Magiceden *MagicedenFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*MagicedenOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenOwnershipHandoverRequestedIterator{contract: _Magiceden.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Magiceden *MagicedenFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *MagicedenOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenOwnershipHandoverRequested)
				if err := _Magiceden.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Magiceden *MagicedenFilterer) ParseOwnershipHandoverRequested(log types.Log) (*MagicedenOwnershipHandoverRequested, error) {
	event := new(MagicedenOwnershipHandoverRequested)
	if err := _Magiceden.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Magiceden *MagicedenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*MagicedenOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenOwnershipTransferredIterator{contract: _Magiceden.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Magiceden *MagicedenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MagicedenOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Magiceden *MagicedenFilterer) ParseOwnershipTransferred(log types.Log) (*MagicedenOwnershipTransferred, error) {
	event := new(MagicedenOwnershipTransferred)
	if err := _Magiceden.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MagicedenSetContractURIIterator is returned from FilterSetContractURI and is used to iterate over the raw logs and unpacked data for SetContractURI events raised by the Magiceden contract.
type MagicedenSetContractURIIterator struct {
	Event *MagicedenSetContractURI // Event containing the contract specifics and raw log

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
func (it *MagicedenSetContractURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetContractURI)
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
		it.Event = new(MagicedenSetContractURI)
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
func (it *MagicedenSetContractURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetContractURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetContractURI represents a SetContractURI event raised by the Magiceden contract.
type MagicedenSetContractURI struct {
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetContractURI is a free log retrieval operation binding the contract event 0x5ca9f750836b0b7efdace104f07b5c9f0df0650c0fd24f5163e99044ae36ea52.
//
// Solidity: event SetContractURI(string uri)
func (_Magiceden *MagicedenFilterer) FilterSetContractURI(opts *bind.FilterOpts) (*MagicedenSetContractURIIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetContractURI")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetContractURIIterator{contract: _Magiceden.contract, event: "SetContractURI", logs: logs, sub: sub}, nil
}

// WatchSetContractURI is a free log subscription operation binding the contract event 0x5ca9f750836b0b7efdace104f07b5c9f0df0650c0fd24f5163e99044ae36ea52.
//
// Solidity: event SetContractURI(string uri)
func (_Magiceden *MagicedenFilterer) WatchSetContractURI(opts *bind.WatchOpts, sink chan<- *MagicedenSetContractURI) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetContractURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetContractURI)
				if err := _Magiceden.contract.UnpackLog(event, "SetContractURI", log); err != nil {
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

// ParseSetContractURI is a log parse operation binding the contract event 0x5ca9f750836b0b7efdace104f07b5c9f0df0650c0fd24f5163e99044ae36ea52.
//
// Solidity: event SetContractURI(string uri)
func (_Magiceden *MagicedenFilterer) ParseSetContractURI(log types.Log) (*MagicedenSetContractURI, error) {
	event := new(MagicedenSetContractURI)
	if err := _Magiceden.contract.UnpackLog(event, "SetContractURI", log); err != nil {
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
// Solidity: event SetCosigner(address indexed cosigner)
func (_Magiceden *MagicedenFilterer) FilterSetCosigner(opts *bind.FilterOpts, cosigner []common.Address) (*MagicedenSetCosignerIterator, error) {

	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetCosigner", cosignerRule)
	if err != nil {
		return nil, err
	}
	return &MagicedenSetCosignerIterator{contract: _Magiceden.contract, event: "SetCosigner", logs: logs, sub: sub}, nil
}

// WatchSetCosigner is a free log subscription operation binding the contract event 0xaea1573caf7b4fdd079b947d86c1be6c725642c47582f8f9bd2c7d2a30bf0bd9.
//
// Solidity: event SetCosigner(address indexed cosigner)
func (_Magiceden *MagicedenFilterer) WatchSetCosigner(opts *bind.WatchOpts, sink chan<- *MagicedenSetCosigner, cosigner []common.Address) (event.Subscription, error) {

	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetCosigner", cosignerRule)
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
// Solidity: event SetCosigner(address indexed cosigner)
func (_Magiceden *MagicedenFilterer) ParseSetCosigner(log types.Log) (*MagicedenSetCosigner, error) {
	event := new(MagicedenSetCosigner)
	if err := _Magiceden.contract.UnpackLog(event, "SetCosigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetDefaultRoyaltyIterator is returned from FilterSetDefaultRoyalty and is used to iterate over the raw logs and unpacked data for SetDefaultRoyalty events raised by the Magiceden contract.
type MagicedenSetDefaultRoyaltyIterator struct {
	Event *MagicedenSetDefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *MagicedenSetDefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetDefaultRoyalty)
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
		it.Event = new(MagicedenSetDefaultRoyalty)
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
func (it *MagicedenSetDefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetDefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetDefaultRoyalty represents a SetDefaultRoyalty event raised by the Magiceden contract.
type MagicedenSetDefaultRoyalty struct {
	Receiver     common.Address
	FeeNumerator *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultRoyalty is a free log retrieval operation binding the contract event 0xa1edde4ed5c1392c90dccd8e051a4080b761850e49a24c77d826348a51e1f8dc.
//
// Solidity: event SetDefaultRoyalty(address receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) FilterSetDefaultRoyalty(opts *bind.FilterOpts) (*MagicedenSetDefaultRoyaltyIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetDefaultRoyalty")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetDefaultRoyaltyIterator{contract: _Magiceden.contract, event: "SetDefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchSetDefaultRoyalty is a free log subscription operation binding the contract event 0xa1edde4ed5c1392c90dccd8e051a4080b761850e49a24c77d826348a51e1f8dc.
//
// Solidity: event SetDefaultRoyalty(address receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) WatchSetDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *MagicedenSetDefaultRoyalty) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetDefaultRoyalty")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetDefaultRoyalty)
				if err := _Magiceden.contract.UnpackLog(event, "SetDefaultRoyalty", log); err != nil {
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

// ParseSetDefaultRoyalty is a log parse operation binding the contract event 0xa1edde4ed5c1392c90dccd8e051a4080b761850e49a24c77d826348a51e1f8dc.
//
// Solidity: event SetDefaultRoyalty(address receiver, uint96 feeNumerator)
func (_Magiceden *MagicedenFilterer) ParseSetDefaultRoyalty(log types.Log) (*MagicedenSetDefaultRoyalty, error) {
	event := new(MagicedenSetDefaultRoyalty)
	if err := _Magiceden.contract.UnpackLog(event, "SetDefaultRoyalty", log); err != nil {
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
	TimestampExpirySeconds *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTimestampExpirySeconds is a free log retrieval operation binding the contract event 0x0f1c5629c9ab6d9b97fd6801d012d74903c2eab7df1abec22bb54f9a05547645.
//
// Solidity: event SetTimestampExpirySeconds(uint256 timestampExpirySeconds)
func (_Magiceden *MagicedenFilterer) FilterSetTimestampExpirySeconds(opts *bind.FilterOpts) (*MagicedenSetTimestampExpirySecondsIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetTimestampExpirySeconds")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetTimestampExpirySecondsIterator{contract: _Magiceden.contract, event: "SetTimestampExpirySeconds", logs: logs, sub: sub}, nil
}

// WatchSetTimestampExpirySeconds is a free log subscription operation binding the contract event 0x0f1c5629c9ab6d9b97fd6801d012d74903c2eab7df1abec22bb54f9a05547645.
//
// Solidity: event SetTimestampExpirySeconds(uint256 timestampExpirySeconds)
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

// ParseSetTimestampExpirySeconds is a log parse operation binding the contract event 0x0f1c5629c9ab6d9b97fd6801d012d74903c2eab7df1abec22bb54f9a05547645.
//
// Solidity: event SetTimestampExpirySeconds(uint256 timestampExpirySeconds)
func (_Magiceden *MagicedenFilterer) ParseSetTimestampExpirySeconds(log types.Log) (*MagicedenSetTimestampExpirySeconds, error) {
	event := new(MagicedenSetTimestampExpirySeconds)
	if err := _Magiceden.contract.UnpackLog(event, "SetTimestampExpirySeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicedenSetTokenURISuffixIterator is returned from FilterSetTokenURISuffix and is used to iterate over the raw logs and unpacked data for SetTokenURISuffix events raised by the Magiceden contract.
type MagicedenSetTokenURISuffixIterator struct {
	Event *MagicedenSetTokenURISuffix // Event containing the contract specifics and raw log

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
func (it *MagicedenSetTokenURISuffixIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicedenSetTokenURISuffix)
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
		it.Event = new(MagicedenSetTokenURISuffix)
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
func (it *MagicedenSetTokenURISuffixIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicedenSetTokenURISuffixIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicedenSetTokenURISuffix represents a SetTokenURISuffix event raised by the Magiceden contract.
type MagicedenSetTokenURISuffix struct {
	Suffix string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetTokenURISuffix is a free log retrieval operation binding the contract event 0x55640d8d12423d0ae4189f0be9ab48d5682c83ac3c23eae8a799459da358e5f4.
//
// Solidity: event SetTokenURISuffix(string suffix)
func (_Magiceden *MagicedenFilterer) FilterSetTokenURISuffix(opts *bind.FilterOpts) (*MagicedenSetTokenURISuffixIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "SetTokenURISuffix")
	if err != nil {
		return nil, err
	}
	return &MagicedenSetTokenURISuffixIterator{contract: _Magiceden.contract, event: "SetTokenURISuffix", logs: logs, sub: sub}, nil
}

// WatchSetTokenURISuffix is a free log subscription operation binding the contract event 0x55640d8d12423d0ae4189f0be9ab48d5682c83ac3c23eae8a799459da358e5f4.
//
// Solidity: event SetTokenURISuffix(string suffix)
func (_Magiceden *MagicedenFilterer) WatchSetTokenURISuffix(opts *bind.WatchOpts, sink chan<- *MagicedenSetTokenURISuffix) (event.Subscription, error) {

	logs, sub, err := _Magiceden.contract.WatchLogs(opts, "SetTokenURISuffix")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicedenSetTokenURISuffix)
				if err := _Magiceden.contract.UnpackLog(event, "SetTokenURISuffix", log); err != nil {
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

// ParseSetTokenURISuffix is a log parse operation binding the contract event 0x55640d8d12423d0ae4189f0be9ab48d5682c83ac3c23eae8a799459da358e5f4.
//
// Solidity: event SetTokenURISuffix(string suffix)
func (_Magiceden *MagicedenFilterer) ParseSetTokenURISuffix(log types.Log) (*MagicedenSetTokenURISuffix, error) {
	event := new(MagicedenSetTokenURISuffix)
	if err := _Magiceden.contract.UnpackLog(event, "SetTokenURISuffix", log); err != nil {
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
	MintFee              *big.Int
	WalletLimit          uint32
	MerkleRoot           [32]byte
	MaxStageSupply       *big.Int
	StartTimeUnixSeconds *big.Int
	EndTimeUnixSeconds   *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateStage is a free log retrieval operation binding the contract event 0xd4d1f42a08182d6508cb22b382bf9fb146c3fd9f27dcbbcbd5d6b6130e4283c7.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint80 mintFee, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint256 startTimeUnixSeconds, uint256 endTimeUnixSeconds)
func (_Magiceden *MagicedenFilterer) FilterUpdateStage(opts *bind.FilterOpts) (*MagicedenUpdateStageIterator, error) {

	logs, sub, err := _Magiceden.contract.FilterLogs(opts, "UpdateStage")
	if err != nil {
		return nil, err
	}
	return &MagicedenUpdateStageIterator{contract: _Magiceden.contract, event: "UpdateStage", logs: logs, sub: sub}, nil
}

// WatchUpdateStage is a free log subscription operation binding the contract event 0xd4d1f42a08182d6508cb22b382bf9fb146c3fd9f27dcbbcbd5d6b6130e4283c7.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint80 mintFee, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint256 startTimeUnixSeconds, uint256 endTimeUnixSeconds)
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

// ParseUpdateStage is a log parse operation binding the contract event 0xd4d1f42a08182d6508cb22b382bf9fb146c3fd9f27dcbbcbd5d6b6130e4283c7.
//
// Solidity: event UpdateStage(uint256 stage, uint80 price, uint80 mintFee, uint32 walletLimit, bytes32 merkleRoot, uint24 maxStageSupply, uint256 startTimeUnixSeconds, uint256 endTimeUnixSeconds)
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
