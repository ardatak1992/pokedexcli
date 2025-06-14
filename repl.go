package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	commands := getCommands()
	config := Config{Page: 0}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		inputArr := cleanInput(reader.Text())
		if com, ok := commands[inputArr[0]]; ok {
			com.callback(&config)
		}

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	seperatedStrings := strings.Fields(text)

	return seperatedStrings
}
