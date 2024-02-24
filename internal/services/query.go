package services

import (
	"errors"
	"log"
	"repositories/cryptobros/internal/config"
	"repositories/cryptobros/internal/repositories/domain"
	"time"
)

var (
	QueryYearThresholdDateFormat = "2006"
)

type CoinsQuerierInterface interface {
	QueryCoins(query Query) error
}

type CoinsQuerier struct {
	Config *config.Config
}

func NewCoinsQuerier(config *config.Config) *CoinsQuerier {
	return &CoinsQuerier{
		Config: config,
	}
}

func (o *CoinsQuerier) QueryCoins(query Query) error {
	var err error

	coins, err := domain.GetCoinsForSpecificDay(query.DateOfData)
	if err != nil {
		return errors.New("error QueryCoins cannot GetCoinsForSpecificDay:" + err.Error())
	}

	filteredCoins := domain.Coins{}
	for _, coin := range coins {
		// Launch date
		launchDateValid, err := isLaunchDateValid(coin, query)
		if err != nil {
			return errors.New("error QueryCoins cannot isLaunchDateValid:" + err.Error())
		}
		if !launchDateValid {
			continue
		}

		// Market Cap - Volume
		marketCapVolumeValid := isMarketCapVolumeValid(coin, query)
		if !marketCapVolumeValid {
			continue
		}

		// Exchanges excluded
		exchangesExcludedValid := isExchangesExcludedValid(coin, query)
		if !exchangesExcludedValid {
			continue
		}

		// Exchanges included
		exchangesIncludedValid := isExchangesIncludedValid(coin, query)
		if !exchangesIncludedValid {
			continue
		}

		// Sourcecode
		sourceCodeValid := isSourceCodeValid(coin, query)
		if !sourceCodeValid {
			continue
		}

		filteredCoins = append(filteredCoins, coin)
	}

	if o.Config.DebugMode {
		log.Println("Number of coins after quering:", len(filteredCoins))
	}

	// save
	err = domain.SaveReults(filteredCoins)
	if err != nil {
		return errors.New("error QueryCoins cannot save results:" + err.Error())
	}

	return nil
}

func isLaunchDateValid(coin *domain.Coin, query Query) (bool, error) {
	yearThresholdParsed, err := time.Parse(QueryYearThresholdDateFormat, query.YearThreshold)
	if err != nil {
		return false, errors.New("error isLaunchDateValid cannot Parse time: " + err.Error())
	}

	projectLaunchDate := coin.FirstDataAt
	if projectLaunchDate.Before(yearThresholdParsed) {
		return false, nil
	}

	return true, nil
}

func isMarketCapVolumeValid(coin *domain.Coin, query Query) bool {
	if (coin.MarketCap < query.MarketCapThresholdDown || coin.MarketCap > query.MarketCapThresholdUp) && coin.MarketCap != 0 {
		return false
	}

	if coin.MarketCap == 0 && coin.Volume < query.VolumeThresholdDown {
		return false
	}

	return true
}

func isExchangesExcludedValid(coin *domain.Coin, query Query) bool {
	maxPermitted := query.ExchangesExcluded.Number
	counter := uint32(0)

	for _, exchange := range query.ExchangesExcluded.ExchangesIDs {
		if _, found := coin.Exchanges[exchange]; found {
			counter++
		}
	}

	return counter <= maxPermitted
}

func isExchangesIncludedValid(coin *domain.Coin, query Query) bool {
	atLeast := query.ExchangesIncluded.Number
	counter := uint32(0)

	for _, exchange := range query.ExchangesIncluded.ExchangesIDs {
		if _, found := coin.Exchanges[exchange]; found {
			counter++
		}
	}

	return counter >= atLeast
}

func isSourceCodeValid(coin *domain.Coin, query Query) bool {
	//
	// TODO: Code here ...
	//

	return true
}
