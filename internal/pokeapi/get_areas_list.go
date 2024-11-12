package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (p *PokeApi) GetAreas(pageURL *string) (AreaResult, error) {

	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := p.cache.Get(url); ok {
		apiResult := AreaResult{}
		json.Unmarshal(val, &apiResult)
		return apiResult, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return AreaResult{}, err
	}

	defer resp.Body.Close()

	var apiResult AreaResult

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return AreaResult{}, err
	}

	json.Unmarshal(data, &apiResult)

	p.cache.Add(url, data)
	return apiResult, nil
}
