package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *Config, params ...string) error {
	if len(params) == 0 {
		return errors.New("please enter a location name")
	}

	fmt.Printf("Exploring %s...\n", params[0])
	pokemonList, err := conf.PokeapiClient.ListPokemon(params[0])
	if err != nil {
		return errors.New("error getting pokemon list")
	}

	fmt.Println("Found Pokemon:")
	for _, p := range pokemonList {
		fmt.Printf("- %s\n", p.Name)
	}

	return nil
}
