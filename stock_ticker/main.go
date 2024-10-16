package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"stockticker/stockdata"
	"strconv"
)

type ResponseData struct {
	Symbol        string
	Days          int
	AveragePrice  float64
	ClosingPrices []stockdata.DailyPrice
}

func main() {
	// Retrieve environment variables
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		fmt.Println("API key is not set")
		return
	}

	symbol := os.Getenv("SYMBOL")
	if symbol == "" {
		fmt.Println("SYMBOL is not set")
		return
	}

	ndaysStr := os.Getenv("NDAYS")
	if ndaysStr == "" {
		fmt.Println("NDAYS is not set")
		return
	}

	// Convert NDAYS to an integer
	ndays, err := strconv.Atoi(ndaysStr)
	if err != nil || ndays <= 0 {
		fmt.Println("Invalid NDAYS value:", ndaysStr)
		return
	}

	// Set up the HTTP handler for the root path to render the home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			renderHome(w) // Serve the landing page at the root URL "/"
			return
		}
		http.NotFound(w, r) // Return 404 for any non-GET requests on "/"
	})

	// Set up the HTTP handler for the "/stock" path to process the form submission
	http.HandleFunc("/stock", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Use environment variables to get the stock data
			prices, avgPrice, err := stockdata.GetStockData(symbol, ndays, apiKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := ResponseData{
				Symbol:        symbol,
				Days:          ndays,
				AveragePrice:  avgPrice,
				ClosingPrices: prices,
			}

			renderResults(w, data)
			return
		}
		http.NotFound(w, r) // Return 404 for any non-POST requests on "/stock"
	})

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Render the home page into a readable format
func renderHome(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(filepath.Join("html", "home.html"))
	if err != nil {
		http.Error(w, "Error loading form html folder", http.StatusInternalServerError)
		log.Println("Error loading form html folder:", err)
		return
	}
	tmpl.Execute(w, nil)
}

// Render the results page inot a readable format
func renderResults(w http.ResponseWriter, data ResponseData) {
	tmpl, err := template.ParseFiles(filepath.Join("html", "results.html"))
	if err != nil {
		http.Error(w, "Error loading results html folder", http.StatusInternalServerError)
		log.Println("Error loading results html folder:", err)
		return
	}
	tmpl.Execute(w, data)
}
