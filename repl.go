package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Pokedex Started. Type 'exit' to quit.")

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		args := cleanInput(scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if len(args) == 0 {
			continue
		}
		command := args[0]
		args = args[1:]

		if cmd, found := getCommands()[command]; found {
			err := cmd.callback(args)
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
}
