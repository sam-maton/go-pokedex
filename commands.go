package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name    string
	desc    string
	command func(config *config) error
}

func commandHelp(config *config) error {
	commands := getCommands()

	for _, v := range commands {
		fmt.Printf(`
Name: %s

Description: %s

--------------
`, v.name, v.desc)
	}
	return nil
}

func commandExit(config *config) error {
	os.Exit(0)

	return nil
}

func commandMap(config *config) error {
	return nil
}

func commandMapB(config *config) error {
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:    "help",
			desc:    "Displays all of the available commands and their descriptions.",
			command: commandHelp,
		},
		"exit": {
			name:    "exit",
			desc:    "Exits the Pokedex",
			command: commandExit,
		},
		"map": {
			name:    "map",
			desc:    "Lists the next set of areas",
			command: commandMap,
		},
		"mapb": {
			name:    "mapb",
			desc:    "Lists the previous set of areas",
			command: commandMapB,
		},
	}
}
