package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/GaussHammer/pokedex/internal/pokecache"
)

func init() {
	Cache = pokecache.NewCache(10 * time.Second)
}

func ExploreLocations(args string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args
	var body []byte
	if cachedData, found := Cache.Get(url); found {
		fmt.Println("Using cached data from previous search")
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
	type pokemonType struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	var apiExplore struct {
		Pokemon_encounters []struct {
			Pokemon pokemonType `json:"pokemon"`
		} `json:"pokemon_encounters"`
	}

	if err := json.Unmarshal(body, &apiExplore); err != nil {
		return err
	}

	for _, pokemon := range apiExplore.Pokemon_encounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
