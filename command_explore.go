package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("no location name provided")
	}

	locationName := args[0]

	EncountersResp, err := cfg.pokeapiClient.ListEncounters(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, mon := range EncountersResp.PokemonEncounters {
		fmt.Println(" -", mon.Pokemon.Name)
	}

	return nil
}
