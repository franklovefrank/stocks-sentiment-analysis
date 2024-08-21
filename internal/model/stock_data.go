package model

type StockData struct {
	Date     string  `json:"date"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	AdjClose float64 `json:"adj_close"`
	Volume   int     `json:"volume"`
}

type StockHistory struct {
	Symbol    string      `json:"symbol"`
	StartDate string      `json:"start_date"`
	EndDate   string      `json:"end_date"`
	Records   []StockData `json:"records"`
}
