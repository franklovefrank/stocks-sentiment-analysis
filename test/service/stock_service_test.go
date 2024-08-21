package service

import (
	"stock-sentiment-cli/internal/model"
	"stock-sentiment-cli/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateReturnsAndDrawdown(t *testing.T) {
	mockFacade := new(MockAPIFacade)
	stockService := service.NewStockService(mockFacade)

	query := model.StockQuery{
		Symbol:    "AAPL",
		StartDate: "2023-01-01",
		EndDate:   "2023-01-31",
	}

	mockFacade.On("FetchStockData", query.Symbol, query.StartDate, query.EndDate).Return(&model.StockHistory{
		Records: []model.StockData{
			{AdjClose: 150},
			{AdjClose: 155},
			{AdjClose: 145},
			{AdjClose: 160},
		},
	}, nil)

	returns, drawdown, err := stockService.CalculateReturnsAndDrawdown(query)
	assert.NoError(t, err)
	assert.Equal(t, 6.67, returns)
	assert.Equal(t, 9.09, drawdown)
}
