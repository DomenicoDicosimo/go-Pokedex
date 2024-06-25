package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	caughtPokemon := cfg.caughtPokemon

	fmt.Println("Your Pokedex:")
	for _, pokemon := range caughtPokemon {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
