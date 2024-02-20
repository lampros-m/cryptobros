package services

import (
	"errors"
	"log"
	"repositories/cryptobros/internal/config"
	"repositories/cryptobros/internal/networking"
	"repositories/cryptobros/internal/repositories/domain"
	"repositories/cryptobros/internal/repositories/paprika"
)

type CoinsCreatorInterface interface {
	CreateTodayCoins() error
}

type CoinsCreator struct {
	Config *config.Config
}

func NewCoinsCreator(config *config.Config) *CoinsCreator {
	return &CoinsCreator{
		Config: config,
	}
}

func (o *CoinsCreator) CreateTodayCoins() error {
	var err error

	// request a new ipv6
	err = networking.RenewIPv6()
	if err != nil {
		return errors.New("error CreateTodayCoins cannot RenewIPv6: " + err.Error())
	}

	// get all active coins
	paprikaActiveCoins, err := paprika.GetAllActiveCoins()
	if err != nil {
		return errors.New("error CreateTodayCoins cannot paprika.GetAllActiveCoins:" + err.Error())
	}

	if o.Config.DebugMode {
		log.Println("Number of active coins to fetched:", len(paprikaActiveCoins))
	}

	// get all active exchanges
	paprikaActiveExchanges, err := paprika.GetAllActiveExchanges()
	if err != nil {
		return errors.New("error CreateTodayCoins cannot paprika.GetAllActiveExchanges:" + err.Error())
	}

	// code for testing and debugging
	if o.Config.DebugMode {
		// testCoins := paprika.Coins{}
		// for _, coin := range paprikaActiveCoins {
		// 	if coin.ID == "ml-mintlayer" || coin.ID == "btc-bitcoin" {
		// 		testCoins = append(testCoins, coin)
		// 	}
		// }
		// paprikaActiveCoins = testCoins
		paprikaActiveCoins = paprikaActiveCoins[:200]
	}

	// encance market cap volume
	for _, coin := range paprikaActiveCoins {
		for {
			marketCapVolume, err := paprika.GetMarketCapAndVolumeByCoinID(coin.ID)
			if err != nil {
				switch err {
				case paprika.ErrCoinIDNotFound:
					break
				case paprika.ErrLimitData, paprika.ErrUnknown:
					err = networking.RenewIPv6()
					if err != nil {
						return errors.New("error CreateTodayCoins cannot RenewIPv6: " + err.Error())
					}
					continue
				default:
					return errors.New("error CreateTodayCoins cannot paprika.GetMarketCapAndVolumeByCoinID:" + err.Error())
				}
			}

			if o.Config.DebugMode {
				log.Println("encance market cap:", coin.Name, coin.Rank, marketCapVolume.MarketCap, marketCapVolume.Volume)
			}

			coin.MarketCapVolume = marketCapVolume
			break
		}
	}

	// encance exchanges
	for _, coin := range paprikaActiveCoins {
		for {
			exchangesIDs, err := paprika.GetExchangesByCoinID(coin.ID)
			if err != nil {
				switch err {
				case paprika.ErrCoinIDNotFound:
					break
				case paprika.ErrLimitData, paprika.ErrUnknown:
					err = networking.RenewIPv6()
					if err != nil {
						return errors.New("error CreateTodayCoins cannot RenewIPv6: " + err.Error())
					}
					continue
				default:
					return errors.New("error CreateTodayCoins cannot paprika.GetExchangesByCoinID:" + err.Error())
				}
			}

			if o.Config.DebugMode {
				log.Println("encance exchanges:", coin.Name, coin.Rank)
			}

			coin.ExchangesIDs = exchangesIDs
			break
		}
	}

	// encance extended info
	for _, coin := range paprikaActiveCoins {
		for {
			extendedInfo, err := paprika.GetExtendedInfoByCoinID(coin.ID)
			if err != nil {
				switch err {
				case paprika.ErrCoinIDNotFound:
					break
				case paprika.ErrLimitData, paprika.ErrUnknown:
					err = networking.RenewIPv6()
					if err != nil {
						return errors.New("error CreateTodayCoins cannot RenewIPv6: " + err.Error())
					}
					continue
				default:
					return errors.New("error CreateTodayCoins cannot paprika.GetExtendedInfoByCoinID:" + err.Error())
				}
			}

			if o.Config.DebugMode {
				log.Println("encance extended info:", coin.Name, coin.Rank)
			}

			coin.ExtendedInfo = extendedInfo
			break
		}
	}

	// mapping
	domainCoins := mapPaprikaCoinsAndExchangesToDomainCoins(paprikaActiveCoins, paprikaActiveExchanges)

	// save
	err = domain.SaveCoins(domainCoins)
	if err != nil {
		return errors.New("error CreateTodayCoins cannot save domain coins:" + err.Error())
	}

	return nil
}
