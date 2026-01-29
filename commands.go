package main

func getCommands() map[string]CLICommand {
	commands := map[string]CLICommand{
		"map": {
			name:        "map",
			description: "Displays areas",
			callback:    commandMap,
		},
		"bmap": {
			name:        "bmap",
			description: "Displays prev areas",
			callback:    commandBMap,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	return commands
}
