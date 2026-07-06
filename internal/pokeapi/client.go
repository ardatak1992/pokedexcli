package pokeapi

import (
	"net/http"
	"time"

	"github.com/ardatak1992/pokedexcli/internal/pokecache"
)

type Pokedex struct {
	pokemon map[string]Pokemon
}

type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(30 * time.Minute),
	}
}
