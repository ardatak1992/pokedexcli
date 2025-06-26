package pokeapi

import (
	"encoding/json"
	"io"

	"net/http"
)

func (c *Client) ListStats(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	var dat []byte

	entry, ok := c.pokemonCache.Get(url)
	if ok {
		dat = entry
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, nil
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.pokemonCache.Add(url, dat)

	}

	var parsed PokemonStatResponse
	if err := json.Unmarshal(dat, &parsed); err != nil {
		return Pokemon{}, err
	}

	return Pokemon{
		parsed.Forms[0].Name,
		parsed.Forms[0].URL,
		parsed.BaseExperience,
		parsed.Height,
		parsed.Weight,
		[]struct {
			BaseStat int
			Stat     struct {
				Name string
				URL  string
			}
		}(parsed.Stats),
		[]struct {
			Type struct {
				Name string
				URL  string
			}
		}(parsed.Types),
	}, nil

}
