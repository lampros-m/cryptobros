package services

import (
	"repositories/cryptobros/internal/repositories/domain"
	"repositories/cryptobros/internal/repositories/github"
	"repositories/cryptobros/internal/repositories/paprika"
)

func mapPaprikaCoinsAndExchangesToDomainCoins(paprikaCoins paprika.Coins, paprikaExchanges paprika.Exchanges) domain.Coins {
	domainCoins := domain.Coins{}

	for _, paprikaCoin := range paprikaCoins {
		domainCoin := domain.Coin{}

		domainCoin.ID = paprikaCoin.ID
		domainCoin.Name = paprikaCoin.Name
		domainCoin.Symbol = paprikaCoin.Symbol
		domainCoin.Rank = paprikaCoin.Rank
		domainCoin.IsActive = paprikaCoin.IsActive
		domainCoin.MarketCap = paprikaCoin.MarketCapVolume.MarketCap
		domainCoin.Volume = paprikaCoin.MarketCapVolume.Volume
		domainCoin.Type = paprikaCoin.ExtendedInfo.Type
		domainCoin.Logo = paprikaCoin.ExtendedInfo.Logo
		domainCoin.Description = paprikaCoin.ExtendedInfo.Description
		domainCoin.DevelopmentStatus = paprikaCoin.ExtendedInfo.DevelopmentStatus
		domainCoin.Links.Explorer = paprikaCoin.ExtendedInfo.Links.Explorer
		domainCoin.Links.Facebook = paprikaCoin.ExtendedInfo.Links.Facebook
		domainCoin.Links.Reddit = paprikaCoin.ExtendedInfo.Links.Reddit
		domainCoin.Links.SourceCode = paprikaCoin.ExtendedInfo.Links.SourceCode
		domainCoin.Links.Website = paprikaCoin.ExtendedInfo.Links.Website
		domainCoin.Links.YouTube = paprikaCoin.ExtendedInfo.Links.YouTube
		domainCoin.FirstDataAt = paprikaCoin.ExtendedInfo.FirstDataAt
		domainCoin.LastDataAt = paprikaCoin.ExtendedInfo.LastDataAt

		domainCoin.LinksExtended = domain.ExtendedLinks{}
		for _, paprikaExtendedLink := range paprikaCoin.ExtendedInfo.LinksExtended {
			domainExtendedLink := domain.ExtendedLink{}
			domainExtendedLink.URL = paprikaExtendedLink.URL
			domainExtendedLink.Type = paprikaExtendedLink.Type
			domainCoin.LinksExtended = append(domainCoin.LinksExtended, domainExtendedLink)
		}

		domainCoin.Exchanges = make(domain.Exchanges)
		for _, paprikaExchange := range paprikaExchanges {
			if _, found := paprikaCoin.ExchangesIDs[paprikaExchange.ID]; found {
				domainCoin.Exchanges[paprikaExchange.ID] = domain.ExchanngeInfo{
					Name: paprikaExchange.Name,
				}
			}
		}

		domainCoins = append(domainCoins, &domainCoin)
	}

	return domainCoins
}

func mapGithubRepositoriesToDomainRepositories(githubRepositories github.Repositories) domain.GitHubRepositories {
	domainRepositories := domain.GitHubRepositories{}

	for _, githubRepository := range githubRepositories {
		domainRepository := domain.GitHubRepository{}

		domainRepository.Name = githubRepository.Name
		domainRepository.UpdatedAt = githubRepository.UpdatedAt
		domainRepository.Stars = githubRepository.Stars

		domainRepositories = append(domainRepositories, domainRepository)
	}

	return domainRepositories
}
