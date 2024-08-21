package facade

import (
	"stock-sentiment-cli/internal/model"
)

type APIFacade interface {
	FetchStockData(symbol, startDate, endDate string) (string, error)
	FetchPosts(symbol, token, startDate, endDate string) (string, error)
	AnalyzeSentiment(text string) (model.SentimentResult, error)
}
