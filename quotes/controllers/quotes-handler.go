package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"quotes/utils"
)

// QuotesHandler encodes the quotes.Quotes slice into JSON and sends it in the HTTP response that serves the frontend.
func QuotesHandler(w http.ResponseWriter, r *http.Request) {
	quotes, err := utils.FetchQuotes()
	if err != nil {
		log.Printf("Error fetching quotes: %v\n", err)
		http.Error(w, fmt.Sprintf("Error fetching quotes: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes.Quotes)
}
