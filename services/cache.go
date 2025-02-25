package services

import (
	"sync"
)


var CoinCache sync.Map


var InProgress sync.Map


type CoinData struct {

	CoinMint             string  `json:"coinMint"`
	Dev                  string  `json:"dev"`
	Name                 string  `json:"name"`
	Ticker               string  `json:"ticker"`
	ImageURL             string  `json:"imageUrl"`
	CreationTime         int64   `json:"creationTime"`
	NumHolders           int     `json:"numHolders"`
	MarketCap            float64 `json:"marketCap"`
	Volume               float64 `json:"volume"`
	BondingCurveProgress float64 `json:"bondingCurveProgress"`
	SniperCount          int     `json:"sniperCount"`
	CurrentMarketPrice   float64 `json:"currentMarketPrice"`
	BundleSupply         float64 `json:"bundleSupply"`
	OwnerHoldings        float64 `json:"ownerHoldings"`
	TenHolderSupply      float64 `json:"tenHolderSupply"`

}



