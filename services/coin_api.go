
// API response struct
package services


import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"sync"
)


var cache sync.Map

func StartDataCollection() {
	
// GatherAboutToGraduateCoins()
// GatherBundleData()



}

func GatherAboutToGraduateCoins () {
	url := "https://advanced-api-v2.pump.fun/coins/about-to-graduate"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close() 

	body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    fmt.Println("Error reading response:", err)
    return
}

var coins []CoinData 

if err := json.Unmarshal(body, &coins); err != nil {
    fmt.Println("Error decoding JSON:", err)
    return
}


for _, coin := range coins {
    cache.Store(coin.CoinMint, coin) 
	fmt.Println(coin)
}

cache.Range(func(key, value interface{}) bool {
	fmt.Printf("Key: %v, Value: %+v\n", key, value)
	return true
})

	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// fmt.Println("Response:", string(body)) // Print response as string
}

func GatherBundleData() {

}