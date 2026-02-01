package pokedexclient

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theunhackable/pokedexcli/internal/models"
	cmd "github.com/theunhackable/pokedexcli/internal/pokedexcommands"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	trimmed := strings.Trim(lower, " ")
	splited := strings.Split(trimmed, " ")
	return splited
}

func StartRepl(config *models.Config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inp := scanner.Text()
		if inp == "" {
			continue
		}
		clean := cleanInput(inp)

		commands := cmd.GetCommands()
		command, exists := commands[clean[0]]

		if exists {
			if len(clean) == 2 {
				config.Param = &clean[1]
			}
			err := command.Callback(config)
			config.Param = nil
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found")
		}
	}
}
