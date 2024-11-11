package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (p *PokeApi) GetAreas(pageURL *string) (APIResult, error) {

	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	resp, err := http.Get(url)
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
