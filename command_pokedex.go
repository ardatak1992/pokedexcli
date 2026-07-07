package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex.entries) == 0 {
		fmt.Println("You haven't caught any pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.pokedex.entries {
		fmt.Printf("- %s\n", key)
	}

	return nil
}
