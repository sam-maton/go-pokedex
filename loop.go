package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	prev string
	next string
}

func startLoop() {
	basicConfig := config{
		prev: "next test",
		next: "next test",
	}
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

		command.command(&basicConfig)
	}
}
