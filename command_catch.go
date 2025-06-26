package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *Config, params ...string) error {

	if len(params) == 0 {
		fmt.Println("Please add a pokemon name")
		fmt.Println("catch <pokemon_name>")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])
	pokemon, ok := cfg.PokeapiClient.Pokedex[params[0]]

	if ok {
		fmt.Printf("%s is already caught.\n", pokemon.Name)
		return nil
	}

	pokemon, err := cfg.PokeapiClient.ListStats(params[0])

	if err != nil {
		return fmt.Errorf("error getting the base experience: %v", err)
	}

	catchProbability := calculateCatchPropability(pokemon.BaseExperience)

	catchTry := rand.Float64()

	if catchTry < catchProbability {
		fmt.Printf("%s was caught!\n", params[0])
		fmt.Println("You may inspect it with the inspect command.")
		cfg.PokeapiClient.Pokedex[params[0]] = pokemon
	} else {
		fmt.Printf("%s escaped\n", pokemon.Name)
	}

	return nil
}

func calculateCatchPropability(base_exp int) float64 {

	return math.Pow(math.E, -float64(base_exp)/200)
}
