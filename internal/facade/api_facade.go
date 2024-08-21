package facade

import (
	"stock-sentiment-cli/internal/model"
)

type APIFacade interface {
	FetchStockData(query model.StockQuery) (*model.StockHistory, error)
	FetchPosts(query model.StockQuery) (*model.BlueskyResponse, error)
	AnalyzeSentiment(text string) (*model.SentimentResult, error)
}
