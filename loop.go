package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sam-maton/go-pokedex/internal/pokeapi"
)

func startLoop() {
	basicConfig := &Config{
		API:     pokeapi.NewClient(5 * time.Second),
		pokemon: make(map[string]pokeapi.PokemonResult),
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	//User input loop
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		command := splitCommand(scanner.Text())
		basicConfig.args = command
		commandFunc, exists := getCommands()[command[0]]

		if commandFunc.args != len(command[1:]) {
			fmt.Printf("%s requires %d args \n", commandFunc.name, commandFunc.args)
			continue
		}

		if !exists {
			fmt.Println("Command does not exist: " + scanner.Text())
			continue
		}

		err := commandFunc.command(basicConfig)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func splitCommand(command string) []string {
	return strings.Split(command, " ")
}
