package model

type SentimentResult struct {
	Text      string  `json:"text"`
	Score     float32 `json:"score"`
	Magnitude float32 `json:"magnitude"`
}

type SentimentSummary struct {
	Positive int `json:"positive"`
	Neutral  int `json:"neutral"`
	Negative int `json:"negative"`
	Total    int `json:"total"`
}
