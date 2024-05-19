package models

import (
	"fmt"
	"github.com/0xBeppo/valueanalyzer/utils"
	"log"
	"os"
	"strconv"
)

const BASE_URL = "https://www.alphavantage.co/query"

type FinancialApi struct {
	ApiKey  string
	BaseUrl string
}

func NewFinancialApi() *FinancialApi {
	return &FinancialApi{
		ApiKey:  os.Getenv("API_KEY"),
		BaseUrl: BASE_URL,
	}
}

func (f *FinancialApi) GetOverview(ticker string) Ticker {
	url := fmt.Sprintf("%s?function=OVERVIEW&symbol=%s&apikey=%s", f.BaseUrl, ticker, f.ApiKey)

	tickerOverview, err := utils.DoApiCall[Ticker](url)
	if err != nil {
		log.Fatalf("failed to fetch data: %v", err)
	}

	return tickerOverview
}

func (f *FinancialApi) GetLastPrice(ticker string) float64 {
	url := fmt.Sprintf("%s?function=TIME_SERIES_INTRADAY&symbol=%s&interval=1min&apikey=%s", f.BaseUrl, ticker, f.ApiKey)

	log.Println(url)

	type MetaData struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		Interval      string `json:"4. Interval"`
		OutputSize    string `json:"5. Output Size"`
		TimeZone      string `json:"6. Time Zone"`
	}

	type TimeSeriesEntry struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	}

	type ApiResponse struct {
		MetaData       MetaData                   `json:"Meta Data"`
		TimeSeries1Min map[string]TimeSeriesEntry `json:"Time Series (1min)"`
	}

	response, err := utils.DoApiCall[ApiResponse](url)
	if err != nil {
		log.Fatalf("failed to fetch data: %v", err)
	}
	lastPrice := response.TimeSeries1Min[response.MetaData.LastRefreshed].Close

	log.Printf("Last price: %s", lastPrice)

	price, err := strconv.ParseFloat(lastPrice, 64)
	if err != nil {
		log.Fatalf("failed to convert last price to float: %v", err)
	}

	return price
}

func (f *FinancialApi) GetIncomeStatement(ticker string) IncomeStatement {
	url := fmt.Sprintf("%s?function=INCOME_STATEMENT&symbol=%s&apikey=%s", f.BaseUrl, ticker, f.ApiKey)

	incomeStatement, err := utils.DoApiCall[IncomeStatement](url)
	if err != nil {
		log.Fatalf("failed to fetch data: %v", err)
	}

	return incomeStatement
}

func (f *FinancialApi) GetBalanceSheet(ticker string) BalanceSheet {
	url := fmt.Sprintf("%s?function=BALANCE_SHEET&symbol=%s&apikey=%s", BASE_URL, ticker, f.ApiKey)
	var balanceSheet BalanceSheet
	balanceSheet, err := utils.DoApiCall[BalanceSheet](url)
	if err != nil {
		log.Fatalf("failed to fetch data: %v", err)
	}

	return balanceSheet
}

func (f *FinancialApi) GetEPS(ticker string) float64 {
	incomeStatement := f.GetIncomeStatement(ticker)
	balanceSheet := f.GetBalanceSheet(ticker)

	netIncome, err := strconv.Atoi(incomeStatement.AnnualReports[0].NetIncome)
	if err != nil {
		log.Fatalf("failed to convert net income to int: %v", err)
	}
	commonStockSharesOutstanding, err := strconv.Atoi(balanceSheet.AnnualReports[0].CommonStockSharesOutstanding)
	if err != nil {
		log.Fatalf("failed to convert net income to int: %v", err)
	}

	eps := float64(netIncome) / float64(commonStockSharesOutstanding)

	return eps
}

func (f *FinancialApi) CalculateNetDebt(ticker string) int {
	balanceSheet := f.GetBalanceSheet(ticker)

	currentDebt, err := strconv.Atoi(balanceSheet.QuarterlyReports[0].CurrentDebt)
	if err != nil {
		log.Fatalf("failed to convert current debt to int: %v", err)
	}
	longTermDebt, err := strconv.Atoi(balanceSheet.QuarterlyReports[0].LongTermDebt)
	if err != nil {
		log.Fatalf("failed to convert long term debt to int: %v", err)
	}
	cashAndEquivalents, err := strconv.Atoi(balanceSheet.QuarterlyReports[0].CashAndCashEquivalentsAtCarryingValue)
	if err != nil {
		log.Fatalf("failed to convert cash and equivalents to int: %v", err)
	}
	shortTermInvestments, err := strconv.Atoi(balanceSheet.QuarterlyReports[0].ShortTermInvestments)
	if err != nil {
		log.Fatalf("failed to convert short term investments to int: %v", err)
	}

	totalDebt := currentDebt + longTermDebt

	totalCash := cashAndEquivalents + shortTermInvestments

	netDebt := totalDebt - totalCash

	return netDebt
}

func (f *FinancialApi) CalculateEnterpriseValue(ticker string) int {

	return 0
}
