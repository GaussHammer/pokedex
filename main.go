package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/GaussHammer/pokedex/internal"
)

var pokeapiConfig = internal.Config{
	Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	Previous: "",
}
var commandMap map[string]cliCommand

func init() {
	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func(args []string) error {
				return commandExit()
			},
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func(args []string) error {
				return commandHelp()
			},
		},
		"map": {
			name:        "map",
			description: "Lists next areas from the Pokemon World",
			callback: func(args []string) error {
				// Pass the global or pre-defined config to GetLocationAreas
				return internal.GetLocationAreas(&pokeapiConfig)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous areas from the Pokemon World",
			callback: func(args []string) error {
				return internal.GetPreviousLocationsAreas(&pokeapiConfig)
			},
		},
		"explore": {
			name:        "explore",
			description: "explore an area from the pokemon world",
			callback: func(args []string) error {
				if len(args) == 0 {
					return fmt.Errorf("explore command requires a location")
				}
				return internal.ExploreLocations(args[0])
			},
		},
		"catch": {
			name:        "catch",
			description: "catch a Pokemon",
			callback: func(args []string) error {
				if len(args) == 0 {
					return fmt.Errorf("catch command requires a Pokemon")
				}
				return internal.CatchPokemon(args[0])
			},
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		parts := strings.Split(input, " ")

		command := parts[0]
		args := parts[1:] // This creates a slice containing all parts after the command

		cmd, ok := commandMap[command]
		if ok {
			err := cmd.callback(args)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	trimmmedText := strings.TrimSpace(lowerText)

	return strings.Fields(trimmmedText)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range commandMap {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
	config      *internal.Config
}
