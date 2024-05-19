package models

type Ticker struct {
	Symbol            string `json:"Symbol"`
	Name              string `json:"Name"`
	Exchange          string `json:"Exchange"`
	Currency          string `json:"Currency"`
	Country           string `json:"Country"`
	Sector            string `json:"Sector"`
	Industry          string `json:"Industry"`
	MarketCap         string `json:"MarketCapitalization"`
	SharesOutstanding string `json:"SharesOutstanding"`
	EarningsPerShare  string
}
