# Initialization

Before starting to use the SDK, you need to prepare the environment. The **tonpay-go-sdk** relies on [tonutils-go](https://github.com/xssnick/tonutils-go) for interacting with the TON Blockchain. To initialize the necessary objects, follow these steps:
```
// First of all, we need to create connection pool
client := liteclient.NewConnectionPool()

// Choose testnet or mainnet
configUrl := "https://ton-blockchain.github.io/global.config.json" // mainnet
// configUrl := "https://ton-blockchain.github.io/testnet-global.config.json"  // testnet
err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
if err != nil {
log.Println(err)
return  false
}

api := ton.NewAPIClient(client)

// Get block
block, err := api.CurrentMasterchainInfo(context.Background())
if err != nil {
	log.Println(err)
	return  false
}
```
Make sure to update the `configURL` based on whether you want to connect to the mainnet or testnet.
