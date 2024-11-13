package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

func commandCatch(config *Config) error {
	result, err := config.API.GetPokemon(config.args[1])

	if err != nil {
		return err
	}

	if result.Experience == 0 && result.Name == "" {
		return errors.New("pokemon could not be found")
	}
	fmt.Println("Throwing a Pokeball at " + result.Name + "...")

	caught := rand.Intn(result.Experience)
	if caught > 50 {
		fmt.Println(result.Name + " escaped!")
		return nil
	}

	fmt.Println(result.Name + " was caught!")

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
		"catch": {
			name:    "catch",
			desc:    "Try and catch a Pokemon",
			command: commandCatch,
			args:    1,
		},
	}
}
