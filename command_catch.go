package main

import (
	"fmt"
	"math"
	"math/rand"
)

func getCatchProbabilty(baseExperience int) float64 {
	midPoint := 171.0
	steepness := 0.02
	sigmoid := 1 / (1 + math.Pow(math.E, (steepness*(float64(baseExperience)-midPoint))))
	return sigmoid
}

func commandCatch(cfg *config, args []string) error {

	if len(args) == 0 {
		fmt.Println("usage: catch <pokemon_name>")
		return nil
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(&pokemonName)
	if err != nil {
		return err
	}

	if _, ok := cfg.pokedex.entries[pokemon.Name]; ok {
		fmt.Printf("You already catched %s\n", pokemon.Name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchProbability := getCatchProbabilty(pokemon.BaseExperience)
	catchAttempt := rand.Float64()

	if catchAttempt < catchProbability {
		fmt.Printf("You caught a %s\n", pokemon.Name)
		cfg.pokedex.entries[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped\n", pokemon.Name)
	}

	return nil
}
