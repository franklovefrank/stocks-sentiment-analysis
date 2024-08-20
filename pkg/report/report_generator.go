package report

import (
	"fmt"
)

func GenerateReport(stockData string, sentimentData map[string]float64) {
	fmt.Println("Stock Data:", stockData)
	fmt.Println("Sentiment Analysis:")
	for sentiment, percentage := range sentimentData {
		fmt.Printf("%s: %.2f%%\n", sentiment, percentage)
	}
}
