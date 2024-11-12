package pokeapi

import (
	"net/http"
	"time"

	"github.com/sam-maton/go-pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type PokeApi struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration) PokeApi {
	return PokeApi{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(10 * time.Second),
	}
}
