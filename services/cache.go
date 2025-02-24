package services

import (
	"sync"
)


var CoinCache sync.Map


var InProgress sync.Map


type CoinData struct {
	CoinMint       string
	OwnerHoldings  float64
	BundleSupply   float64
	TenHolderSupply float64
}
