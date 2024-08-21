package facade

import (
	"stock-sentiment-cli/internal/api"
	"stock-sentiment-cli/internal/model"
)

type APIFacadeImpl struct {
	yahooClient     *api.YahooFinanceClient
	blueskyClient   *api.BlueskyClient
	googleNLPClient *api.GoogleNLPClient
}

func NewAPIFacade(yahooClient *api.YahooFinanceClient, blueskyClient *api.BlueskyClient, googleNLPClient *api.GoogleNLPClient) *APIFacadeImpl {
	return &APIFacadeImpl{
		yahooClient:     yahooClient,
		blueskyClient:   blueskyClient,
		googleNLPClient: googleNLPClient,
	}
}

func (f *APIFacadeImpl) FetchStockData(query model.StockQuery) (*model.StockHistory, error) {
	return f.yahooClient.FetchData(query)
}

func (f *APIFacadeImpl) FetchPosts(query model.StockQuery) (*model.BlueskyResponse, error) {
	return f.blueskyClient.FetchPosts(query)
}

func (f *APIFacadeImpl) AnalyzeSentiment(text string) (*model.SentimentResult, error) {
	return f.googleNLPClient.Analyze(text)
}
