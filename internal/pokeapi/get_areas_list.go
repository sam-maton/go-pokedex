package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetAreas(pokiApi PokeApi) (APIResult, error) {
	resp, err := http.Get(baseURL + "/location-area")
	if err != nil {
		return APIResult{}, err
	}

	defer resp.Body.Close()

	var apiResult APIResult

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&apiResult)

	if err != nil {
		return APIResult{}, err
	}

	return apiResult, nil
}
