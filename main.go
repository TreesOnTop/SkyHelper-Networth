package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"constants"
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

type NetworthError struct {
	Message string
}

func (e *NetworthError) Error() string {
	return e.Message
}

type PricesError struct {
	Message string
}

func (e *PricesError) Error() string {
	return e.Message
}

type Options struct {
	V2Endpoint     bool
	Cache          bool
	OnlyNetworth   bool
	Prices         map[string]float64
	ReturnItemData bool
	MuseumData     interface{}
}

type ProfileData struct {
	Currencies struct {
		CoinPurse float64 `json:"coin_purse"`
	} `json:"currencies"`
	CoinPurse float64 `json:"coin_purse"`
	Profile   struct {
		BankAccount float64 `json:"bank_account"`
	} `json:"profile"`
}

type Item struct {
	Tag interface{} `json:"tag"`
	Exp interface{} `json:"exp"`
}

var cachedPrices struct {
	Prices    map[string]float64
	LastCache time.Time
}

var isLoadingPrices bool

func getNetworth(profileData ProfileData, bankBalance float64, options Options) (map[string]interface{}, error) {
	if profileData == (ProfileData{}) {
		return nil, &NetworthError{Message: "Invalid profile data provided"}
	}
	purse := profileData.CoinPurse
	if options.V2Endpoint {
		purse = profileData.Currencies.CoinPurse
	}
	personalBankBalance := 0.0
	if options.V2Endpoint {
		personalBankBalance = profileData.Profile.BankAccount
	}
	prices, err := parsePrices(options.Prices, options.Cache)
	if err != nil {
		return nil, err
	}
	items, err := parseItems(profileData, options.MuseumData, options.V2Endpoint)
	if err != nil {
		return nil, err
	}
	return calculateNetworth(items, purse, bankBalance, personalBankBalance, prices, options.OnlyNetworth, options.ReturnItemData), nil
}

func getPreDecodedNetworth(profileData ProfileData, items map[string]interface{}, bankBalance float64, options Options) (map[string]interface{}, error) {
	purse := profileData.CoinPurse
	if options.V2Endpoint {
		purse = profileData.Currencies.CoinPurse
	}
	personalBankBalance := 0.0
	if options.V2Endpoint {
		personalBankBalance = profileData.Profile.BankAccount
	}
	err := postParseItems(profileData, items, options.V2Endpoint)
	if err != nil {
		return nil, err
	}
	prices, err := parsePrices(options.Prices, options.Cache)
	if err != nil {
		return nil, err
	}
	return calculateNetworth(items, purse, bankBalance, personalBankBalance, prices, options.OnlyNetworth, options.ReturnItemData), nil
}

func getItemNetworth(item Item, options Options) (map[string]interface{}, error) {
	if item.Tag == nil && item.Exp == nil {
		return nil, &NetworthError{Message: "Invalid item provided"}
	}
	prices, err := parsePrices(options.Prices, options.Cache)
	if err != nil {
		return nil, err
	}
	return calculateItemNetworth(item, prices, options.ReturnItemData), nil
}

func parsePrices(prices map[string]float64, cache bool) (map[string]float64, error) {
	if prices != nil {
		for id, price := range prices {
			prices[id] = price
		}
	}
	if prices == nil {
		return getPrices(cache, 3)
	}
	return prices, nil
}

func getPrices(cache bool, retries int) (map[string]float64, error) {
	if retries <= 0 {
		return nil, &PricesError{Message: "Failed to retrieve prices"}
	}
	if cachedPrices.LastCache.After(time.Now().Add(-5 * time.Minute)) && cache {
		return cachedPrices.Prices, nil
	}
	if isLoadingPrices {
		for isLoadingPrices {
			time.Sleep(100 * time.Millisecond)
		}
		return getPrices(cache, retries)
	}
	isLoadingPrices = true
	resp, err := http.Get("https://raw.githubusercontent.com/SkyHelperBot/Prices/main/prices.json")
	if err != nil {
		isLoadingPrices = false
		if retries <= 0 {
			return nil, &PricesError{Message: fmt.Sprintf("Failed to retrieve prices with status code %v", err)}
		}
		fmt.Printf("[SKYHELPER-NETWORTH] Failed to retrieve prices with status code %v. Retrying (%d attempt(s) left)...\n", err, retries)
		return getPrices(cache, retries-1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		isLoadingPrices = false
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		isLoadingPrices = false
		return nil, err
	}
	prices := make(map[string]float64)
	for item, price := range data {
		prices[item] = price.(float64)
	}
	cachedPrices = struct {
		Prices    map[string]float64
		LastCache time.Time
	}{Prices: prices, LastCache: time.Now()}
	isLoadingPrices = false
	return prices, nil
}

func checkForUpdate() {
	resp, err := http.Get("https://registry.npmjs.org/skyhelper-networth")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var packageInfo map[string]interface{}
	err = json.Unmarshal(body, &packageInfo)
	if err != nil {
		return
	}
	latestVersion := packageInfo["dist-tags"].(map[string]interface{})["latest"].(string)
	currentVersion := "1.25.0"
	if latestVersion != currentVersion {
		fmt.Printf("[SKYHELPER-NETWORTH] An update is available! Current version: %s, Latest version: %s\n", currentVersion, latestVersion)
	}
}

func main() {
	checkForUpdate()
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	for range ticker.C {
		checkForUpdate()
	}
}
