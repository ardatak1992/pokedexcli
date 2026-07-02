package main

import (
	"fmt"
	"os"
)

func commmandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
