package arbitrum

import (
	"fmt"
	"log"
	"math/big"
)

func (kc *KingdomlyClient) Mint(amount int) {
	// Get the fees quote
	quote, err := kc.QuoteFees(amount)
	if err != nil {
		return
	}
	kc.TransactionSigner.Value = quote

	// Mint NFT
	tx, err := kc.Contract.BatchMint(kc.TransactionSigner, big.NewInt(int64(amount)), big.NewInt(0))
	if err != nil {
		fmt.Println("Failed to batch mint NFTs: ", err)
		return
	}

	log.Printf("Batch Mint transaction sent: %s\n", tx.Hash().Hex())
}
