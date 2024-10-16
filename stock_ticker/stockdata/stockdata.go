package stockdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type TimeSeriesDaily struct {
	Close string `json:"4. close"`
}

type APIResponse struct {
	TimeSeries map[string]TimeSeriesDaily `json:"Time Series (Daily)"`
}

type DailyPrice struct {
	Date  string
	Price float64
}

func GetStockData(symbol string, ndays int, apikey string) ([]DailyPrice, float64, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s", apikey, symbol)
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, 0, fmt.Errorf("failed to decode JSON: %v", err)
	}

	// Collect and parse dates
	var dates []time.Time
	dateMap := make(map[time.Time]TimeSeriesDaily)
	for dateStr, dailyData := range apiResponse.TimeSeries {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Printf("Failed to parse date %s: %v", dateStr, err)
			continue
		}
		dates = append(dates, date)
		dateMap[date] = dailyData
	}

	// Sort dates in descending order (latest first)
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].After(dates[j])
	})

	// Ensure we have at least ndays worth of data
	if len(dates) < ndays {
		return nil, 0, fmt.Errorf("not enough data available for the last %d days", ndays)
	}

	// Select the latest ndays worth of data
	var closingPrices []DailyPrice
	total := 0.0
	for i := 0; i < ndays; i++ {
		date := dates[i]
		dailyData := dateMap[date]

		closePrice, err := strconv.ParseFloat(dailyData.Close, 64)
		if err != nil {
			log.Printf("Failed to parse closing price for date %s: %v", date.Format("2006-01-02"), err)
			continue
		}

		// Format date to MM-DD-YYYY
		formattedDate := date.Format("01-02-2006")

		closingPrices = append(closingPrices, DailyPrice{Date: formattedDate, Price: closePrice})
		total += closePrice
	}

	avgPrice := total / float64(len(closingPrices))
	return closingPrices, avgPrice, nil
}
