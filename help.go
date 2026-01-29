package main

import (
	"fmt"
	"strings"
)

func commandHelp() error {
	var display strings.Builder

	display.WriteString("Welcome to the Pokedex!\n")
	display.WriteString("Usage:\n\n")

	commands := getCommands()

	for _, command := range commands {
		statement := fmt.Sprintf("%s: %s\n", command.name, command.description)
		display.WriteString(statement)
	}
	fmt.Println(display.String())

	return nil
}
