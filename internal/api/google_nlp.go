package api

import (
	"context"
	"fmt"
	"stock-sentiment-cli/internal/model"

	language "cloud.google.com/go/language/apiv2"
	"cloud.google.com/go/language/apiv2/languagepb"
)

type GoogleNLPClient struct{}

func (c *GoogleNLPClient) Analyze(text string) (*model.SentimentResult, error) {
	ctx := context.Background()

	client, err := language.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})

	if err != nil {
		return nil, fmt.Errorf("AnalyzeSentiment: %w", err)
	}

	result := &model.SentimentResult{
		Text:      text,
		Score:     resp.DocumentSentiment.Score,
		Magnitude: resp.DocumentSentiment.Magnitude,
	}
	return result, nil
}
