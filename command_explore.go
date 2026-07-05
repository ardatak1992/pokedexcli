package main

import "fmt"

func commandExplore(cfg *config, args []string) error{
	if len(args) == 0 {
		fmt.Println("usage: explore <location>")
		return nil
	}

	location := args[0]
	

	return nil
}
