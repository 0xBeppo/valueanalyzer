package main

import (
	"github.com/0xBeppo/valueanalyzer/models"
	"log"
)

func main() {

	financialApi := models.NewFinancialApi()

	ev := financialApi.CalculateEnterpriseValue("AAPL")

	log.Printf("EV: %f", ev)

}
