package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (p *PokeApi) GetAreas(pageURL *string) (APIResult, error) {

	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := p.cache.Get(url); ok {
		apiResult := APIResult{}
		json.Unmarshal(val, &apiResult)
		return apiResult, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return APIResult{}, err
	}

	defer resp.Body.Close()

	var apiResult APIResult

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return APIResult{}, err
	}

	json.Unmarshal(data, &apiResult)

	p.cache.Add(url, data)
	return apiResult, nil
}
