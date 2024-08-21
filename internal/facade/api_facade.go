package facade

import (
	"stock-sentiment-cli/internal/model"
)

type APIFacade interface {
	FetchStockData(symbol, startDate, endDate string) (*model.StockHistory, error)
	FetchPosts(query, token, since, until string) (*model.BlueskyResponse, error)
	AnalyzeSentiment(text string) (*model.SentimentResult, error)
}
