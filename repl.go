package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ardatak1992/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {

	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}

type pokedex struct {
	entries map[string]pokeapi.Pokemon
}

type config struct {
	pokeapiClient pokeapi.Client
	previousURL   *string
	nextURL       *string
	pokedex       pokedex
}

func startRepl(cfg *config) error {
	scanner := bufio.NewScanner(os.Stdin)
	commands := populateCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		inputSlice := cleanInput(userInput)

		if len(inputSlice) == 0 {
			continue
		}

		command, ok := commands[inputSlice[0]]
		args := inputSlice[1:]

		if !ok {
			fmt.Printf("Unknown command: %s\n", inputSlice[0])
			continue
		}

		err := command.callback(cfg, args)
		if err != nil {
			return err
		}

		scanner.Err()
	}
}
