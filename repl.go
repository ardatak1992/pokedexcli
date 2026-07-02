package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}

func startRepl() error {
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

		if !ok {
			fmt.Printf("Unknown command: %s\n", inputSlice[0])
			continue
		}

		err := command.callback()
		if err != nil {
			return err
		}

	}
}
