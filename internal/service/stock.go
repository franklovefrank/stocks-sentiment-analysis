package service

import (
	"fmt"
	"stock-sentiment-cli/internal/facade"
	"stock-sentiment-cli/internal/model"
	"stock-sentiment-cli/internal/util"
)

type StockService struct {
	apiFacade facade.APIFacade
}

func NewStockService(apiFacade facade.APIFacade) *StockService {
	return &StockService{apiFacade: apiFacade}
}

func (s *StockService) CalculateReturnsAndDrawdown(query model.StockQuery) (float64, float64, error) {
	stockHistory, err := s.apiFacade.FetchStockData(query)
	if err != nil {
		return 0, 0, err
	}

	prices := extractPrices(stockHistory)
	if len(prices) < 2 {
		return 0, 0, fmt.Errorf("not enough data to calculate returns and drawdown")
	}

	startPrice := prices[0]
	endPrice := prices[len(prices)-1]
	maxDrawdown := calculateDrawdown(prices)

	simpleReturn := calculateSimpleReturn(startPrice, endPrice)
	return util.RoundToTwoDecimalPlaces(simpleReturn), util.RoundToTwoDecimalPlaces(maxDrawdown), nil
}

func extractPrices(stockHistory *model.StockHistory) []float64 {
	prices := make([]float64, len(stockHistory.Records))
	for i, record := range stockHistory.Records {
		prices[i] = record.AdjClose
	}
	return prices
}

func calculateSimpleReturn(startPrice, endPrice float64) float64 {
	return (endPrice - startPrice) / startPrice * 100
}

func calculateDrawdown(prices []float64) float64 {
	if len(prices) == 0 {
		return 0
	}

	maxDrawdown := 0.0
	peakPrice := prices[0]

	for _, price := range prices {
		if price > peakPrice {
			peakPrice = price
		}

		currentDrawdown := (peakPrice - price) / peakPrice * 100
		if currentDrawdown > maxDrawdown {
			maxDrawdown = currentDrawdown
		}
	}

	return maxDrawdown
}
