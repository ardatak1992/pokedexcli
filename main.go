package main

import (
	"log"
	"time"

	"github.com/ardatak1992/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokedex := pokedex{entries: map[string]pokeapi.Pokemon{}}

	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}

	err := startRepl(cfg)
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
