package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type PokeApi struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) PokeApi {
	return PokeApi{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
