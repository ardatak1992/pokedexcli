package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	commands := populateCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Name\tDescription")
	fmt.Println("--------------------")
	for _, command := range commands {
		fmt.Printf("%s\t%s\n", command.name, command.description)
	}

	return nil
}
