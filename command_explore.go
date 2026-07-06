package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("usage: explore <location>")
		return nil
	}

	areaName := args[0]
	loc, err := cfg.pokeapiClient.GetAreaInfo(&areaName)
	if err != nil {
		return err
	}

	

	for _, pokemon := range loc.PokemonEncounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
