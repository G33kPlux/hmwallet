package hdwallet

func init() {
	coins[DASH] = newDASH
}

type dash struct {
	*btc
}

func newDASH(key *Key) Wallet {
	token := newBTC(key).(*btc)
	token.name = "Dash