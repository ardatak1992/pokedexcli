package pokeapi

import (
	"net/http"
	"time"

	"github.com/ardatak1992/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient    http.Client
	locationCache *pokecache.Cache
	pokemonCache  *pokecache.Cache
	Pokedex       map[string]Pokemon
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		locationCache: pokecache.NewCache(10 * time.Second),
		pokemonCache:  pokecache.NewCache(10 * time.Second),
		Pokedex:       map[string]Pokemon{},
	}
}
