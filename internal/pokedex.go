package internal

import "fmt"

func YourPokedex() error {
	if len(Pokedex) == 0 {
		return fmt.Errorf("You haven't caught any Pokemon yet!")
	}
	for pokemon, _ := range Pokedex {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}
