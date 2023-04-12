package main

import (
	"fmt"

	tonpaygo "github.com/TheTonpay/tonpay-go-sdk/src"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
)

func main() {
	owner := "EQASwvTkaA3TyA-OFyhj2-H7pX5ShRbA0mD3dLRGhayvoOtd"
	store := "EQBeEgLLL369qn-SK8ikE2FzX9MnblI_RDP5WLxCHm6Ihl_o"
	storeAddr := address.MustParseAddr(store)
	amount := tlb.MustFromTON("1").NanoTON().Uint64()
	invoiceID := "created using tonpaygo, paid by wallet"

	// Building request cell. It will be used to generate url for the wallet
	requestCell, err := tonpaygo.RequestPurchaseMessage(invoiceID, "", amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Building invoice config. It will be used to calculate invoice address
	config := tonpaygo.InvoiceData{
		Store:       store,
		Merchant:    owner,
		Beneficiary: "EQD2wz8Rq5QDj9iK2Z_leGQu-Rup__y-Z4wo8Lm7-tSD6Iz2",
		HasCustomer: false,
		Customer:    tonpaygo.ZeroAddress,
		InvoiceID:   invoiceID,
		Metadata:    "",
		Amount:      amount,
		Paid:        false,
		Active:      true,
	}

	// Calculating invoice address
	invoiceAddress, err := tonpaygo.PrecalculateInvoiceAddress(0, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Building url for the wallet
	url := tonpaygo.BuildURLForWallet(storeAddr.String(), requestCell, amount, tonpaygo.RequestPurchaseFee)

	fmt.Println("Send this url to the customer:", url)
	fmt.Println("After payment you can find the invoice at this address:", invoiceAddress)
}
