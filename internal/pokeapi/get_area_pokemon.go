package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (p *PokeApi) GetAreaPokemon(area string) (AreaPokemonResult, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := p.cache.Get(url); ok {
		apiResult := AreaPokemonResult{}
		json.Unmarshal(val, &apiResult)
		return apiResult, nil
	}

	var result AreaPokemonResult
	resp, err := http.Get(url)

	if err != nil {
		return AreaPokemonResult{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaPokemonResult{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return AreaPokemonResult{}, err
	}

	p.cache.Add(url, body)

	return result, nil
}
