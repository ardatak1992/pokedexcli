package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}

type Place struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Config struct {
	Page     int
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Place `json:"results"`
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays last 20 location areas",
			callback:    commandMapb,
		},
	}

}

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(conf *Config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(conf *Config) error {

	url := ""
	if conf.Next != nil {
		url = *conf.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while responding: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading body: %v", err)
	}

	err = json.Unmarshal([]byte(body), conf)
	if err != nil {
		fmt.Println(err)
	}

	for _, place := range conf.Results {
		fmt.Println(place.Name)
	}

	fmt.Println()

	return nil
}

func commandMapb(conf *Config) error {
	url := ""
	if conf.Previous != nil {
		url = *conf.Previous
	} else {
		fmt.Println("You are on the first page")
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error while responding: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error while reading body: %v", err)
	}

	err = json.Unmarshal([]byte(body), conf)
	if err != nil {
		fmt.Println(err)
	}

	for _, place := range conf.Results {
		fmt.Println(place.Name)
	}

	fmt.Println()

	return nil

}
