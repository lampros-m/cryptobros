package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"repositories/cryptobros/internal/config"
	"repositories/cryptobros/internal/repositories/domain"
	"strings"
	"time"
)

var (
	QueryYearThresholdDateFormat = "2006"
)

type Repository struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

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
	if !query.SourceCode {
		return true
	}

	if len(coin.Links.SourceCode) == 0 {
		return false
	}

	//apiUrl := convertToAPIUrl(coin.Links.SourceCode[0])
	//
	//if apiUrl == "" {
	//	return false
	//}
	//
	//repos, err := fetchRepositories(apiUrl)
	//
	//if err != nil {
	//	fmt.Println("error fetching repos from api url: %s", apiUrl)
	//	return false
	//}
	//
	//if len(repos) < query.GithubReposThreshold {
	//	return false
	//}

	return true
}

func convertToAPIUrl(url string) string {
	// Removing the protocol and get the rest (e.g., github.com/mintlayer or github.com/opentensor/BitTensor)
	parts := strings.SplitN(url, "github.com/", 2)

	// Further split to isolate organization/user name
	if len(parts) < 2 {
		return "Invalid GitHub URL"
	}

	orgUserParts := strings.SplitN(parts[1], "/", 2)

	// Construct the API URL
	apiUrl := fmt.Sprintf("https://api.github.com/orgs/%s/repos", orgUserParts[0])

	return apiUrl
}

func fetchRepositories(apiURL string) ([]Repository, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
