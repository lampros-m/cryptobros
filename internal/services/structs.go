package services

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
