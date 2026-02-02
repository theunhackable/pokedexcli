package commands

import "github.com/theunhackable/pokedexcli/internal/models"

func GetCommands() map[string]models.CLICommand {
	commands := map[string]models.CLICommand{
		"map": {
			Name:        "map",
			Description: "Displays areas",
			Callback:    commandMap,
		},
		"bmap": {
			Name:        "bmap",
			Description: "Displays prev areas",
			Callback:    commandBMap,
		},
		"explore": {
			Name:        "explore",
			Description: "See the pokemons in the area.",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon.",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inpect details of caught pokemon.",
			Callback:    commandInspect,
		},

		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},

		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}

	return commands
}
