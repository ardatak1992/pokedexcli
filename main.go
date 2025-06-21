package main

import (
	"time"

	"github.com/ardatak1992/pokedexcli/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(10 * time.Second)
	
	cfg := &Config{
		PokeapiClient: pokeClient,
	}

	startRepl(cfg)

}
