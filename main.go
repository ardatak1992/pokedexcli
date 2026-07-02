package main

import "log"

func main() {
	err := startRepl()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
