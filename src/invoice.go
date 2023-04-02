package tonpaygo

import (
	"context"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
)

var (
	DeployInvoiceFee   = tlb.MustFromTON("0.005").NanoTON().Uint64()
	EditInvoiceFee     = tlb.MustFromTON("0.005").NanoTON().Uint64()
	ActiveInvoiceFee   = tlb.MustFromTON("0.005").NanoTON().Uint64()
	DeactiveInvoiceFee = tlb.MustFromTON("0.005").NanoTON().Uint64()
)

type InvoiceData struct {
	Store       string
	Merchant    string
	Beneficiary string
	HasCustomer bool
	Customer    string
	InvoiceID   string
	Amount      uint64
	Paid        bool
	Active      bool
	Version     int64
}

func GetStore(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_store")
	if err != nil {
		return nil, err
	}

	owner := res.MustSlice(0).MustLoadAddr()
	return owner, nil
}

func GetMerchant(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_merchant")
	if err != nil {
		return nil, err
	}

	owner := res.MustSlice(0).MustLoadAddr()
	return owner, nil
}

func GetCustomer(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_customer")
	if err != nil {
		return nil, err
	}

	customer := res.MustSlice(0).MustLoadAddr()
	return customer, nil
}

func HasCustomer(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_has_customer")
	if err != nil {
		return false, err
	}

	hasCustomer := res.MustInt(0).Int64()
	return hasCustomer == -1, nil
}

func GetInvoiceID(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_id")
	if err != nil {
		return "", err
	}

	id := res.MustCell(0).BeginParse().MustLoadStringSnake()
	return id, nil
}

func GetAmount(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (uint64, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_amount")
	if err != nil {
		return 0, err
	}

	amount := res.MustInt(0)
	return amount.Uint64(), nil
}

func IsPaid(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_paid")
	if err != nil {
		return false, err
	}

	isPaid := res.MustInt(0).Int64()
	return isPaid == -1, nil
}

func IsActive(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_active")
	if err != nil {
		return false, err
	}

	isActive := res.MustInt(0).Int64()
	return isActive == -1, nil
}

func GetInvoiceVersion(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (uint64, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_version")
	if err != nil {
		return 0, err
	}

	version := res.MustInt(0)
	return version.Uint64(), nil
}

func GetInvoiceData(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (InvoiceData, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_invoice_data")
	if err != nil {
		return InvoiceData{}, err
	}

	store := res.MustSlice(0).MustLoadAddr()
	merchant := res.MustSlice(1).MustLoadAddr()
	beneficiary := res.MustSlice(2).MustLoadAddr()
	hasCustomer := res.MustInt(3)
	customer := res.MustCell(4).BeginParse().MustLoadAddr()
	invoiceID := res.MustCell(5).BeginParse().MustLoadStringSnake()
	amount := res.MustInt(6)
	paid := res.MustInt(7)
	active := res.MustInt(8)
	version := res.MustInt(9)

	return InvoiceData{
		Store:       store.String(),
		Merchant:    merchant.String(),
		Beneficiary: beneficiary.String(),
		HasCustomer: hasCustomer.Int64() == -1,
		Customer:    customer.String(),
		InvoiceID:   invoiceID,
		Amount:      amount.Uint64(),
		Paid:        paid.Int64() == -1,
		Active:      active.Int64() == -1,
		Version:     version.Int64(),
	}, nil
}
