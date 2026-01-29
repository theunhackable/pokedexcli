package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	return fmt.Errorf("Closing the Pokedex... Goodbye!")
}

func commandHelp() error {
	commands := []cliCommand{
		{
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	var display strings.Builder

	display.WriteString("Welcome to the Pokedex!\n")
	display.WriteString("Usage:\n\n")

	for _, command := range commands {
		statement := fmt.Sprintf("%s: %s\n", command.name, command.description)
		display.WriteString(statement)
	}

	return fmt.Errorf("%s", display.String())
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inp := scanner.Text()
		clean := cleanInput(inp)
		switch clean[0] {
		case "exit":
			err := commandExit()
			fmt.Println(err)
			os.Exit(0)
		case "help":
			err := commandHelp()
			fmt.Println(err)

		}
	}
}
