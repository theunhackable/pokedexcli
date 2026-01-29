package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	trimmed := strings.Trim(lower, " ")
	splited := strings.Split(trimmed, " ")
	return splited
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inp := scanner.Text()
		if inp == "" {
			continue
		}
		clean := cleanInput(inp)

		commands := getCommands()
		command, exists := commands[clean[0]]

		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found")
		}
	}
}
