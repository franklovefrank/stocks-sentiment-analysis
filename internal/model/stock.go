package model

type StockQuery struct {
	Symbol    string
	StartDate string
	EndDate   string
}

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
	Query   StockQuery  `json:"query"`
	Records []StockData `json:"records"`
}
