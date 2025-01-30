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

// AllowListData is an auto generated low-level Go binding around an user-defined struct.
type AllowListData struct {
	MerkleRoot    [32]byte
	PublicKeyURIs []string
	AllowListURI  string
}

// MintParams is an auto generated low-level Go binding around an user-defined struct.
type MintParams struct {
	MintPrice                *big.Int
	MaxTotalMintableByWallet *big.Int
	StartTime                *big.Int
	EndTime                  *big.Int
	DropStageIndex           *big.Int
	MaxTokenSupplyForStage   *big.Int
	FeeBps                   *big.Int
	RestrictFeeRecipients    bool
}

// PublicDrop is an auto generated low-level Go binding around an user-defined struct.
type PublicDrop struct {
	MintPrice                *big.Int
	StartTime                *big.Int
	EndTime                  *big.Int
	MaxTotalMintableByWallet uint16
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// SignedMintValidationParams is an auto generated low-level Go binding around an user-defined struct.
type SignedMintValidationParams struct {
	MinMintPrice                *big.Int
	MaxMaxTotalMintableByWallet *big.Int
	MinStartTime                *big.Int
	MaxEndTime                  *big.Int
	MaxMaxTokenSupplyForStage   *big.Int
	MinFeeBps                   uint16
	MaxFeeBps                   uint16
}

// TokenGatedDropStage is an auto generated low-level Go binding around an user-defined struct.
type TokenGatedDropStage struct {
	MintPrice                *big.Int
	MaxTotalMintableByWallet uint16
	StartTime                *big.Int
	EndTime                  *big.Int
	DropStageIndex           uint8
	MaxTokenSupplyForStage   uint32
	FeeBps                   uint16
	RestrictFeeRecipients    bool
}

// TokenGatedMintParams is an auto generated low-level Go binding around an user-defined struct.
type TokenGatedMintParams struct {
	AllowedNftToken    common.Address
	AllowedNftTokenIds []*big.Int
}

// OpenseaMetaData contains all meta data concerning the Opensea contract.
var OpenseaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CreatorPayoutAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicatePayer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"IncorrectPayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveredSigner\",\"type\":\"address\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedEndTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumOrMaximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMaxTotalMintableByWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedMintPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"InvalidSignedStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintQuantityCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowed\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxMintedPerWallet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"}],\"name\":\"MintQuantityExceedsMaxTokenSupplyForStage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"}],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"OnlyINonFungibleSeaDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignedMintsMustRestrictFeeRecipients\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerNotPresent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeDropToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropAllowedNftTokenCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenGatedDropStageNotPresent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedNotTokenOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"TokenGatedTokenIdAlreadyRedeemed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"publicKeyURI\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"name\":\"AllowListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedFeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPayoutAddress\",\"type\":\"address\"}],\"name\":\"CreatorPayoutAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDropURI\",\"type\":\"string\"}],\"name\":\"DropURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"PayerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"PublicDropUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unitMintPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"}],\"name\":\"SeaDropMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"SignedMintValidationParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"TokenGatedDropStageUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getAllowListMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getAllowedFeeRecipients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowedNftTokenId\",\"type\":\"uint256\"}],\"name\":\"getAllowedNftTokenIdIsRedeemed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getCreatorPayoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"getFeeRecipientIsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"}],\"name\":\"getPayerIsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getPayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getPublicDrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getSignedMintValidationParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"getTokenGatedAllowedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"}],\"name\":\"getTokenGatedDrop\",\"outputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"mintAllowList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"allowedNftTokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structTokenGatedMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"}],\"name\":\"mintAllowedTokenHolder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"mintPublic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minterIfNotPayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dropStageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structMintParams\",\"name\":\"mintParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintSigned\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"publicKeyURIs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"allowListURI\",\"type\":\"string\"}],\"internalType\":\"structAllowListData\",\"name\":\"allowListData\",\"type\":\"tuple\"}],\"name\":\"updateAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateAllowedFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payoutAddress\",\"type\":\"address\"}],\"name\":\"updateCreatorPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dropURI\",\"type\":\"string\"}],\"name\":\"updateDropURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updatePayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structPublicDrop\",\"name\":\"publicDrop\",\"type\":\"tuple\"}],\"name\":\"updatePublicDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"minMintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint24\",\"name\":\"maxMaxTotalMintableByWallet\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"minStartTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxEndTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"maxMaxTokenSupplyForStage\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minFeeBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxFeeBps\",\"type\":\"uint16\"}],\"internalType\":\"structSignedMintValidationParams\",\"name\":\"signedMintValidationParams\",\"type\":\"tuple\"}],\"name\":\"updateSignedMintValidationParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"allowedNftToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint80\",\"name\":\"mintPrice\",\"type\":\"uint80\"},{\"internalType\":\"uint16\",\"name\":\"maxTotalMintableByWallet\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"startTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTime\",\"type\":\"uint48\"},{\"internalType\":\"uint8\",\"name\":\"dropStageIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenSupplyForStage\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"feeBps\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"restrictFeeRecipients\",\"type\":\"bool\"}],\"internalType\":\"structTokenGatedDropStage\",\"name\":\"dropStage\",\"type\":\"tuple\"}],\"name\":\"updateTokenGatedDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OpenseaABI is the input ABI used to generate the binding from.
// Deprecated: Use OpenseaMetaData.ABI instead.
var OpenseaABI = OpenseaMetaData.ABI

// Opensea is an auto generated Go binding around an Ethereum contract.
type Opensea struct {
	OpenseaCaller     // Read-only binding to the contract
	OpenseaTransactor // Write-only binding to the contract
	OpenseaFilterer   // Log filterer for contract events
}

// OpenseaCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpenseaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpenseaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpenseaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenseaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpenseaSession struct {
	Contract     *Opensea          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenseaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpenseaCallerSession struct {
	Contract *OpenseaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OpenseaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpenseaTransactorSession struct {
	Contract     *OpenseaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OpenseaRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpenseaRaw struct {
	Contract *Opensea // Generic contract binding to access the raw methods on
}

// OpenseaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpenseaCallerRaw struct {
	Contract *OpenseaCaller // Generic read-only contract binding to access the raw methods on
}

// OpenseaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpenseaTransactorRaw struct {
	Contract *OpenseaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpensea creates a new instance of Opensea, bound to a specific deployed contract.
func NewOpensea(address common.Address, backend bind.ContractBackend) (*Opensea, error) {
	contract, err := bindOpensea(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Opensea{OpenseaCaller: OpenseaCaller{contract: contract}, OpenseaTransactor: OpenseaTransactor{contract: contract}, OpenseaFilterer: OpenseaFilterer{contract: contract}}, nil
}

// NewOpenseaCaller creates a new read-only instance of Opensea, bound to a specific deployed contract.
func NewOpenseaCaller(address common.Address, caller bind.ContractCaller) (*OpenseaCaller, error) {
	contract, err := bindOpensea(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaCaller{contract: contract}, nil
}

// NewOpenseaTransactor creates a new write-only instance of Opensea, bound to a specific deployed contract.
func NewOpenseaTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenseaTransactor, error) {
	contract, err := bindOpensea(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenseaTransactor{contract: contract}, nil
}

// NewOpenseaFilterer creates a new log filterer instance of Opensea, bound to a specific deployed contract.
func NewOpenseaFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenseaFilterer, error) {
	contract, err := bindOpensea(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenseaFilterer{contract: contract}, nil
}

// bindOpensea binds a generic wrapper to an already deployed contract.
func bindOpensea(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpenseaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opensea *OpenseaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opensea.Contract.OpenseaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opensea *OpenseaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.Contract.OpenseaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opensea *OpenseaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opensea.Contract.OpenseaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Opensea *OpenseaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Opensea.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Opensea *OpenseaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Opensea.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Opensea *OpenseaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Opensea.Contract.contract.Transact(opts, method, params...)
}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Opensea *OpenseaCaller) GetAllowListMerkleRoot(opts *bind.CallOpts, nftContract common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getAllowListMerkleRoot", nftContract)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Opensea *OpenseaSession) GetAllowListMerkleRoot(nftContract common.Address) ([32]byte, error) {
	return _Opensea.Contract.GetAllowListMerkleRoot(&_Opensea.CallOpts, nftContract)
}

// GetAllowListMerkleRoot is a free data retrieval call binding the contract method 0x32bf11f5.
//
// Solidity: function getAllowListMerkleRoot(address nftContract) view returns(bytes32)
func (_Opensea *OpenseaCallerSession) GetAllowListMerkleRoot(nftContract common.Address) ([32]byte, error) {
	return _Opensea.Contract.GetAllowListMerkleRoot(&_Opensea.CallOpts, nftContract)
}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Opensea *OpenseaCaller) GetAllowedFeeRecipients(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getAllowedFeeRecipients", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Opensea *OpenseaSession) GetAllowedFeeRecipients(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetAllowedFeeRecipients(&_Opensea.CallOpts, nftContract)
}

// GetAllowedFeeRecipients is a free data retrieval call binding the contract method 0x68632274.
//
// Solidity: function getAllowedFeeRecipients(address nftContract) view returns(address[])
func (_Opensea *OpenseaCallerSession) GetAllowedFeeRecipients(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetAllowedFeeRecipients(&_Opensea.CallOpts, nftContract)
}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Opensea *OpenseaCaller) GetAllowedNftTokenIdIsRedeemed(opts *bind.CallOpts, nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getAllowedNftTokenIdIsRedeemed", nftContract, allowedNftToken, allowedNftTokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Opensea *OpenseaSession) GetAllowedNftTokenIdIsRedeemed(nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	return _Opensea.Contract.GetAllowedNftTokenIdIsRedeemed(&_Opensea.CallOpts, nftContract, allowedNftToken, allowedNftTokenId)
}

// GetAllowedNftTokenIdIsRedeemed is a free data retrieval call binding the contract method 0x88aa3d37.
//
// Solidity: function getAllowedNftTokenIdIsRedeemed(address nftContract, address allowedNftToken, uint256 allowedNftTokenId) view returns(bool)
func (_Opensea *OpenseaCallerSession) GetAllowedNftTokenIdIsRedeemed(nftContract common.Address, allowedNftToken common.Address, allowedNftTokenId *big.Int) (bool, error) {
	return _Opensea.Contract.GetAllowedNftTokenIdIsRedeemed(&_Opensea.CallOpts, nftContract, allowedNftToken, allowedNftTokenId)
}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Opensea *OpenseaCaller) GetCreatorPayoutAddress(opts *bind.CallOpts, nftContract common.Address) (common.Address, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getCreatorPayoutAddress", nftContract)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Opensea *OpenseaSession) GetCreatorPayoutAddress(nftContract common.Address) (common.Address, error) {
	return _Opensea.Contract.GetCreatorPayoutAddress(&_Opensea.CallOpts, nftContract)
}

// GetCreatorPayoutAddress is a free data retrieval call binding the contract method 0x5cb3c4d3.
//
// Solidity: function getCreatorPayoutAddress(address nftContract) view returns(address)
func (_Opensea *OpenseaCallerSession) GetCreatorPayoutAddress(nftContract common.Address) (common.Address, error) {
	return _Opensea.Contract.GetCreatorPayoutAddress(&_Opensea.CallOpts, nftContract)
}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Opensea *OpenseaCaller) GetFeeRecipientIsAllowed(opts *bind.CallOpts, nftContract common.Address, feeRecipient common.Address) (bool, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getFeeRecipientIsAllowed", nftContract, feeRecipient)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Opensea *OpenseaSession) GetFeeRecipientIsAllowed(nftContract common.Address, feeRecipient common.Address) (bool, error) {
	return _Opensea.Contract.GetFeeRecipientIsAllowed(&_Opensea.CallOpts, nftContract, feeRecipient)
}

