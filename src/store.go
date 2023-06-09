package tonpaygo

import (
	"context"
	"fmt"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

// OpCodes
const (
	IssueInvoice       = 0x4b4e70b0
	RequestPurchase    = 0x36b795b5
	EditStore          = 0xa0b2b61d
	DeleteStore        = 0xfb4aca1a // UNUSED
	DeactivateStore    = 0xf9bf9637
	ActivateStore      = 0x97500daf
	UpgradeCodeFull    = 0xb43bbb52
	UpgradeCodeStore   = 0xacb08f28
	UpgradeCodeInvoice = 0xb5f1424f
)

var (
	DeployStoreFee     = tlb.MustFromTON("0.005").NanoTON().Uint64()
	EditStoreFee       = tlb.MustFromTON("0.005").NanoTON().Uint64()
	ActivateStoreFee   = tlb.MustFromTON("0.005").NanoTON().Uint64()
	DeactivateStoreFee = tlb.MustFromTON("0.005").NanoTON().Uint64()
	IssueInvoiceFee    = tlb.MustFromTON("0.02").NanoTON().Uint64()
	RequestPurchaseFee = tlb.MustFromTON("0.05").NanoTON().Uint64()
	FullUpgradeFee     = tlb.MustFromTON("0.006").NanoTON().Uint64()
	InvoiceUpgradeFee  = tlb.MustFromTON("0.006").NanoTON().Uint64()
)

type StoreData struct {
	Owner       *address.Address
	Name        string
	Description string
	Image       string
	Webhook     string
	MccCode     uint64
	IsActive    bool
	InvoiceCode *cell.Cell
	Version     int64
}

func StoreDataToCell(storeConfig StoreData) (*cell.Cell, error) {
	nameCell, err := wallet.CreateCommentCell(storeConfig.Name)
	if err != nil {
		return nil, err
	}
	descriptionCell, err := wallet.CreateCommentCell(storeConfig.Description)
	if err != nil {
		return nil, err
	}
	imageCell, err := wallet.CreateCommentCell(storeConfig.Image)
	if err != nil {
		return nil, err
	}
	webhookCell, err := wallet.CreateCommentCell(storeConfig.Webhook)
	if err != nil {
		return nil, err
	}

	var IsActive int64 = -1
	if !storeConfig.IsActive {
		IsActive = 0
	}

	return cell.BeginCell().
		MustStoreAddr(storeConfig.Owner).
		MustStoreRef(nameCell).
		MustStoreRef(descriptionCell).
		MustStoreRef(cell.BeginCell().
			MustStoreRef(imageCell).
			MustStoreRef(webhookCell).
			EndCell()).
		MustStoreUInt(storeConfig.MccCode, 16).
		MustStoreInt(IsActive, 2).
		MustStoreRef(storeConfig.InvoiceCode).
		EndCell(), nil
}

func IssueInvoiceMessage(customer *address.Address, invoiceHasCustomer bool, invoiceID string, metadata string, amount uint64) (*cell.Cell, error) {
	invoiceIDCell, err := wallet.CreateCommentCell(invoiceID)
	if err != nil {
		return nil, err
	}

	metadataCell, err := wallet.CreateCommentCell(metadata)
	if err != nil {
		return nil, err
	}

	var hasCustomer int64 = 0
	if invoiceHasCustomer {
		hasCustomer = -1
		if customer != nil || customer.IsAddrNone() {
			return nil, fmt.Errorf("customer address is not set")
		}
	}

	customerAddressCell := cell.BeginCell().MustStoreAddr(customer).EndCell()
	return cell.BeginCell().
		MustStoreUInt(IssueInvoice, 32).
		MustStoreUInt(0, 64).
		MustStoreInt(hasCustomer, 2).
		MustStoreRef(customerAddressCell).
		MustStoreRef(invoiceIDCell).
		MustStoreRef(metadataCell).
		MustStoreUInt(amount, 64).
		EndCell(), nil
}

func RequestPurchaseMessage(invoiceID string, metadata string, amount uint64) (*cell.Cell, error) {
	invoiceIDCell, err := wallet.CreateCommentCell(invoiceID)
	if err != nil {
		return nil, err
	}

	metadataCell, err := wallet.CreateCommentCell(metadata)
	if err != nil {
		return nil, err
	}

	return cell.BeginCell().
		MustStoreUInt(RequestPurchase, 32).
		MustStoreUInt(0, 64).
		MustStoreRef(invoiceIDCell).
		MustStoreRef(metadataCell).
		MustStoreUInt(amount, 64).
		EndCell(), nil
}

func EditStoreMessage(storeName string, storeDescription string, storeImage string, storeWebhook string, mccCode uint64) (*cell.Cell, error) {
	name, err := wallet.CreateCommentCell(storeName)
	if err != nil {
		return nil, err
	}
	description, err := wallet.CreateCommentCell(storeDescription)
	if err != nil {
		return nil, err
	}
	image, err := wallet.CreateCommentCell(storeImage)
	if err != nil {
		return nil, err
	}
	webhook, err := wallet.CreateCommentCell(storeWebhook)
	if err != nil {
		return nil, err
	}

	return cell.BeginCell().
		MustStoreUInt(EditStore, 32).
		MustStoreUInt(0, 64).
		MustStoreRef(name).
		MustStoreRef(description).
		MustStoreRef(image).
		MustStoreRef(webhook).
		MustStoreUInt(mccCode, 16).
		EndCell(), nil
}

func DeactivateStoreMessage() (*cell.Cell, error) {
	return cell.BeginCell().
		MustStoreUInt(DeactivateStore, 32).
		MustStoreUInt(0, 64).
		EndCell(), nil
}

func ActivateStoreMessage() (*cell.Cell, error) {
	return cell.BeginCell().
		MustStoreUInt(ActivateStore, 32).
		MustStoreUInt(0, 64).
		EndCell(), nil
}

func UpgradeCodeFullMessage(storeCode *cell.Cell, invoiceCode *cell.Cell, hasNewData bool, newData *cell.Cell) (*cell.Cell, error) {
	if hasNewData && newData == nil {
		return nil, fmt.Errorf("newData is not set")
	}

	var hasNewDataFlag int64 = 0
	if hasNewData {
		hasNewDataFlag = -1
	}

	return cell.BeginCell().
		MustStoreUInt(UpgradeCodeFull, 32).
		MustStoreUInt(0, 64).
		MustStoreRef(storeCode).
		MustStoreRef(invoiceCode).
		MustStoreInt(hasNewDataFlag, 2).
		MustStoreRef(newData).
		EndCell(), nil
}

func UpgradeCodeStoreMessage(storeCode *cell.Cell, hasNewData bool, newData *cell.Cell) (*cell.Cell, error) {
	if hasNewData && newData == nil {
		return nil, fmt.Errorf("newData is not set")
	}

	var hasNewDataFlag int64 = 0
	if hasNewData {
		hasNewDataFlag = -1
	}

	return cell.BeginCell().
		MustStoreUInt(UpgradeCodeStore, 32).
		MustStoreUInt(0, 64).
		MustStoreRef(storeCode).
		MustStoreInt(hasNewDataFlag, 2).
		MustStoreRef(newData).
		EndCell(), nil
}

func UpgradeCodeInvoiceMessage(invoiceCode *cell.Cell) (*cell.Cell, error) {
	return cell.BeginCell().
		MustStoreUInt(UpgradeCodeInvoice, 32).
		MustStoreUInt(0, 64).
		MustStoreRef(invoiceCode).
		EndCell(), nil
}

func GetOwner(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_owner")
	if err != nil {
		return nil, err
	}

	owner := res.MustSlice(0).MustLoadAddr()
	return owner, nil
}

func GetName(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_name")
	if err != nil {
		return "", err
	}

	name := res.MustCell(0).BeginParse().MustLoadStringSnake()
	return name, nil
}

func GetDescription(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_description")
	if err != nil {
		return "", err
	}

	description := res.MustCell(0).BeginParse().MustLoadStringSnake()
	return description, nil
}

func GetImage(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_image")
	if err != nil {
		return "", err
	}

	image := res.MustCell(0).BeginParse().MustLoadStringSnake()
	return image, nil
}

func GetStoreWebhook(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_webhook")
	if err != nil {
		return "", err
	}

	webhook := res.MustCell(0).BeginParse().MustLoadStringSnake()
	return webhook, nil
}

func GetMccCode(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (int64, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_mcc_code")
	if err != nil {
		return -1, err
	}

	mccCode := res.MustInt(0).Int64()
	return mccCode, nil
}

func GetIsActive(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_active")
	if err != nil {
		return false, err
	}

	isActive := res.MustInt(0).Int64()
	return isActive == -1, nil
}

func GetStoreVersion(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (uint64, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_version")
	if err != nil {
		return 0, err
	}

	version := res.MustInt(0).Uint64()
	return version, nil
}

func GetStoreData(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (StoreData, error) {
	res, err := api.RunGetMethod(context.Background(), block, addr, "get_store_data")
	if err != nil {
		return StoreData{}, err
	}

	owner := res.MustSlice(0).MustLoadAddr()
	name := res.MustCell(1).BeginParse().MustLoadStringSnake()
	description := res.MustCell(2).BeginParse().MustLoadStringSnake()
	image := res.MustCell(3).BeginParse().MustLoadStringSnake()
	webhook := res.MustCell(4).BeginParse().MustLoadStringSnake()
	mccCode := res.MustInt(5)
	isActive := res.MustInt(6)
	invoiceCode := res.MustCell(7)
	version := res.MustInt(8)

	return StoreData{
		Owner:       owner,
		Name:        name[4:],
		Description: description[4:],
		Image:       image[4:],
		Webhook:     webhook[4:],
		MccCode:     mccCode.Uint64(),
		IsActive:    isActive.Int64() == -1,
		InvoiceCode: invoiceCode,
		Version:     version.Int64(),
	}, nil
}
