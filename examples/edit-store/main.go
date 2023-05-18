package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	tonpaygo "github.com/TheTonpay/tonpay-go-sdk/src"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

func EditStore(addr *address.Address, version wallet.Version, seedWords []string) bool {
	client := liteclient.NewConnectionPool()

	// Choose testnet or mainnet
	configUrl := "https://ton-blockchain.github.io/global.config.json" // mainnet
	// configUrl := "https://ton-blockchain.github.io/testnet-global.config.json" // testnet
	err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
	if err != nil {
		log.Println(err)
		return false
	}

	api := ton.NewAPIClient(client)

	editStoreCell, err := tonpaygo.EditStoreMessage("Durger King 3.1", "Fast food", "7E5D1FDEA9EEC531F3E6BAA69E22615FB4CEA111BAE40219AB4E60053489DD64", "", 541)
	if err != nil {
		fmt.Println(err)
		return false
	}

	w, err := wallet.FromSeed(api, seedWords, wallet.V4R2)
	if err != nil {
		fmt.Println(err)
		return false
	}

	message := wallet.Message{
		Mode: 1,
		InternalMessage: &tlb.InternalMessage{
			Bounce:  true,
			DstAddr: addr,
			Amount:  tlb.FromNanoTONU(tonpaygo.EditStoreFee),
			Body:    editStoreCell,
		},
	}

	err = w.Send(context.Background(), &message, true)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func main() {
	storeAddr := address.MustParseAddr("EQBOPmMoxApwWAUW77ou70qOqqlDjCZLdqbdo6xRm3tUxP_7")
	walletVersion := wallet.V4R2
	seedWords := strings.Split("seed words here ...", " ")
	isUpdated := EditStore(storeAddr, walletVersion, seedWords)
	fmt.Println("isUpdated:", isUpdated)
}
