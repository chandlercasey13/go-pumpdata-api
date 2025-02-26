package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"Go-Server/graph/model"
	"context"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

)


func fetchData(url string) (*model.Coin, error) {
	resp, err := http.Get(url) 
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close() 


	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("API Raw Response:", string(body))
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	
	var rawData map[string]interface{}
	if err := json.Unmarshal(body, &rawData); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	
	coin := &model.Coin{
		Mint:         getString(rawData, "mint"),
		Name:         getString(rawData, "name"),
		Symbol:       getString(rawData, "symbol"),
		Description:  getString(rawData, "description"),
		ImageURL:     getString(rawData, "image_uri"),
		MetadataURL:  getString(rawData, "metadata_uri"),
		Twitter:      getString(rawData, "twitter"),
		Telegram:     getString(rawData, "telegram"),
		Creator:      getString(rawData, "creator"),
		Website:      getString(rawData, "website"),
		MarketCap:    getFloat(rawData, "market_cap"),
		UsdMarketCap: getFloat(rawData, "usd_market_cap"),
		
	}

	return coin, nil
}



func getString(data map[string]interface{}, key string) *string {
	if val, exists := data[key]; exists {
		if strVal, ok := val.(string); ok {
			return &strVal
		}
	}
	return nil
}

func getFloat(data map[string]interface{}, key string) *float64 {
	if val, exists := data[key]; exists {
		if floatVal, ok := val.(float64); ok {
			return &floatVal
		}
	}
	return nil
}

func getInt64(data map[string]interface{}, key string) *int64 {
	if val, exists := data[key]; exists {
		if floatVal, ok := val.(float64); ok {
			intVal := int64(floatVal) 
			return &intVal
		}
	}
	return nil
}






func (r *queryResolver) GetCoin(ctx context.Context, mintAddress string) (*model.Coin, error) {
	
	url := fmt.Sprintf("https://frontend-api-v3.pump.fun/coins/%s", mintAddress)
	


	coinData, err := fetchData(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching coin data: %v", err)
	}


	return coinData, nil
}


func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
