package service

import (
	"fmt"
	"strings"
	"sync"

	"stock-sentiment-cli/internal/facade"
	"stock-sentiment-cli/internal/model"
)

type SentimentService struct {
	apiFacade facade.APIFacade
}

func NewSentimentService(apiFacade facade.APIFacade) *SentimentService {
	return &SentimentService{apiFacade: apiFacade}
}

// fetches posts and analyzes sentiment.
func (s *SentimentService) AnalyzePosts(query model.StockQuery) (model.SentimentSummary, error) {
	postsResponse, err := s.apiFacade.FetchPosts(query)
	if err != nil {
		return model.SentimentSummary{}, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var positive, neutral, negative int

	for _, post := range postsResponse.Posts {
		postText := post.Record.Text
		if strings.TrimSpace(postText) == "" {
			continue
		}

		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			result, err := s.apiFacade.AnalyzeSentiment(text)
			if err != nil {
				fmt.Printf("Error analyzing sentiment: %v\n", err)
				return
			}

			mu.Lock()
			defer mu.Unlock()

			if result.Score > 0 {
				positive++
			} else if result.Score == 0 {
				neutral++
			} else {
				negative++
			}
		}(postText)
	}

	wg.Wait()

	summary := model.SentimentSummary{
		Positive: positive,
		Neutral:  neutral,
		Negative: negative,
		Total:    positive + neutral + negative,
	}

	return summary, nil
}