// GetFeeRecipientIsAllowed is a free data retrieval call binding the contract method 0x322e75d1.
//
// Solidity: function getFeeRecipientIsAllowed(address nftContract, address feeRecipient) view returns(bool)
func (_Opensea *OpenseaCallerSession) GetFeeRecipientIsAllowed(nftContract common.Address, feeRecipient common.Address) (bool, error) {
	return _Opensea.Contract.GetFeeRecipientIsAllowed(&_Opensea.CallOpts, nftContract, feeRecipient)
}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Opensea *OpenseaCaller) GetPayerIsAllowed(opts *bind.CallOpts, nftContract common.Address, payer common.Address) (bool, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getPayerIsAllowed", nftContract, payer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Opensea *OpenseaSession) GetPayerIsAllowed(nftContract common.Address, payer common.Address) (bool, error) {
	return _Opensea.Contract.GetPayerIsAllowed(&_Opensea.CallOpts, nftContract, payer)
}

// GetPayerIsAllowed is a free data retrieval call binding the contract method 0xe583141d.
//
// Solidity: function getPayerIsAllowed(address nftContract, address payer) view returns(bool)
func (_Opensea *OpenseaCallerSession) GetPayerIsAllowed(nftContract common.Address, payer common.Address) (bool, error) {
	return _Opensea.Contract.GetPayerIsAllowed(&_Opensea.CallOpts, nftContract, payer)
}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Opensea *OpenseaCaller) GetPayers(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getPayers", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Opensea *OpenseaSession) GetPayers(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetPayers(&_Opensea.CallOpts, nftContract)
}

// GetPayers is a free data retrieval call binding the contract method 0x7c35b982.
//
// Solidity: function getPayers(address nftContract) view returns(address[])
func (_Opensea *OpenseaCallerSession) GetPayers(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetPayers(&_Opensea.CallOpts, nftContract)
}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Opensea *OpenseaCaller) GetPublicDrop(opts *bind.CallOpts, nftContract common.Address) (PublicDrop, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getPublicDrop", nftContract)

	if err != nil {
		return *new(PublicDrop), err
	}

	out0 := *abi.ConvertType(out[0], new(PublicDrop)).(*PublicDrop)

	return out0, err

}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Opensea *OpenseaSession) GetPublicDrop(nftContract common.Address) (PublicDrop, error) {
	return _Opensea.Contract.GetPublicDrop(&_Opensea.CallOpts, nftContract)
}

// GetPublicDrop is a free data retrieval call binding the contract method 0xbc6a629c.
//
// Solidity: function getPublicDrop(address nftContract) view returns((uint80,uint48,uint48,uint16,uint16,bool))
func (_Opensea *OpenseaCallerSession) GetPublicDrop(nftContract common.Address) (PublicDrop, error) {
	return _Opensea.Contract.GetPublicDrop(&_Opensea.CallOpts, nftContract)
}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Opensea *OpenseaCaller) GetSignedMintValidationParams(opts *bind.CallOpts, nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getSignedMintValidationParams", nftContract, signer)

	if err != nil {
		return *new(SignedMintValidationParams), err
	}

	out0 := *abi.ConvertType(out[0], new(SignedMintValidationParams)).(*SignedMintValidationParams)

	return out0, err

}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Opensea *OpenseaSession) GetSignedMintValidationParams(nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	return _Opensea.Contract.GetSignedMintValidationParams(&_Opensea.CallOpts, nftContract, signer)
}

// GetSignedMintValidationParams is a free data retrieval call binding the contract method 0x81bf9af3.
//
// Solidity: function getSignedMintValidationParams(address nftContract, address signer) view returns((uint80,uint24,uint40,uint40,uint40,uint16,uint16))
func (_Opensea *OpenseaCallerSession) GetSignedMintValidationParams(nftContract common.Address, signer common.Address) (SignedMintValidationParams, error) {
	return _Opensea.Contract.GetSignedMintValidationParams(&_Opensea.CallOpts, nftContract, signer)
}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Opensea *OpenseaCaller) GetSigners(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getSigners", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Opensea *OpenseaSession) GetSigners(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetSigners(&_Opensea.CallOpts, nftContract)
}

// GetSigners is a free data retrieval call binding the contract method 0x7e3ba6af.
//
// Solidity: function getSigners(address nftContract) view returns(address[])
func (_Opensea *OpenseaCallerSession) GetSigners(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetSigners(&_Opensea.CallOpts, nftContract)
}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Opensea *OpenseaCaller) GetTokenGatedAllowedTokens(opts *bind.CallOpts, nftContract common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getTokenGatedAllowedTokens", nftContract)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Opensea *OpenseaSession) GetTokenGatedAllowedTokens(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetTokenGatedAllowedTokens(&_Opensea.CallOpts, nftContract)
}

// GetTokenGatedAllowedTokens is a free data retrieval call binding the contract method 0x2db526eb.
//
// Solidity: function getTokenGatedAllowedTokens(address nftContract) view returns(address[])
func (_Opensea *OpenseaCallerSession) GetTokenGatedAllowedTokens(nftContract common.Address) ([]common.Address, error) {
	return _Opensea.Contract.GetTokenGatedAllowedTokens(&_Opensea.CallOpts, nftContract)
}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Opensea *OpenseaCaller) GetTokenGatedDrop(opts *bind.CallOpts, nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	var out []interface{}
	err := _Opensea.contract.Call(opts, &out, "getTokenGatedDrop", nftContract, allowedNftToken)

	if err != nil {
		return *new(TokenGatedDropStage), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenGatedDropStage)).(*TokenGatedDropStage)

	return out0, err

}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Opensea *OpenseaSession) GetTokenGatedDrop(nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	return _Opensea.Contract.GetTokenGatedDrop(&_Opensea.CallOpts, nftContract, allowedNftToken)
}

