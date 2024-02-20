package paprika

import "time"

type Coins []*Coin

type Coin struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Symbol       string       `json:"symbol"`
	Rank         int          `json:"rank"`
	IsActive     bool         `json:"is_active"`
	ExchangesIDs ExchangesIDs `json:"exchanges"`
	MarketCapVolume
	ExtendedInfo
}

type ExchangesIDs map[string]struct{}

type MarketCapVolume struct {
	MarketCap uint64 `json:"market_cap"`
	Volume    uint64 `json:"volume"`
}

type ExtendedInfo struct {
	IsNew             bool           `json:"is_new"`
	IsActive          bool           `json:"is_active"`
	Type              string         `json:"type"`
	Logo              string         `json:"logo"`
	Description       string         `json:"description"`
	StartedAt         time.Time      `json:"started_at"`
	DevelopmentStatus string         `json:"development_status"`
	Links             Links          `json:"links"`
	LinksExtended     []ExtendedLink `json:"links_extended"`
	FirstDataAt       time.Time      `json:"first_data_at"`
	LastDataAt        time.Time      `json:"last_data_at"`
}

type Links struct {
	Explorer   []string `json:"explorer"`
	Facebook   []string `json:"facebook"`
	Reddit     []string `json:"reddit"`
	SourceCode []string `json:"source_code"`
	Website    []string `json:"website"`
	YouTube    []string `json:"youtube"`
}

type ExtendedLink struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type Exchanges []*Exchange

type Exchange struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type PaprikaError struct {
	Error string `json:"error"`
}
