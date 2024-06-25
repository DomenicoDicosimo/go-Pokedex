package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DomenicoDicosimo/go-Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.RespPokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Pokedex Started. Type 'exit' to quit.")

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if len(words) == 0 {
			continue
		}
		command := words[0]
		args := words[1:]

		if command, found := getCommands()[command]; found {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", command)
			continue
		}
	}
}
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name:        "mapf",
			description: "Lists the next set of locations in Pokemon",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous set of locations in Pokemon",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists the encounters in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches a Pokemon!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspects a pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
