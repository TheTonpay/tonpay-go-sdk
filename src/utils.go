package tonpaygo

import (
	"encoding/base64"
	"fmt"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func BuildURLForWallet(reciver string, cell_ *cell.Cell, amount uint64, comission uint64) string {
	bin := base64.StdEncoding.EncodeToString(cell_.ToBOCWithFlags(false))
	total := amount + comission

	return fmt.Sprintf("ton://transfer/%s?bin=%s&amount=%d", reciver, bin, total)
}

func PrecalculateInvoiceAddress(workchain int, config InvoiceData) (string, error) {
	invoiceCell, err := InvoiceConfigToCell(config)
	if err != nil {
		return "", err
	}

	codeCellBytes, err := base64.StdEncoding.DecodeString(InvoiceCode)
	if err != nil {
		panic(err)
	}
	contractCodeCell, err := cell.FromBOC(codeCellBytes)
	if err != nil {
		panic(err)
	}

	return PrecalculateContractAddress(workchain, contractCodeCell, invoiceCell), nil
}

func PrecalculateContractAddress(workchain int, code *cell.Cell, data *cell.Cell) string {
	contractCell, err := tlb.ToCell(tlb.StateInit{
		Code: code,
		Data: data,
	})

	if err != nil {
		panic(err)
	}

	return address.NewAddress(0, byte(workchain), contractCell.Hash()).String()
}
