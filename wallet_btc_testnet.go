
package hdwallet

func init() {
	coins[BTCTestnet] = newBTCTestnet
}

type btcTestnet struct {
	name   string
	symbol string
	key    *Key
}

func newBTCTestnet(key *Key) Wallet {
	key.Opt.Params = &BTCTestnetParams
	return &btcTestnet{
		name:   "Bitcoin Testnet",
		symbol: "BTCTestnet",
		key:    key,
	}
}