package main

import (
	"fmt"
)

func commandMap(cfg *config) error {

	locations, err := cfg.pokeapiClient.GetAreas(cfg.nextURL)
	if err != nil {
		return fmt.Errorf("Error in GetAreas: %v", err)
	}

	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {

	if cfg.previousURL == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	locations, err := cfg.pokeapiClient.GetAreas(cfg.previousURL)
	if err != nil {
		return fmt.Errorf("Error in GetAreas: %v", err)
	}

	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
