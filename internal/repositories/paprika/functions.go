package paprika

import (
	"encoding/json"
	"errors"
	"fmt"
	"repositories/cryptobros/internal/networking"
)

var (
	PaprikaEndpointListCoins             = "https://api.coinpaprika.com/v1/coins"
	PaprikaEndpointListExhanges          = "https://api.coinpaprika.com/v1/exchanges"
	PaprikaEndpointOHLC                  = "https://api.coinpaprika.com/v1/coins/%s/ohlcv/latest"
	PaprikaEndpointExchangesByCoinID     = "https://api.coinpaprika.com/v1/coins/%s/exchanges"
	PaprikaEndpointExtendedInfoByCoindID = "https://api.coinpaprika.com/v1/coins/%s"
)

func GetAllActiveCoins() (Coins, error) {
	body, err := networking.GetRequest(PaprikaEndpointListCoins)
	if err != nil {
		return Coins{}, errors.New("error GetAllActiveCoins cannot GetRequest: " + err.Error())
	}

	var coins Coins
	err = json.Unmarshal(body, &coins)
	if err != nil {
		return Coins{}, errors.New("error GetAllActiveCoins cannot Unmarshal: " + err.Error())
	}

	var activeCoins Coins
	for _, coin := range coins {
		if coin.IsActive {
			activeCoins = append(activeCoins, coin)
		}
	}

	return activeCoins, nil
}

func GetAllActiveExchanges() (Exchanges, error) {
	body, err := networking.GetRequest(PaprikaEndpointListExhanges)
	if err != nil {
		return Exchanges{}, errors.New("error GetAllActiveExchanges cannot GetRequest: " + err.Error())
	}

	var exchanges Exchanges
	err = json.Unmarshal(body, &exchanges)
	if err != nil {
		return Exchanges{}, errors.New("error GetAllActiveExchanges cannot Unmarshal: " + err.Error())
	}

	var activeExchanges Exchanges
	for _, exchange := range exchanges {
		if exchange.Active {
			activeExchanges = append(activeExchanges, exchange)
		}
	}

	return activeExchanges, nil
}

func GetMarketCapAndVolumeByCoinID(coinID string) (MarketCapVolume, error) {
	url := fmt.Sprintf(PaprikaEndpointOHLC, coinID)
	body, err := networking.GetRequest(url)
	if err != nil {
		return MarketCapVolume{}, errors.New("error GetMarketCapAndVolumeByCoinID cannot GetRequest:" + err.Error())
	}

	var paprikaError PaprikaError
	_ = json.Unmarshal(body, &paprikaError)

	if paprikaError.Error == ErrCoinIDNotFound.Error() {
		return MarketCapVolume{}, nil
	}

	if paprikaError.Error == ErrLimitData.Error() {
		return MarketCapVolume{}, ErrLimitData
	}

	if paprikaError.Error != "" {
		return MarketCapVolume{}, ErrUnknown
	}

	var marketCapVolume []MarketCapVolume
	_ = json.Unmarshal(body, &marketCapVolume)

	response := MarketCapVolume{}
	if len(marketCapVolume) > 0 {
		response = marketCapVolume[0]
	}

	return response, nil
}

func GetExchangesByCoinID(coinID string) (ExchangesIDs, error) {
	url := fmt.Sprintf(PaprikaEndpointExchangesByCoinID, coinID)
	body, err := networking.GetRequest(url)
	if err != nil {
		return ExchangesIDs{}, errors.New("error GetExchangesByCoinID cannot GetRequest:" + err.Error())
	}

	var paprikaError PaprikaError
	_ = json.Unmarshal(body, &paprikaError)

	if paprikaError.Error == ErrCoinIDNotFound.Error() {
		return ExchangesIDs{}, nil
	}

	if paprikaError.Error == ErrLimitData.Error() {
		return ExchangesIDs{}, ErrLimitData
	}

	if paprikaError.Error != "" {
		return ExchangesIDs{}, ErrUnknown
	}

	var exchanges []Exchange
	_ = json.Unmarshal(body, &exchanges)

	exchangesIDs := make(ExchangesIDs)
	for _, exchange := range exchanges {
		exchangesIDs[exchange.ID] = struct{}{}
	}

	return exchangesIDs, nil
}

func GetExtendedInfoByCoinID(coinID string) (ExtendedInfo, error) {
	url := fmt.Sprintf(PaprikaEndpointExtendedInfoByCoindID, coinID)
	body, err := networking.GetRequest(url)
	if err != nil {
		return ExtendedInfo{}, errors.New("error GetExtendedInfoByCoinID cannot GetRequest: " + err.Error())
	}

	var paprikaError PaprikaError
	_ = json.Unmarshal(body, &paprikaError)

	if paprikaError.Error == ErrCoinIDNotFound.Error() {
		return ExtendedInfo{}, nil
	}

	if paprikaError.Error == ErrLimitData.Error() {
		return ExtendedInfo{}, ErrLimitData
	}

	if paprikaError.Error != "" {
		return ExtendedInfo{}, ErrUnknown
	}

	var extendedInfo ExtendedInfo
	_ = json.Unmarshal(body, &extendedInfo)

	return extendedInfo, nil
}
