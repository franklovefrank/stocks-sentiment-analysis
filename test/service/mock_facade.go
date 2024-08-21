package service

import (
	"stock-sentiment-cli/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockAPIFacade struct {
	mock.Mock
}

func (m *MockAPIFacade) FetchPosts(symbol, token, startDate, endDate string) (*model.BlueskyResponse, error) {
	args := m.Called(symbol, token, startDate, endDate)
	return args.Get(0).(*model.BlueskyResponse), args.Error(1)
}

func (m *MockAPIFacade) AnalyzeSentiment(text string) (*model.SentimentResult, error) {
	args := m.Called(text)
	return args.Get(0).(*model.SentimentResult), args.Error(1)
}

func (m *MockAPIFacade) FetchStockData(symbol, startDate, endDate string) (*model.StockHistory, error) {
	args := m.Called(symbol, startDate, endDate)
	return args.Get(0).(*model.StockHistory), args.Error(1)
}
