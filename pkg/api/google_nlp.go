package api

import (
	"context"

	language "cloud.google.com/go/language/apiv1"
	"google.golang.org/api/option"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func AnalyzeSentiment(text string) (*languagepb.AnalyzeSentimentResponse, error) {
	ctx := context.Background()
	client, err := language.NewClient(ctx, option.WithCredentialsFile("path/to/credentials.json"))
	if err != nil {
		return nil, err
	}

	document := &languagepb.Document{
		Type:    languagepb.Document_PLAIN_TEXT,
		Content: text,
	}

	resp, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: document,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
