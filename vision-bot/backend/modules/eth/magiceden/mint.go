package eth

import (
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	MINT_ABI = `[ 
		{"constant": false, "inputs": [
            {"name": "qty", "type": "uint32"},
            {"name": "proof", "type": "bytes32[]"},
            {"name": "timestamp", "type": "uint64"},
            {"name": "signature", "type": "bytes"}
        ], "name": "mint", "outputs": [], "stateMutability": "payable", "type": "function"}
	]`
)

func (mc *MagicEdenClient) Mint(amount int) {
	// Get the stage info with mint price
	info, err := mc.GetStageInfo()
	if err != nil {
		return
	}
	mc.TransactionSigner.Value = info.Price

	// Get params
	var proof [][32]byte
	timestamp := uint64(time.Now().Unix())

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(MINT_ABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Pack the parameters for the `mint` function
	data, err := parsedABI.Pack("mint", uint32(amount), proof, timestamp, []byte{})
	if err != nil {
		log.Fatalf("Failed to ABI encode parameters: %v", err)
	}

	// Hash the encoded data
	hash := crypto.Keccak256Hash(data)

	// Get signature param
	signature, err := crypto.Sign(hash.Bytes(), mc.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to sign the hash: %v", err)
	}

	tx, err := mc.Contract.Mint(mc.TransactionSigner, uint32(amount), proof, timestamp, signature)
	if err != nil {
		log.Fatalf("Failed to mint NFTs: %v", err)
	}

	log.Printf("Mint transaction sent: %s\n", tx.Hash().Hex())
}
