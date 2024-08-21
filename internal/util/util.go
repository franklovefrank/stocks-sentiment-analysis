package util

import (
	"math"
	"time"
)

// converts a date string in YYYY-MM-DD format to a Unix timestamp (in seconds).
func DateToUnix(date string) (int64, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// converts a date string in YYYY-MM-DD format to ISO8601 format.
func ToISO8601(dateStr string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}
	return parsedDate.Format("2006-01-02T15:04:05Z07:00"), nil
}

func RoundToTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}
