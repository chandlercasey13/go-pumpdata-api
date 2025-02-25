// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model




type Coin struct {
	Mint                   *string  `json:"mint,omitempty"`
	Name                   *string  `json:"name,omitempty"`
	Symbol                 *string  `json:"symbol,omitempty"`
	Description            *string  `json:"description,omitempty"`
	ImageURL               *string  `json:"imageUrl,omitempty"`
	MetadataURL            *string  `json:"metadataUrl,omitempty"`
	Twitter                *string  `json:"twitter,omitempty"`
	Telegram               *string  `json:"telegram,omitempty"`
	BondingCurve           *string  `json:"bondingCurve,omitempty"`
	AssociatedBondingCurve *string  `json:"associatedBondingCurve,omitempty"`
	Creator                *string  `json:"creator,omitempty"`
	CreatedTimestamp       *int32   `json:"createdTimestamp,omitempty"`
	RaydiumPool            *string  `json:"raydiumPool,omitempty"`
	Complete               *bool    `json:"complete,omitempty"`
	VirtualSolReserves     *float64 `json:"virtualSolReserves,omitempty"`
	VirtualTokenReserves   *float64 `json:"virtualTokenReserves,omitempty"`
	TotalSupply            *float64 `json:"totalSupply,omitempty"`
	Website                *string  `json:"website,omitempty"`
	ShowName               *bool    `json:"showName,omitempty"`
	KingOfTheHillTimestamp *int32   `json:"kingOfTheHillTimestamp,omitempty"`
	MarketCap              *float64 `json:"marketCap,omitempty"`
	ReplyCount             *int32   `json:"replyCount,omitempty"`
	LastReply              *int32   `json:"lastReply,omitempty"`
	Nsfw                   *bool    `json:"nsfw,omitempty"`
	MarketID               *string  `json:"marketId,omitempty"`
	Inverted               *bool    `json:"inverted,omitempty"`
	IsCurrentlyLive        *bool    `json:"isCurrentlyLive,omitempty"`
	Username               *string  `json:"username,omitempty"`
	ProfileMint            *string  `json:"profileMint,omitempty"`
	UsdMarketCap           *float64 `json:"usdMarketCap,omitempty"`
}







type Query struct {
}
