package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchStockData(symbol string, startDate, endDate string) (string, error) {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%s&period2=%s&interval=1d", symbol, startDate, endDate)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
