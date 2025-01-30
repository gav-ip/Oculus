package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	abstract "github.com/scombs97/vision-bot/backend/modules/abstract/magiceden"
)

const (
	arb_test_contract  = "0x3CBa380D9722028e590a2132324A5859Fb2AAeC0"
	eth_test_contract  = "0x0d5D728159E18ea10D8CA90D29b69A5a286CcdAB"
	base_test_contract = "0x80958DC45286f460eCbd0x4691281dDBBd2C017edcB3C41ae7acFB689b30fe174FD74e832Dd13AFED6"
	abstract_contract  = "0x8DB93eE05B9eEfe2c1A151080201922A7b51dB14"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	privateKeyHex := os.Getenv("PRIVATE_KEY")

	client, err := abstract.NewMagicEdenClient(privateKeyHex, abstract_contract)
	if err != nil {
		log.Fatalf("Error making client")
		return
	}

	// client.Mint(1)

	info, err := client.GetStageInfo()
	if err != nil {
		log.Fatalf("Error getting info")
		return
	}
	fmt.Println(info.MintFee)
	fmt.Println(info.Price)
}
