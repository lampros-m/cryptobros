package domain

import "time"

type Coins []*Coin

type Coin struct {
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	Symbol            string        `json:"symbol"`
	Rank              int           `json:"rank"`
	IsActive          bool          `json:"is_active"`
	Exchanges         Exchanges     `json:"exchanges"`
	MarketCap         uint64        `json:"market_cap"`
	Volume            uint64        `json:"volume"`
	Type              string        `json:"type"`
	Logo              string        `json:"logo"`
	Description       string        `json:"description"`
	DevelopmentStatus string        `json:"development_status"`
	Links             Links         `json:"links"`
	LinksExtended     ExtendedLinks `json:"links_extended"`
	FirstDataAt       time.Time     `json:"first_data_at"`
	LastDataAt        time.Time     `json:"last_data_at"`
	GitHubRepositories
}

type ExchanngeInfo struct {
	Name string `json:"name"`
}

type Exchanges map[string]ExchanngeInfo

type Links struct {
	Explorer   []string `json:"explorer"`
	Facebook   []string `json:"facebook"`
	Reddit     []string `json:"reddit"`
	SourceCode []string `json:"source_code"`
	Website    []string `json:"website"`
	YouTube    []string `json:"youtube"`
}

type ExtendedLinks []ExtendedLink

type ExtendedLink struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type GitHubRepository struct {
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	Stars     int32     `json:"stargazers_count"`
}

type GitHubRepositories []GitHubRepository
