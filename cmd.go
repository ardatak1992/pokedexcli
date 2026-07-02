package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
