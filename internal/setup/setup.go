package setup

import (
	"stock-sentiment-cli/internal/api"
	"stock-sentiment-cli/internal/facade"
	"stock-sentiment-cli/internal/service"
)

func InitializeServices() (*service.StockService, *service.SentimentService) {
	yahooClient := &api.YahooFinanceClient{}
	blueskyClient := &api.BlueskyClient{}
	googleNLPClient := &api.GoogleNLPClient{}

	apiFacade := facade.NewAPIFacade(yahooClient, blueskyClient, googleNLPClient)
	stockService := service.NewStockService(apiFacade)
	sentimentService := service.NewSentimentService(apiFacade)

	return stockService, sentimentService
}
