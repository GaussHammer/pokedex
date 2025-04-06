package internal

import "fmt"

func InspectPokemon(args string) error {
	_, ok := Pokedex[args]
	if !ok {
		return fmt.Errorf("%s wasn't caught", args)
	}
	pokemon := Pokedex[args]
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.Base_stat)
	}
	fmt.Println("Types:")
	for _, pokemontype := range pokemon.Types {
		fmt.Printf("- %s\n", pokemontype.Type.Name)
	}
	return nil
}
