package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(cityName string) ([]Pokemon, error) {
	url := baseURL + "/location-area/" + cityName
	var dat []byte
	entry, ok := c.pokemonCache.Get(url)
	if ok {
		dat = entry
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		c.pokemonCache.Add(url, dat)
	}

	var apiResp LocationAreaResponse
	err := json.Unmarshal(dat, &apiResp)
	if err != nil {
		return nil, err
	}

	pokemonArr := make([]Pokemon, 0, len(apiResp.PokemonEncounters))
	for _, pe := range apiResp.PokemonEncounters {
		pokemonArr = append(pokemonArr, Pokemon{
			Name: pe.Pokemon.Name,
			Url:  pe.Pokemon.URL,
		})
	}

	return pokemonArr, nil
}
