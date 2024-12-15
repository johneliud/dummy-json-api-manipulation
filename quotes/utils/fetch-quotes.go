package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"quotes/models"
)

// FetchQuotes uses a GET request to fetch quotes from the provided URL. Fetches every available quotes by specifying the limit query to 0.
func FetchQuotes() (models.Quotes, error) {
	url := "https://dummyjson.com/quotes?limit=0"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Quotes{}, fmt.Errorf("error creating request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return models.Quotes{}, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.Quotes{}, fmt.Errorf("received status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Quotes{}, fmt.Errorf("error reading response body: %v", err)
	}

	var quotes models.Quotes
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return models.Quotes{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return quotes, nil
}
