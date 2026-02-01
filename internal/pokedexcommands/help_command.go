package pokedexcommands

import (
	"fmt"
	"strings"

	"github.com/theunhackable/pokedexcli/internal/models"
)

func commandHelp(config *models.Config) error {
	var display strings.Builder

	display.WriteString("Welcome to the Pokedex!\n")
	display.WriteString("Usage:\n\n")

	commands := GetCommands()

	for _, command := range commands {
		statement := fmt.Sprintf("%s: %s\n", command.Name, command.Description)
		display.WriteString(statement)
	}
	fmt.Println(display.String())

	return nil
}
