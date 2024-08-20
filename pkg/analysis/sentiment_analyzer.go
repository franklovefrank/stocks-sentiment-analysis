package analysis

import (
	"stock-sentiment-cli/pkg/api"
)

func AnalyzeTweets(tweets string) (map[string]float64, error) {
	sentimentResponse, err := api.AnalyzeSentiment(tweets)
	if err != nil {
		return nil, err
	}

	sentimentData := make(map[string]float64)
	// Process sentimentResponse to fill sentimentData map

	return sentimentData, nil
}
