package main

import (
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CLICommand struct {
	name    string
	desc    string
	command func(config *Config) error
	args    int
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

	if err != nil {
		return err
	}

	for _, r := range results.Results {
		fmt.Println(r.Name)
	}
	config.locationPrev = results.Previous
	config.locationNext = results.Next
	return nil
}

func commandExplore(config *Config) error {
	results, err := config.API.GetAreaPokemon(config.args[1])

	if err != nil {
		return err
	}

	for _, r := range results.Encounters {
		fmt.Println(cases.Title(language.English).String(r.Pokemon.Name))
	}
	return nil
}

func getCommands() map[string]CLICommand {
	return map[string]CLICommand{
		"help": {
			name:    "help",
			desc:    "Displays all of the available commands and their descriptions.",
			command: commandHelp,
			args:    0,
		},
		"exit": {
			name:    "exit",
			desc:    "Exits the Pokedex",
			command: commandExit,
			args:    0,
		},
		"map": {
			name:    "map",
			desc:    "Lists the next set of areas",
			command: commandMap,
			args:    0,
		},
		"mapb": {
			name:    "mapb",
			desc:    "Lists the previous set of areas",
			command: commandMapB,
			args:    0,
		},
		"explore": {
			name:    "explore",
			desc:    "Searches an area for Pokemon",
			command: commandExplore,
			args:    1,
		},
	}
}
