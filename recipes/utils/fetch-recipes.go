package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"recipes/models"
)

// FetchRecipes uses a GET request to fetch recipes from the provided URL. Fetches every available recipe by specifying the limit query to 0.
func FetchRecipes() (models.Recipes, error) {
	url := "https://dummyjson.com/recipes?limit=0"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Recipes{}, fmt.Errorf("error creating request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return models.Recipes{}, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.Recipes{}, fmt.Errorf("received status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Recipes{}, fmt.Errorf("error reading response body: %v", err)
	}

	var Recipes models.Recipes
	err = json.Unmarshal(body, &Recipes)
	if err != nil {
		return models.Recipes{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return Recipes, nil
}