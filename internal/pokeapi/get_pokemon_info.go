package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName *string) (Pokemon, error) {
	url := apiURL + pokemonEndpoint + *pokemonName
	var pokemon Pokemon
	if cachedItem, ok := c.pokeCache.Get(url); ok {
		err := json.Unmarshal(cachedItem, &pokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error getting cached item: %v", err)
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error creating request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error getting response: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Can't find pokemon: %s\n", *pokemonName)
		return Pokemon{}, nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error reading data: %v", err)
	}

	c.pokeCache.Add(url, data)

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error unmarshalling data: %v", err)
	}

	return pokemon, nil
}
