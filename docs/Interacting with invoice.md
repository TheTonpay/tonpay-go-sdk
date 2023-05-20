# Interacting with invoice

When working with the store, you had a chance to get the invoice. Here you'll learn what you can do with it.

## Edit an invoice
To edit an invoice, you can follow the same formula as before:
1. Connect the owner's wallet using the seed phrases:
```
words := strings.Split("seed phrase ...", " ")
// Wallet version you can find in explorer, for example on tonscan: https://tonscan.org/address/<wallet_address>
w, err := wallet.FromSeed(api, words, wallet.V4R2)
if err != nil {
	panic(err)
}
```
2. Build the message cell:
```  
editInvoiceMessage, err := tonpaygo.EditInvoiceMessage(false, "", "testid", "", 700000000)
if err != nil {
	panic(err)
}
```
3. Send an internal message with the cell (ensure that the fee is `EditInvoiceFee`):
```
message := wallet.Message{
	Mode: 1,
	InternalMessage: &tlb.InternalMessage{
	Bounce: true,
	DstAddr: addr,
	Amount: tlb.FromNanoTONU(tonpaygo.EditInvoiceFee),
	Body: editInvoiceMessage,
	},
}

err = w.Send(context.Background(), &message, true)
if err != nil {
	panic(err)
}
```
**Warning:** Only the merchant can edit the invoice. It's important to note that an invoice cannot be edited once it has been paid or if it is inactive.

## Manage active status
To manage the active status of an invoice, you can use the following steps:
```
activateInvoiceMessage, err := tonpaygo.ActivateInvoiceMessage()
// or
deactivateInvoiceMessage, err := tonpaygo.DeactivateInvoiceMessage()

// Then do step 3(send message) from previous example
```
After confirming the transaction, the invoice will be activated/deactivated.

**Warning:** Only the merchant can activate or deactivate the invoice. It's important to note that when the invoice is deactivated, it will not accept any payments and cannot be edited.

## Paying the invoice
After the invoice is issued, the customer can make the payment in two ways:
1. Provide the customer with the payment link. You can find an example of this method in the [create-and-pay-invoice-url](https://github.com/TheTonpay/tonpay-go-sdk/tree/main/examples/create-and-pay-invoice-url) example.
2. Send a payment to the invoice programmatically. To do this, you can use the following steps:

a. Get the payment message using `tonpaygo.PayInvoiceMessage()`.
```
message := wallet.Message{
	...
	InternalMessage: &tlb.InternalMessage{
		...
		Amount: invoiceAmount + tlb.FromNanoTONU(tonpaygo.ActiveInvoiceFee),
	},
}
```
b. Send the payment message using the same process as mentioned before.

## Get invoice info
Like in example below you can get other contract data:
```
func GetStore(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error)
func GetMerchant(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error)
func GetCustomer(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error)
func HasCustomer(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error)
func GetInvoiceID(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error)
func GetInvoiceMetadata(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error)
func GetAmount(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (uint64, error)
func IsPaid(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error)
func IsActive(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error)
func GetInvoiceVersion(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (uint64, error)

```
Or get all in one:
```
func  GetInvoiceData(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (InvoiceData, error)
```
