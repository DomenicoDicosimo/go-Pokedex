package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("enter the name of a pokemon")
	}

	pokemon := args[0]

	if mon, ok := cfg.caughtPokemon[pokemon]; !ok {
		return fmt.Errorf("you have not caught that pokemon")
	} else {
		fmt.Println("Name: ", mon.Name)
		fmt.Println("Height: ", mon.Height)
		fmt.Println("Height: ", mon.Weight)
		fmt.Println("Stats: ")
		for _, stat := range mon.Stats {
			fmt.Printf("  -%s: %d \n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types: ")
		for _, typ := range mon.Types {
			fmt.Printf("  -%s \n", typ.Type.Name)
		}
	}
	return nil
}
