package service

import (
	"stock-sentiment-cli/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockAPIFacade struct {
	mock.Mock
}

func (m *MockAPIFacade) FetchPosts(query model.StockQuery) (*model.BlueskyResponse, error) {
	args := m.Called(query)
	return args.Get(0).(*model.BlueskyResponse), args.Error(1)
}

func (m *MockAPIFacade) AnalyzeSentiment(text string) (*model.SentimentResult, error) {
	args := m.Called(text)
	return args.Get(0).(*model.SentimentResult), args.Error(1)
}

func (m *MockAPIFacade) FetchStockData(query model.StockQuery) (*model.StockHistory, error) {
	args := m.Called(query)
	return args.Get(0).(*model.StockHistory), args.Error(1)
}
