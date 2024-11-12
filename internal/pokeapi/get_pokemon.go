package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (p *PokeApi) GetPokemon(name string) (PokemonResult, error) {

	url := baseURL + "/pokemon/" + name

	var result PokemonResult

	if val, ok := p.cache.Get(url); ok {
		json.Unmarshal(val, &result)
		return result, nil
	}

	resp, err := http.Get(url)

	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	apiJson, err := io.ReadAll(resp.Body)

	if err != nil {
		return result, err
	}

	json.Unmarshal(apiJson, &result)

	return result, nil
}
