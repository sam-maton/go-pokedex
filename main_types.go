package main

import "github.com/sam-maton/go-pokedex/internal/pokeapi"

type Config struct {
	locationPrev *string
	locationNext *string
	API          pokeapi.PokeApi
	args         []string
}

type CLICommand struct {
	name    string
	desc    string
	command func(config *Config) error
	args    int
}
