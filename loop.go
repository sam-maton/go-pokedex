package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/sam-maton/go-pokedex/internal/pokeapi"
)

type Config struct {
	Prev string
	Next string
	API  pokeapi.PokeApi
}

func startLoop() {
	basicConfig := Config{
		Prev: "",
		Next: "https://pokeapi.co/api/v2/location-area",
		API:  pokeapi.NewClient(5 * time.Second),
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

		err := command.command(&basicConfig)

		if err != nil {
			fmt.Println(err)
		}
	}
}
