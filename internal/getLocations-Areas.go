package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/GaussHammer/pokedex/internal/pokecache"
)

var Cache *pokecache.Cache

func init() {
	Cache = pokecache.NewCache(10 * time.Second)
}

func GetLocationAreas(config *Config) error {
	url := config.Next
	var body []byte
	if cachedData, found := Cache.Get(url); found {
		fmt.Println("Using cached data for location areas")
		body = cachedData
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
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

	// Display the locations (names) to the user
	for _, location := range apiResponse.Results {
		fmt.Println(location.Name)
	}

	// Update globalConfig with new pagination state
	config.Next = apiResponse.Next
	config.Previous = apiResponse.Previous

	return nil
}

type Config struct {
	Next     string
	Previous string
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
