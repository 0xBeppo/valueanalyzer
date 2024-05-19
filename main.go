package main

import (
	"github.com/0xBeppo/valueanalyzer/models"
	"log"
)

func main() {

	financialApi := models.NewFinancialApi()

	lastPrice := financialApi.GetLastPrice("AAPL")
	log.Printf("Last Price: %f", lastPrice)

	overview := financialApi.GetOverview("AAPL")

	log.Printf("Overview: %v", overview)

	eps := financialApi.GetEPS("AAPL")

	log.Printf("EPS: %f", eps)

	//url := fmt.Sprintf("%s?function=GLOBAL_QUOTE&symbol=AAPL&apikey=%s", BASE_URL, API_KEY)

}
