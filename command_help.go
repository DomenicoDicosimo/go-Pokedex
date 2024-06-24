package main

import (
	"fmt"
)

func commandHelp(args []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