// GetTokenGatedDrop is a free data retrieval call binding the contract method 0x0b0e8a6e.
//
// Solidity: function getTokenGatedDrop(address nftContract, address allowedNftToken) view returns((uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool))
func (_Opensea *OpenseaCallerSession) GetTokenGatedDrop(nftContract common.Address, allowedNftToken common.Address) (TokenGatedDropStage, error) {
	return _Opensea.Contract.GetTokenGatedDrop(&_Opensea.CallOpts, nftContract, allowedNftToken)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Opensea *OpenseaTransactor) MintAllowList(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "mintAllowList", nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Opensea *OpenseaSession) MintAllowList(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Opensea.Contract.MintAllowList(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowList is a paid mutator transaction binding the contract method 0x4300a4e6.
//
// Solidity: function mintAllowList(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, bytes32[] proof) payable returns()
func (_Opensea *OpenseaTransactorSession) MintAllowList(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, proof [][32]byte) (*types.Transaction, error) {
	return _Opensea.Contract.MintAllowList(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, proof)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Opensea *OpenseaTransactor) MintAllowedTokenHolder(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "mintAllowedTokenHolder", nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Opensea *OpenseaSession) MintAllowedTokenHolder(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Opensea.Contract.MintAllowedTokenHolder(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintAllowedTokenHolder is a paid mutator transaction binding the contract method 0x99eb900f.
//
// Solidity: function mintAllowedTokenHolder(address nftContract, address feeRecipient, address minterIfNotPayer, (address,uint256[]) mintParams) payable returns()
func (_Opensea *OpenseaTransactorSession) MintAllowedTokenHolder(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, mintParams TokenGatedMintParams) (*types.Transaction, error) {
	return _Opensea.Contract.MintAllowedTokenHolder(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, mintParams)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Opensea *OpenseaTransactor) MintPublic(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "mintPublic", nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Opensea *OpenseaSession) MintPublic(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Opensea.Contract.MintPublic(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintPublic is a paid mutator transaction binding the contract method 0x161ac21f.
//
// Solidity: function mintPublic(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity) payable returns()
func (_Opensea *OpenseaTransactorSession) MintPublic(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Opensea.Contract.MintPublic(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Opensea *OpenseaTransactor) MintSigned(opts *bind.TransactOpts, nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "mintSigned", nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Opensea *OpenseaSession) MintSigned(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Opensea.Contract.MintSigned(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// MintSigned is a paid mutator transaction binding the contract method 0x4b61cd6f.
//
// Solidity: function mintSigned(address nftContract, address feeRecipient, address minterIfNotPayer, uint256 quantity, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool) mintParams, uint256 salt, bytes signature) payable returns()
func (_Opensea *OpenseaTransactorSession) MintSigned(nftContract common.Address, feeRecipient common.Address, minterIfNotPayer common.Address, quantity *big.Int, mintParams MintParams, salt *big.Int, signature []byte) (*types.Transaction, error) {
	return _Opensea.Contract.MintSigned(&_Opensea.TransactOpts, nftContract, feeRecipient, minterIfNotPayer, quantity, mintParams, salt, signature)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Opensea *OpenseaTransactor) UpdateAllowList(opts *bind.TransactOpts, allowListData AllowListData) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updateAllowList", allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Opensea *OpenseaSession) UpdateAllowList(allowListData AllowListData) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateAllowList(&_Opensea.TransactOpts, allowListData)
}

// UpdateAllowList is a paid mutator transaction binding the contract method 0xebb4a55f.
//
// Solidity: function updateAllowList((bytes32,string[],string) allowListData) returns()
func (_Opensea *OpenseaTransactorSession) UpdateAllowList(allowListData AllowListData) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateAllowList(&_Opensea.TransactOpts, allowListData)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Opensea *OpenseaTransactor) UpdateAllowedFeeRecipient(opts *bind.TransactOpts, feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updateAllowedFeeRecipient", feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Opensea *OpenseaSession) UpdateAllowedFeeRecipient(feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateAllowedFeeRecipient(&_Opensea.TransactOpts, feeRecipient, allowed)
}

// UpdateAllowedFeeRecipient is a paid mutator transaction binding the contract method 0x8e7d1e43.
//
// Solidity: function updateAllowedFeeRecipient(address feeRecipient, bool allowed) returns()
func (_Opensea *OpenseaTransactorSession) UpdateAllowedFeeRecipient(feeRecipient common.Address, allowed bool) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateAllowedFeeRecipient(&_Opensea.TransactOpts, feeRecipient, allowed)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address _payoutAddress) returns()
func (_Opensea *OpenseaTransactor) UpdateCreatorPayoutAddress(opts *bind.TransactOpts, _payoutAddress common.Address) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updateCreatorPayoutAddress", _payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address _payoutAddress) returns()
func (_Opensea *OpenseaSession) UpdateCreatorPayoutAddress(_payoutAddress common.Address) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateCreatorPayoutAddress(&_Opensea.TransactOpts, _payoutAddress)
}

// UpdateCreatorPayoutAddress is a paid mutator transaction binding the contract method 0x12738db8.
//
// Solidity: function updateCreatorPayoutAddress(address _payoutAddress) returns()
func (_Opensea *OpenseaTransactorSession) UpdateCreatorPayoutAddress(_payoutAddress common.Address) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateCreatorPayoutAddress(&_Opensea.TransactOpts, _payoutAddress)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Opensea *OpenseaTransactor) UpdateDropURI(opts *bind.TransactOpts, dropURI string) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updateDropURI", dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Opensea *OpenseaSession) UpdateDropURI(dropURI string) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateDropURI(&_Opensea.TransactOpts, dropURI)
}

// UpdateDropURI is a paid mutator transaction binding the contract method 0xb957d0cb.
//
// Solidity: function updateDropURI(string dropURI) returns()
func (_Opensea *OpenseaTransactorSession) UpdateDropURI(dropURI string) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateDropURI(&_Opensea.TransactOpts, dropURI)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Opensea *OpenseaTransactor) UpdatePayer(opts *bind.TransactOpts, payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updatePayer", payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Opensea *OpenseaSession) UpdatePayer(payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Opensea.Contract.UpdatePayer(&_Opensea.TransactOpts, payer, allowed)
}

// UpdatePayer is a paid mutator transaction binding the contract method 0x7f2a5cca.
//
// Solidity: function updatePayer(address payer, bool allowed) returns()
func (_Opensea *OpenseaTransactorSession) UpdatePayer(payer common.Address, allowed bool) (*types.Transaction, error) {
	return _Opensea.Contract.UpdatePayer(&_Opensea.TransactOpts, payer, allowed)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Opensea *OpenseaTransactor) UpdatePublicDrop(opts *bind.TransactOpts, publicDrop PublicDrop) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updatePublicDrop", publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Opensea *OpenseaSession) UpdatePublicDrop(publicDrop PublicDrop) (*types.Transaction, error) {
	return _Opensea.Contract.UpdatePublicDrop(&_Opensea.TransactOpts, publicDrop)
}

// UpdatePublicDrop is a paid mutator transaction binding the contract method 0x01308e65.
//
// Solidity: function updatePublicDrop((uint80,uint48,uint48,uint16,uint16,bool) publicDrop) returns()
func (_Opensea *OpenseaTransactorSession) UpdatePublicDrop(publicDrop PublicDrop) (*types.Transaction, error) {
	return _Opensea.Contract.UpdatePublicDrop(&_Opensea.TransactOpts, publicDrop)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Opensea *OpenseaTransactor) UpdateSignedMintValidationParams(opts *bind.TransactOpts, signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updateSignedMintValidationParams", signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Opensea *OpenseaSession) UpdateSignedMintValidationParams(signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateSignedMintValidationParams(&_Opensea.TransactOpts, signer, signedMintValidationParams)
}

// UpdateSignedMintValidationParams is a paid mutator transaction binding the contract method 0x4d380178.
//
// Solidity: function updateSignedMintValidationParams(address signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams) returns()
func (_Opensea *OpenseaTransactorSession) UpdateSignedMintValidationParams(signer common.Address, signedMintValidationParams SignedMintValidationParams) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateSignedMintValidationParams(&_Opensea.TransactOpts, signer, signedMintValidationParams)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Opensea *OpenseaTransactor) UpdateTokenGatedDrop(opts *bind.TransactOpts, allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Opensea.contract.Transact(opts, "updateTokenGatedDrop", allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Opensea *OpenseaSession) UpdateTokenGatedDrop(allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateTokenGatedDrop(&_Opensea.TransactOpts, allowedNftToken, dropStage)
}

// UpdateTokenGatedDrop is a paid mutator transaction binding the contract method 0xfd9ab22a.
//
// Solidity: function updateTokenGatedDrop(address allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage) returns()
func (_Opensea *OpenseaTransactorSession) UpdateTokenGatedDrop(allowedNftToken common.Address, dropStage TokenGatedDropStage) (*types.Transaction, error) {
	return _Opensea.Contract.UpdateTokenGatedDrop(&_Opensea.TransactOpts, allowedNftToken, dropStage)
}

// OpenseaAllowListUpdatedIterator is returned from FilterAllowListUpdated and is used to iterate over the raw logs and unpacked data for AllowListUpdated events raised by the Opensea contract.
type OpenseaAllowListUpdatedIterator struct {
	Event *OpenseaAllowListUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaAllowListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaAllowListUpdated)
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
		it.Event = new(OpenseaAllowListUpdated)
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
func (it *OpenseaAllowListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaAllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaAllowListUpdated represents a AllowListUpdated event raised by the Opensea contract.
type OpenseaAllowListUpdated struct {
	NftContract        common.Address
	PreviousMerkleRoot [32]byte
	NewMerkleRoot      [32]byte
	PublicKeyURI       []string
	AllowListURI       string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAllowListUpdated is a free log retrieval operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Opensea *OpenseaFilterer) FilterAllowListUpdated(opts *bind.FilterOpts, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (*OpenseaAllowListUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var previousMerkleRootRule []interface{}
	for _, previousMerkleRootItem := range previousMerkleRoot {
		previousMerkleRootRule = append(previousMerkleRootRule, previousMerkleRootItem)
	}
	var newMerkleRootRule []interface{}
	for _, newMerkleRootItem := range newMerkleRoot {
		newMerkleRootRule = append(newMerkleRootRule, newMerkleRootItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaAllowListUpdatedIterator{contract: _Opensea.contract, event: "AllowListUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListUpdated is a free log subscription operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Opensea *OpenseaFilterer) WatchAllowListUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaAllowListUpdated, nftContract []common.Address, previousMerkleRoot [][32]byte, newMerkleRoot [][32]byte) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var previousMerkleRootRule []interface{}
	for _, previousMerkleRootItem := range previousMerkleRoot {
		previousMerkleRootRule = append(previousMerkleRootRule, previousMerkleRootItem)
	}
	var newMerkleRootRule []interface{}
	for _, newMerkleRootItem := range newMerkleRoot {
		newMerkleRootRule = append(newMerkleRootRule, newMerkleRootItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "AllowListUpdated", nftContractRule, previousMerkleRootRule, newMerkleRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaAllowListUpdated)
				if err := _Opensea.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
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

// ParseAllowListUpdated is a log parse operation binding the contract event 0xefcd7e019bc8b47d27881fd59e2619280ca5894f285950f10ab049870652efa5.
//
// Solidity: event AllowListUpdated(address indexed nftContract, bytes32 indexed previousMerkleRoot, bytes32 indexed newMerkleRoot, string[] publicKeyURI, string allowListURI)
func (_Opensea *OpenseaFilterer) ParseAllowListUpdated(log types.Log) (*OpenseaAllowListUpdated, error) {
	event := new(OpenseaAllowListUpdated)
	if err := _Opensea.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaAllowedFeeRecipientUpdatedIterator is returned from FilterAllowedFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for AllowedFeeRecipientUpdated events raised by the Opensea contract.
type OpenseaAllowedFeeRecipientUpdatedIterator struct {
	Event *OpenseaAllowedFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaAllowedFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaAllowedFeeRecipientUpdated)
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
		it.Event = new(OpenseaAllowedFeeRecipientUpdated)
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
func (it *OpenseaAllowedFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaAllowedFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaAllowedFeeRecipientUpdated represents a AllowedFeeRecipientUpdated event raised by the Opensea contract.
type OpenseaAllowedFeeRecipientUpdated struct {
	NftContract  common.Address
	FeeRecipient common.Address
	Allowed      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAllowedFeeRecipientUpdated is a free log retrieval operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Opensea *OpenseaFilterer) FilterAllowedFeeRecipientUpdated(opts *bind.FilterOpts, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (*OpenseaAllowedFeeRecipientUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaAllowedFeeRecipientUpdatedIterator{contract: _Opensea.contract, event: "AllowedFeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedFeeRecipientUpdated is a free log subscription operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Opensea *OpenseaFilterer) WatchAllowedFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaAllowedFeeRecipientUpdated, nftContract []common.Address, feeRecipient []common.Address, allowed []bool) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "AllowedFeeRecipientUpdated", nftContractRule, feeRecipientRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaAllowedFeeRecipientUpdated)
				if err := _Opensea.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
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

// ParseAllowedFeeRecipientUpdated is a log parse operation binding the contract event 0x6486c31f9d664e241acf94ec2541d328f6b9e97257ae16a1d887f296f879719f.
//
// Solidity: event AllowedFeeRecipientUpdated(address indexed nftContract, address indexed feeRecipient, bool indexed allowed)
func (_Opensea *OpenseaFilterer) ParseAllowedFeeRecipientUpdated(log types.Log) (*OpenseaAllowedFeeRecipientUpdated, error) {
	event := new(OpenseaAllowedFeeRecipientUpdated)
	if err := _Opensea.contract.UnpackLog(event, "AllowedFeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaCreatorPayoutAddressUpdatedIterator is returned from FilterCreatorPayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for CreatorPayoutAddressUpdated events raised by the Opensea contract.
type OpenseaCreatorPayoutAddressUpdatedIterator struct {
	Event *OpenseaCreatorPayoutAddressUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaCreatorPayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaCreatorPayoutAddressUpdated)
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
		it.Event = new(OpenseaCreatorPayoutAddressUpdated)
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
func (it *OpenseaCreatorPayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaCreatorPayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaCreatorPayoutAddressUpdated represents a CreatorPayoutAddressUpdated event raised by the Opensea contract.
type OpenseaCreatorPayoutAddressUpdated struct {
	NftContract      common.Address
	NewPayoutAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreatorPayoutAddressUpdated is a free log retrieval operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Opensea *OpenseaFilterer) FilterCreatorPayoutAddressUpdated(opts *bind.FilterOpts, nftContract []common.Address, newPayoutAddress []common.Address) (*OpenseaCreatorPayoutAddressUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaCreatorPayoutAddressUpdatedIterator{contract: _Opensea.contract, event: "CreatorPayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCreatorPayoutAddressUpdated is a free log subscription operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Opensea *OpenseaFilterer) WatchCreatorPayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaCreatorPayoutAddressUpdated, nftContract []common.Address, newPayoutAddress []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var newPayoutAddressRule []interface{}
	for _, newPayoutAddressItem := range newPayoutAddress {
		newPayoutAddressRule = append(newPayoutAddressRule, newPayoutAddressItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "CreatorPayoutAddressUpdated", nftContractRule, newPayoutAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaCreatorPayoutAddressUpdated)
				if err := _Opensea.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
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

// ParseCreatorPayoutAddressUpdated is a log parse operation binding the contract event 0x0c69f21751e800ea5960436c9a94370c7adbf54c733a20a025293fbbe8f16252.
//
// Solidity: event CreatorPayoutAddressUpdated(address indexed nftContract, address indexed newPayoutAddress)
func (_Opensea *OpenseaFilterer) ParseCreatorPayoutAddressUpdated(log types.Log) (*OpenseaCreatorPayoutAddressUpdated, error) {
	event := new(OpenseaCreatorPayoutAddressUpdated)
	if err := _Opensea.contract.UnpackLog(event, "CreatorPayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaDropURIUpdatedIterator is returned from FilterDropURIUpdated and is used to iterate over the raw logs and unpacked data for DropURIUpdated events raised by the Opensea contract.
type OpenseaDropURIUpdatedIterator struct {
	Event *OpenseaDropURIUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaDropURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaDropURIUpdated)
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
		it.Event = new(OpenseaDropURIUpdated)
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
func (it *OpenseaDropURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaDropURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaDropURIUpdated represents a DropURIUpdated event raised by the Opensea contract.
type OpenseaDropURIUpdated struct {
	NftContract common.Address
	NewDropURI  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDropURIUpdated is a free log retrieval operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Opensea *OpenseaFilterer) FilterDropURIUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*OpenseaDropURIUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaDropURIUpdatedIterator{contract: _Opensea.contract, event: "DropURIUpdated", logs: logs, sub: sub}, nil
}

// WatchDropURIUpdated is a free log subscription operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Opensea *OpenseaFilterer) WatchDropURIUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaDropURIUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "DropURIUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaDropURIUpdated)
				if err := _Opensea.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
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

// ParseDropURIUpdated is a log parse operation binding the contract event 0xa0295608d25b3033c2e2c41cbac8746c2d08767bcfde6d47fae1ed7ba1d32150.
//
// Solidity: event DropURIUpdated(address indexed nftContract, string newDropURI)
func (_Opensea *OpenseaFilterer) ParseDropURIUpdated(log types.Log) (*OpenseaDropURIUpdated, error) {
	event := new(OpenseaDropURIUpdated)
	if err := _Opensea.contract.UnpackLog(event, "DropURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaPayerUpdatedIterator is returned from FilterPayerUpdated and is used to iterate over the raw logs and unpacked data for PayerUpdated events raised by the Opensea contract.
type OpenseaPayerUpdatedIterator struct {
	Event *OpenseaPayerUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaPayerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaPayerUpdated)
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
		it.Event = new(OpenseaPayerUpdated)
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
func (it *OpenseaPayerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaPayerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaPayerUpdated represents a PayerUpdated event raised by the Opensea contract.
type OpenseaPayerUpdated struct {
	NftContract common.Address
	Payer       common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayerUpdated is a free log retrieval operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Opensea *OpenseaFilterer) FilterPayerUpdated(opts *bind.FilterOpts, nftContract []common.Address, payer []common.Address, allowed []bool) (*OpenseaPayerUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaPayerUpdatedIterator{contract: _Opensea.contract, event: "PayerUpdated", logs: logs, sub: sub}, nil
}

// WatchPayerUpdated is a free log subscription operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Opensea *OpenseaFilterer) WatchPayerUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaPayerUpdated, nftContract []common.Address, payer []common.Address, allowed []bool) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var allowedRule []interface{}
	for _, allowedItem := range allowed {
		allowedRule = append(allowedRule, allowedItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "PayerUpdated", nftContractRule, payerRule, allowedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaPayerUpdated)
				if err := _Opensea.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
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

// ParsePayerUpdated is a log parse operation binding the contract event 0x55a5cfa4bc68ffb9d833b75bf93f6d9c9aadc558dbfa587a9b5bb0ea7d5c38a3.
//
// Solidity: event PayerUpdated(address indexed nftContract, address indexed payer, bool indexed allowed)
func (_Opensea *OpenseaFilterer) ParsePayerUpdated(log types.Log) (*OpenseaPayerUpdated, error) {
	event := new(OpenseaPayerUpdated)
	if err := _Opensea.contract.UnpackLog(event, "PayerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaPublicDropUpdatedIterator is returned from FilterPublicDropUpdated and is used to iterate over the raw logs and unpacked data for PublicDropUpdated events raised by the Opensea contract.
type OpenseaPublicDropUpdatedIterator struct {
	Event *OpenseaPublicDropUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaPublicDropUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaPublicDropUpdated)
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
		it.Event = new(OpenseaPublicDropUpdated)
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
func (it *OpenseaPublicDropUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaPublicDropUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaPublicDropUpdated represents a PublicDropUpdated event raised by the Opensea contract.
type OpenseaPublicDropUpdated struct {
	NftContract common.Address
	PublicDrop  PublicDrop
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPublicDropUpdated is a free log retrieval operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Opensea *OpenseaFilterer) FilterPublicDropUpdated(opts *bind.FilterOpts, nftContract []common.Address) (*OpenseaPublicDropUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaPublicDropUpdatedIterator{contract: _Opensea.contract, event: "PublicDropUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicDropUpdated is a free log subscription operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Opensea *OpenseaFilterer) WatchPublicDropUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaPublicDropUpdated, nftContract []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "PublicDropUpdated", nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaPublicDropUpdated)
				if err := _Opensea.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
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

// ParsePublicDropUpdated is a log parse operation binding the contract event 0x3e30d8e1f739ea4795c481b21c23f905e938b80339305f3508e43c558e5dead3.
//
// Solidity: event PublicDropUpdated(address indexed nftContract, (uint80,uint48,uint48,uint16,uint16,bool) publicDrop)
func (_Opensea *OpenseaFilterer) ParsePublicDropUpdated(log types.Log) (*OpenseaPublicDropUpdated, error) {
	event := new(OpenseaPublicDropUpdated)
	if err := _Opensea.contract.UnpackLog(event, "PublicDropUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaSeaDropMintIterator is returned from FilterSeaDropMint and is used to iterate over the raw logs and unpacked data for SeaDropMint events raised by the Opensea contract.
type OpenseaSeaDropMintIterator struct {
	Event *OpenseaSeaDropMint // Event containing the contract specifics and raw log

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
func (it *OpenseaSeaDropMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaSeaDropMint)
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
		it.Event = new(OpenseaSeaDropMint)
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
func (it *OpenseaSeaDropMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaSeaDropMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaSeaDropMint represents a SeaDropMint event raised by the Opensea contract.
type OpenseaSeaDropMint struct {
	NftContract    common.Address
	Minter         common.Address
	FeeRecipient   common.Address
	Payer          common.Address
	QuantityMinted *big.Int
	UnitMintPrice  *big.Int
	FeeBps         *big.Int
	DropStageIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSeaDropMint is a free log retrieval operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Opensea *OpenseaFilterer) FilterSeaDropMint(opts *bind.FilterOpts, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (*OpenseaSeaDropMintIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaSeaDropMintIterator{contract: _Opensea.contract, event: "SeaDropMint", logs: logs, sub: sub}, nil
}

// WatchSeaDropMint is a free log subscription operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Opensea *OpenseaFilterer) WatchSeaDropMint(opts *bind.WatchOpts, sink chan<- *OpenseaSeaDropMint, nftContract []common.Address, minter []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "SeaDropMint", nftContractRule, minterRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaSeaDropMint)
				if err := _Opensea.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
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

// ParseSeaDropMint is a log parse operation binding the contract event 0xe90cf9cc0a552cf52ea6ff74ece0f1c8ae8cc9ad630d3181f55ac43ca076b7d6.
//
// Solidity: event SeaDropMint(address indexed nftContract, address indexed minter, address indexed feeRecipient, address payer, uint256 quantityMinted, uint256 unitMintPrice, uint256 feeBps, uint256 dropStageIndex)
func (_Opensea *OpenseaFilterer) ParseSeaDropMint(log types.Log) (*OpenseaSeaDropMint, error) {
	event := new(OpenseaSeaDropMint)
	if err := _Opensea.contract.UnpackLog(event, "SeaDropMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaSignedMintValidationParamsUpdatedIterator is returned from FilterSignedMintValidationParamsUpdated and is used to iterate over the raw logs and unpacked data for SignedMintValidationParamsUpdated events raised by the Opensea contract.
type OpenseaSignedMintValidationParamsUpdatedIterator struct {
	Event *OpenseaSignedMintValidationParamsUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaSignedMintValidationParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaSignedMintValidationParamsUpdated)
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
		it.Event = new(OpenseaSignedMintValidationParamsUpdated)
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
func (it *OpenseaSignedMintValidationParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaSignedMintValidationParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaSignedMintValidationParamsUpdated represents a SignedMintValidationParamsUpdated event raised by the Opensea contract.
type OpenseaSignedMintValidationParamsUpdated struct {
	NftContract                common.Address
	Signer                     common.Address
	SignedMintValidationParams SignedMintValidationParams
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterSignedMintValidationParamsUpdated is a free log retrieval operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Opensea *OpenseaFilterer) FilterSignedMintValidationParamsUpdated(opts *bind.FilterOpts, nftContract []common.Address, signer []common.Address) (*OpenseaSignedMintValidationParamsUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaSignedMintValidationParamsUpdatedIterator{contract: _Opensea.contract, event: "SignedMintValidationParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchSignedMintValidationParamsUpdated is a free log subscription operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Opensea *OpenseaFilterer) WatchSignedMintValidationParamsUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaSignedMintValidationParamsUpdated, nftContract []common.Address, signer []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "SignedMintValidationParamsUpdated", nftContractRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaSignedMintValidationParamsUpdated)
				if err := _Opensea.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
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

// ParseSignedMintValidationParamsUpdated is a log parse operation binding the contract event 0xcaeb4009c05208df426d15ff50b608287b05d21dee1f790552ea451a540a7be0.
//
// Solidity: event SignedMintValidationParamsUpdated(address indexed nftContract, address indexed signer, (uint80,uint24,uint40,uint40,uint40,uint16,uint16) signedMintValidationParams)
func (_Opensea *OpenseaFilterer) ParseSignedMintValidationParamsUpdated(log types.Log) (*OpenseaSignedMintValidationParamsUpdated, error) {
	event := new(OpenseaSignedMintValidationParamsUpdated)
	if err := _Opensea.contract.UnpackLog(event, "SignedMintValidationParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpenseaTokenGatedDropStageUpdatedIterator is returned from FilterTokenGatedDropStageUpdated and is used to iterate over the raw logs and unpacked data for TokenGatedDropStageUpdated events raised by the Opensea contract.
type OpenseaTokenGatedDropStageUpdatedIterator struct {
	Event *OpenseaTokenGatedDropStageUpdated // Event containing the contract specifics and raw log

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
func (it *OpenseaTokenGatedDropStageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpenseaTokenGatedDropStageUpdated)
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
		it.Event = new(OpenseaTokenGatedDropStageUpdated)
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
func (it *OpenseaTokenGatedDropStageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpenseaTokenGatedDropStageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpenseaTokenGatedDropStageUpdated represents a TokenGatedDropStageUpdated event raised by the Opensea contract.
type OpenseaTokenGatedDropStageUpdated struct {
	NftContract     common.Address
	AllowedNftToken common.Address
	DropStage       TokenGatedDropStage
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenGatedDropStageUpdated is a free log retrieval operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Opensea *OpenseaFilterer) FilterTokenGatedDropStageUpdated(opts *bind.FilterOpts, nftContract []common.Address, allowedNftToken []common.Address) (*OpenseaTokenGatedDropStageUpdatedIterator, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Opensea.contract.FilterLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return &OpenseaTokenGatedDropStageUpdatedIterator{contract: _Opensea.contract, event: "TokenGatedDropStageUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenGatedDropStageUpdated is a free log subscription operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Opensea *OpenseaFilterer) WatchTokenGatedDropStageUpdated(opts *bind.WatchOpts, sink chan<- *OpenseaTokenGatedDropStageUpdated, nftContract []common.Address, allowedNftToken []common.Address) (event.Subscription, error) {

	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var allowedNftTokenRule []interface{}
	for _, allowedNftTokenItem := range allowedNftToken {
		allowedNftTokenRule = append(allowedNftTokenRule, allowedNftTokenItem)
	}

	logs, sub, err := _Opensea.contract.WatchLogs(opts, "TokenGatedDropStageUpdated", nftContractRule, allowedNftTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpenseaTokenGatedDropStageUpdated)
				if err := _Opensea.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
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

// ParseTokenGatedDropStageUpdated is a log parse operation binding the contract event 0xc695f93ae16034280e4fc93181b6afca9af23027ac1f1842a2287ba25cdc4476.
//
// Solidity: event TokenGatedDropStageUpdated(address indexed nftContract, address indexed allowedNftToken, (uint80,uint16,uint48,uint48,uint8,uint32,uint16,bool) dropStage)
func (_Opensea *OpenseaFilterer) ParseTokenGatedDropStageUpdated(log types.Log) (*OpenseaTokenGatedDropStageUpdated, error) {
	event := new(OpenseaTokenGatedDropStageUpdated)
	if err := _Opensea.contract.UnpackLog(event, "TokenGatedDropStageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
