package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config , []string) error
}

func populateCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commmandExit,
		},
		"help": {
			name:        "help",
			description: "Lists the available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Lists previous locations",
			callback:    commandMapb,
		},
		"explore": {
			name: "explore",
			description: "lists all pokemon in an area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "tries to catch the given pokemon",
			callback: commandCatch,
		},
	}
}
