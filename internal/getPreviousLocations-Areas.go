package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPreviousLocationsAreas(config *Config) error {
	if config.Previous == "" {
		fmt.Println("You're already on the first page")
		return nil
	}
	fmt.Println("DEBUG: Previous URL from config:", config.Previous)
	url := config.Previous
	var body []byte

	// Try to get from cache first
	if cachedData, found := Cache.Get(url); found {
		fmt.Println("Using cached data for previous location areas")
		body = cachedData
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Note: Using = instead of := to avoid variable shadowing
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// Add to cache for future use
		Cache.Add(url, body)
	}

	// Define the structure to unmarshal the JSON response
	var apiResponse struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	}

	// Parse the JSON response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}

	// Display the locations
	for _, location := range apiResponse.Results {
		fmt.Println(location.Name)
	}

	// Update config with new pagination state
	config.Next = apiResponse.Next
	config.Previous = apiResponse.Previous

	return nil
}
