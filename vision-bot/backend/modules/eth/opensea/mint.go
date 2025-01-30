package eth

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	SEADROP_FEE_ADDR = "0x0000a26b00c1F0DF003000390027140000fAa719"
)

func (oc *OpenseaClient) Mint(amount int) {
	// Get the mint price
	publicDrop, err := oc.GetPublicDrop()
	if err != nil {
		return
	}
	oc.TransactionSigner.Value = publicDrop.MintPrice

	// Mint NFT
	tx, err := oc.Contract.MintPublic(oc.TransactionSigner, oc.ContractAddress, common.HexToAddress(SEADROP_FEE_ADDR), common.HexToAddress("0x0000000000000000000000000000000000000000"), big.NewInt(int64(amount)))
	if err != nil {
		fmt.Println("Failed to mint NFTs: ", err)
		return
	}

	log.Printf("Batch Mint transaction sent: %s\n", tx.Hash().Hex())
}
