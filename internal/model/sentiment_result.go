package model

// SentimentResult represents the result of sentiment analysis for a piece of text.
type SentimentResult struct {
	Text      string  `json:"text"`
	Score     float32 `json:"score"`
	Magnitude float32 `json:"magnitude"`
}

// SentimentSummary summarizes the sentiment analysis results for multiple texts.
type SentimentSummary struct {
	Positive int `json:"positive"`
	Neutral  int `json:"neutral"`
	Negative int `json:"negative"`
	Total    int `json:"total"`
}
