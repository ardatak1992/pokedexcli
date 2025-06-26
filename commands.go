package main

type CliCommand struct {
	name        string
	description string
	callback    func(conf *Config, params ...string) error
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
		"explore": {
			name:        "explore",
			description: "displays the pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "random chance of catching a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists pokemon you caught",
			callback:    CommandPokedex,
		},
	}

}
