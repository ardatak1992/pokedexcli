package main

import "fmt"

func commandInspect(cfg *config, args []string) error {

	if len(args) == 0 {
		fmt.Println("usage: inspect <pokemon_name>")
		return nil
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex.entries[pokemonName]
	if !ok {
		fmt.Printf("You didn't catch %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
