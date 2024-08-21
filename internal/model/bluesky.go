package model

type BlueskyResponse struct {
	Posts []struct {
		Record struct {
			Text string `json:"text"`
		} `json:"record"`
	} `json:"posts"`
}
