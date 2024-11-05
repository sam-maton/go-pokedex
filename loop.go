package main

import (
	"bufio"
	"fmt"
	"os"
)

func startLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		command, exists := getCommands()[scanner.Text()]

		if !exists {
			fmt.Println("Command does not exist.")
			continue
		}

		command.command()
	}
}

type cliCommand struct {
	name    string
	desc    string
	command func() error
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
	}
}

func commandHelp() error {
	fmt.Println("Help text here...")
	return nil
}

func commandExit() error {
	os.Exit(0)

	return nil
}
