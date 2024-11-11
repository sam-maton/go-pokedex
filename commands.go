package main

import (
	"fmt"
	"os"
)

type CLICommand struct {
	name    string
	desc    string
	command func(config *Config) error
}

func commandHelp(config *Config) error {
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

func commandExit(config *Config) error {
	os.Exit(0)

	return nil
}

func commandMap(config *Config) error {
	results, err := config.API.GetAreas(config.locationNext)

	for _, r := range results.Results {
		fmt.Println(r.Name)
	}

	config.locationPrev = results.Previous
	config.locationNext = results.Next

	return err
}

func commandMapB(config *Config) error {
	results, err := config.API.GetAreas(config.locationPrev)
	for _, r := range results.Results {
		fmt.Println(r.Name)
	}
	config.locationPrev = results.Previous
	config.locationNext = results.Next
	return err
}

func getCommands() map[string]CLICommand {
	return map[string]CLICommand{
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
