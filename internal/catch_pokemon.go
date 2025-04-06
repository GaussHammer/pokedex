package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func init() {
	Pokedex = make(map[string]Pokemon)
}

func CatchPokemon(args string) error {
	var body []byte
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + args)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s isn't a Pokemon", args)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args)
	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}
	catchChance := 100 - pokemon.Base_experience/4
	random := rand.Intn(100)
	if random < catchChance {
		Pokedex[args] = pokemon
		fmt.Printf("%s was caught!\n", args)
	} else {
		fmt.Printf("%s escaped!\n", args)
	}
	return nil
}
