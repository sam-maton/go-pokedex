package main

import "github.com/sam-maton/go-pokedex/internal/pokeapi"

type Config struct {
	locationPrev *string
	locationNext *string
	API          pokeapi.PokeApi
	args         []string
	pokemon      map[string]pokeapi.PokemonResult
}

type CLICommand struct {
	name    string
	desc    string
	command func(config *Config) error
	args    int
}
