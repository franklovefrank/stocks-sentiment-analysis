package service

import (
	"stock-sentiment-cli/internal/api"
	"stock-sentiment-cli/internal/model"
)

type BlueskyService struct {
	blueskyClient *api.BlueskyClient
}

func NewBlueskyService(blueskyClient *api.BlueskyClient) *BlueskyService {
	return &BlueskyService{blueskyClient: blueskyClient}
}

func (s *BlueskyService) GetPosts(query, token, since, until string) (*model.BlueskyResponse, error) {
	response, err := s.blueskyClient.FetchPosts(query, token, since, until)
	if err != nil {
		return nil, err
	}

	// Additional processing of response data can be done here if needed.
	return response, nil
}
