package service

/*
didn't have time to implement the details of each strategy but wanted to do something like this for the sentiment service

type SentimentStrategy interface {
	Analyze(tweets []Tweet) (AnalysisResult, error)
}

type BasicSentimentStrategy struct{}

func (b *BasicSentimentStrategy) Analyze(tweets []Tweet) (AnalysisResult, error) {
	return AnalysisResult{}, nil
}

type AdvancedSentimentStrategy struct{}

func (a *AdvancedSentimentStrategy) Analyze(tweets []Tweet) (AnalysisResult, error) {
	return AnalysisResult{}, nil
}

type CustomSentimentStrategy struct {
	Config map[string]interface{}
}

func (c *CustomSentimentStrategy) Analyze(tweets []Tweet) (AnalysisResult, error) {
	return AnalysisResult{}, nil
}

type Tweet struct {
	Text string
}

type AnalysisResult struct {
	Score float64
}

*/
