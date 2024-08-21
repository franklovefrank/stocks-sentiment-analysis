package facade

import (
	"stock-sentiment-cli/internal/model"
)

type APIFacade interface {
	FetchStockData(stock_info model.StockQuery) (*model.StockHistory, error)
	FetchPosts(stock_info model.StockQuery) (*model.BlueskyResponse, error)
	AnalyzeSentiment(text string) (*model.SentimentResult, error)
}
