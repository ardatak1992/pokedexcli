package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ardatak1992/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		inputArr := cleanInput(reader.Text())
		params := []string{}
		if len(inputArr) > 1 {
			params = inputArr[1:]
		}
		if com, ok := commands[inputArr[0]]; ok {
			err := com.callback(cfg, params...)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	seperatedStrings := strings.Fields(text)

	return seperatedStrings
}
