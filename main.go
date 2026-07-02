package main

import (
	"log"
	"time"

	"github.com/ardatak1992/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	err := startRepl(cfg)
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
