package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"stock-sentiment-cli/internal/model"
	"stock-sentiment-cli/internal/util"
)

type BlueskyClient struct{}

func (c *BlueskyClient) FetchPosts(query, token, since, until string) (*model.BlueskyResponse, error) {
	baseURL := "https://public.api.bsky.app/xrpc/app.bsky.feed.searchPosts"
	params := url.Values{}
	params.Add("q", query)
	if since != "" {
		sinceISO8601, err := util.ToISO8601(since)
		if err != nil {
			return nil, err
		}
		params.Add("since", sinceISO8601)
	}
	if until != "" {
		untilISO8601, err := util.ToISO8601(until)
		if err != nil {
			return nil, err
		}
		params.Add("until", untilISO8601)
	}
	fullURL := baseURL + "?" + params.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %s", res.Status)
	}

	var response model.BlueskyResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &response, nil
}
