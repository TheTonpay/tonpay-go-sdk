
# Interacting with Store

Congratulations on initializing the SDK! Now let's move on to interacting with stores.

## Getting a store

To start interacting with a store, you need to obtain the Store instance. To do that, you'll need the store ID, which can be found on your [store page](https://business.thetonpay.app/stores) in the [merchant portal](https://business.thetonpay.app/). Click the "COPY ID" button to obtain the store ID.

![enter image description here](https://tonbyte.com/gateway/3A3420357AA219F7487DEDFB1F031A04CD9392073D8B2D9B54F34BDE06772D60/Screenshot_20230421_184454.png)

Let's assume that our store ID is `EQCtRWrkfxovf-H02xc5CjHpC4t3VnxAIrCg6LJaaYflj4RY`. Add the following code after the initialization:
```
storeAddress := "EQCPibUxBk7xwwcKBja6SbC5Pm1BCihGDA6SY5wTmFhrmYSh"
storeOwner, err := tonpaygo.GetOwner(api, block, storeAddress)
if err != nil {
fmt.Println(err)
return  false
}

fmt.Printf(storeName)
```
Congratulations! You have now retrieved the store owner's address directly from the Blockchain contract.

## Edit store
To edit store data, you need to follow these three steps: 
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
editStoreCell, err := tonpaygo.EditStoreMessage("Durger King 3.0", "Fast food", "7E5D1FDEA9EEC531F3E6BAA69E22615FB4CEA111BAE40219AB4E60053489DD64", "", 541)
if err != nil {
	panic(err)
}
```
3. Send an internal message with the cell:
```
message := wallet.Message{
	Mode: 1,
	InternalMessage: &tlb.InternalMessage{
	Bounce: true,
	DstAddr: addr,
	Amount: tlb.FromNanoTONU(tonpaygo.EditStoreFee),
	Body: editStoreCell,
	},
}

err = w.Send(context.Background(), &message, true)
if err != nil {
	panic(err)
}
```
**Note:** Make sure to replace "seed phrase ..." with the actual seed phrases of the owner's wallet.

## Create invoice
To create an invoice, there are two methods available:

1.  Create a cell that stores all the required information about the payment, predicts the future invoice contract address, and packs it into a URL. You can find an example of this method [here](https://github.com/TheTonpay/tonpay-go-sdk/tree/main/examples/create-and-pay-invoice-url). This method doesn't require the store wallet seeds on the backend.
    
2.  Deploy the invoice contract from the backend and send its address to the customer. First, retrieve the wallet from the previous section, "Edit store":
```
words := strings.Split("seed phrase ...", " ")

// Wallet version you can find in explorer, for example on tonscan: https://tonscan.org/address/<wallet_address>
w, err := wallet.FromSeed(api, words, wallet.V4R2)
if err != nil {
	panic(err)
}
```
Then, build and send the message (ensure that you double-check the fee you send):
```  
createIssueStoreCell, err := tonpaygo.IssueInvoiceMessage(nil, false, "Invoice from go", "", tlb.MustFromTON("1").NanoTON().Uint64())
if err != nil {
	panic(err)
}

message := wallet.Message{
	Mode: 1,
	InternalMessage: &tlb.InternalMessage{
		Bounce: true,
		DstAddr: addr,
		Amount: tlb.FromNanoTONU(tonpaygo.IssueInvoiceFee),
		Body: createIssueStoreCell,
	},
}

err = w.Send(context.Background(), &message, true)
if err != nil {
	panic(err)
}
```


## Get store info
Like in example below you can get other contract data:
```
func GetOwner(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (*address.Address, error)
func GetName(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error)
func GetDescription(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error)
func GetImage(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error)
func GetStoreWebhook(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (string, error)
func GetMccCode(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (int64, error)
func GetIsActive(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (bool, error)
func GetStoreVersion(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (uint64, error)
```
Or get all in one:
```
func GetStoreData(api *ton.APIClient, block *ton.BlockIDExt, addr *address.Address) (StoreData, error)
```
