package main

import "fmt"

func CommandPokedex(cfg *Config, params ...string) error {

	if len(cfg.PokeapiClient.Pokedex) == 0 {
		fmt.Println("You have no pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pok := range cfg.PokeapiClient.Pokedex {
		fmt.Printf("- %s\n", pok.Name)
	}

	return nil
}
