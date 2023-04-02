package main

import (
	"fmt"
	tonpaygo "tonpay-go-sdk/src"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
)

func main() {
	storeAddr := address.MustParseAddr("EQCA_bhKoqvxxG44_fBDC4VFQ36Sdzu6i8KaD2SMaZwWS_N9")
	amount := tlb.MustFromTON("1").NanoTON().Uint64()

	requestCell, err := tonpaygo.RequestPurchaseMessage("created using tonpaygo, paid by wallet", amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := tonpaygo.BuildURLForWallet(storeAddr.String(), requestCell, amount, tonpaygo.RequestPurchaseFee)

	fmt.Printf("Send this url to the customer: %s", url)
}
