package services

var (
// DataDate                    = "20240219"                      // Date of data retrieval
// MarketCapThresholdDown      = 1000000000                      // Minimum market cap
// MarketCapThresholdUp        = 5000000000                      // Maximum market cap
// VolumeThreshold             = 5000000                         // Minimum volume
// ExchangesIncludedAtLeastOne = []string{"mexc", "gateio"}      // At least one of these exchanges must be included
// ExchangesExcludedAll        = []string{"binance", "coinbase"} // All of these exchanges must be excluded
// YearThreshold               = "2023"                          // Year threshold of project launch
)

type Query struct {
	DateOfData             string         // Date of data retrieval
	MarketCapThresholdDown uint64         // Minimum market cap
	MarketCapThresholdUp   uint64         // Maximum market cap
	VolumeThresholdDown    uint64         // Minimum volume
	ExchangesIncluded      QueryExchanges // Exchanges to be included - At least a number of exchanges
	ExchangesExcluded      QueryExchanges // Exchanges to be excluded - Max number of exchanges permitted
	YearThreshold          string         // Year threshold of project launch
	SourceCode             bool           // Source code must be available
}

type QueryExchanges struct {
	ExchangesIDs []string
	Number       uint32
}
