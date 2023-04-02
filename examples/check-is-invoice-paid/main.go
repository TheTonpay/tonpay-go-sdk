package main

import (
	"context"
	"fmt"
	"log"
	tonpaygo "tonpay-go-sdk/src"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

func CheckIsInvoicePaid(addr *address.Address) bool {
	// First of all, we need to create connection pool
	client := liteclient.NewConnectionPool()

	// Choose testnet or mainnet
	// configUrl := "https://ton-blockchain.github.io/global-config.json"	   // mainnet
	configUrl := "https://ton-blockchain.github.io/testnet-global.config.json" // testnet
	err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
	if err != nil {
		log.Println(err)
		return false
	}

	api := ton.NewAPIClient(client)

	// Get block
	block, err := api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}

	isPaid, err := tonpaygo.IsPaid(api, block, addr)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return isPaid
}

func main() {
	// This one paid
	addr := address.MustParseAddr("EQCPibUxBk7xwwcKBja6SbC5Pm1BCihGDA6SY5wTmFhrmYSh")
	isPaid := CheckIsInvoicePaid(addr)
	fmt.Println(isPaid)

	// This one not
	addr2 := address.MustParseAddr("EQAO4tDMxk1NX6BXJt_LpzHSxfOGwX45M3VfL5Yqu-exS7S6")
	isPaid2 := CheckIsInvoicePaid(addr2)
	fmt.Println(isPaid2)
}
