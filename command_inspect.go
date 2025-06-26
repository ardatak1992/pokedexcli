package main

import "fmt"

func commandInspect(cfg *Config, params ...string) error {

	if len(params) == 0 {
		fmt.Println("Please add a pokemon name")
		fmt.Println("inspect <pokemon_name>")
		return nil
	}

	pokemon, ok := cfg.PokeapiClient.Pokedex[params[0]]
	if !ok {
		fmt.Printf("%s wasn't caught.\n", params[0])
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("-%s\n", t.Type.Name)
	}

	return nil
}
