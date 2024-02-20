package main

import (
	"log"
	"os"
	"repositories/cryptobros/internal/config"
	"repositories/cryptobros/internal/services"
)

var (
	DebugMode = true

	// Query Documentation
	// ===================
	// DateOfData: 				Date of data retrieval
	// 	MarketCapThresholdDown: Minimum market cap
	// 	MarketCapThresholdUp: 	Maximum market cap
	// 	VolumeThresholdDown: 	Minimum volume
	// 	ExchangesIncluded: 		Exchanges to be included - At least a number of exchanges
	// 	ExchangesExcluded: 		Exchanges to be excluded - Max number of exchanges permitted
	// 	YearThreshold: 			Year threshold of project's data
	// 	SourceCode: 			Source code must be available
	//
	// Query Example
	// =============
	// "20240219"
	// 1000000000
	// 5000000000
	// 5000000
	// []string{"mexc", "gateio", "okx"}, 2
	// []string{"binance", "coinbase"}, 1
	// "2023"
	//
	Query = services.Query{
		DateOfData:             "20240220",
		MarketCapThresholdDown: 1000000,
		MarketCapThresholdUp:   1000000000,
		VolumeThresholdDown:    100000,
		ExchangesIncluded: services.QueryExchanges{
			ExchangesIDs: []string{"kucoin", "bybit", "bitfinex"},
			Number:       0,
		},
		ExchangesExcluded: services.QueryExchanges{
			ExchangesIDs: []string{"binance"},
			Number:       0,
		},
		YearThreshold: "1900",
		SourceCode:    true,
	}
)

func main() {
	config := config.NewConfig(DebugMode)
	coinsQuerier := services.NewCoinsQuerier(config)

	err := coinsQuerier.QueryCoins(Query)
	if err != nil {
		log.Println("error making query: ", err)
		os.Exit(1)
	}

	log.Println("Query executed successfully")
}
