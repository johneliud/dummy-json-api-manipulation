package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"quotes/models"
)

// FetchQuotes uses a GET request to fetch quotes from the provided URL. Fetches every available quotes by specifying the limit query to 0.
func FetchQuotes() (models.Quotes, error) {
	url := "https://dummyjson.com/quotes?limit=0"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return models.Quotes{}, fmt.Errorf("error creating request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error making HTTP request: %v\n", err)
		return models.Quotes{}, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Invalid status code %v received.\n", res.StatusCode)
		return models.Quotes{}, fmt.Errorf("received status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response: %v\n", err)
		return models.Quotes{}, fmt.Errorf("error reading response body: %v", err)
	}

	var quotes models.Quotes
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v\n", err)
		return models.Quotes{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return quotes, nil
}
