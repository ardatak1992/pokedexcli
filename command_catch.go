package main

import "fmt"

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

	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)

	fmt.Printf("%s - %d xp\n", pokemon.Name, pokemon.BaseExperience)

	return nil
}
