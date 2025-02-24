package graph

import (
    "Go-Server/graph/model"
    "context"
    "fmt"
)


func (r *queryResolver) GetCoin(ctx context.Context, mintAddress string) (*model.Coin, error) {
    
    fakeDatabase := map[string]*model.Coin{
        "0x123456": {MintAddress: strPtr("0x123456"), MarketCap: floatPtr(1000000.50), Name: strPtr("Bitcoin")},
        "0xabcdef": {MintAddress: strPtr("0xabcdef"), MarketCap: floatPtr(500000.75), Name: strPtr("Ethereum")},
    }

 
    if coin, exists := fakeDatabase[mintAddress]; exists {
        return coin, nil
    }

    return nil, fmt.Errorf("coin not found for mintAddress: %s", mintAddress)
}


func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }


func strPtr(s string) *string {
    return &s
}


func floatPtr(f float64) *float64 {
    return &f
}
