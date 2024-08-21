package api

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"stock-sentiment-cli/internal/model"
	"stock-sentiment-cli/internal/util"
	"strconv"
)

type YahooFinanceClient struct{}

func (c *YahooFinanceClient) FetchData(symbol, startDate, endDate string) (*model.StockHistory, error) {
	startUnix, err := util.DateToUnix(startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %w", err)
	}
	endUnix, err := util.DateToUnix(endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %w", err)
	}
	url := buildYahooFinanceURL(symbol, startUnix, endUnix)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV data: %w", err)
	}

	if len(records) <= 1 {
		return nil, fmt.Errorf("no data returned")
	}

	var stockHistory model.StockHistory
	stockHistory.Symbol = symbol
	stockHistory.StartDate = startDate
	stockHistory.EndDate = endDate

	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		date := record[0]
		open, _ := strconv.ParseFloat(record[1], 64)
		high, _ := strconv.ParseFloat(record[2], 64)
		low, _ := strconv.ParseFloat(record[3], 64)
		close, _ := strconv.ParseFloat(record[4], 64)
		adjClose, _ := strconv.ParseFloat(record[5], 64)
		volume, _ := strconv.Atoi(record[6])

		stockData := model.StockData{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			AdjClose: adjClose,
			Volume:   volume,
		}

		stockHistory.Records = append(stockHistory.Records, stockData)
	}

	return &stockHistory, nil
}

func buildYahooFinanceURL(symbol string, startUnix, endUnix int64) string {
	return fmt.Sprintf(
		"https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%d&period2=%d&interval=1d&events=history&includeAdjustedClose=true",
		symbol, startUnix, endUnix)
}
