package service

/*
didn't have time to implement the details of each strategy but this is approximately the design pattern I had in mind

type SentimentService struct {
	nlpClient *api.GoogleNLPClient
}

func NewSentimentService(nlpClient *api.GoogleNLPClient) *SentimentService {
	return &SentimentService{nlpClient: nlpClient}
}
type SentimentStrategy interface {
	Analyze(tweet string) (AnalysisResult, error)
}

type BasicSentimentStrategy struct{}

func (b *BasicSentimentStrategy) Analyze(tweets []Tweet) (AnalysisResult, error) {
}

type AdvancedSentimentStrategy struct{}

func (a *AdvancedSentimentStrategy) Analyze(tweets []Tweet) (AnalysisResult, error) {
}

type CustomSentimentStrategy struct {
	// Config options for custom behavior
	Config map[string]interface{}
}

func (c *CustomSentimentStrategy) Analyze(tweets []Tweet) (AnalysisResult, error) {
}

*/
