package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaInfo(areaName *string) (Location, error) {
	url := apiURL + locationEndpoint + *areaName
	var loc Location
	if cachedItem, ok := c.pokeCache.Get(url); ok {
		err := json.Unmarshal(cachedItem, &loc)
		if err != nil {
			return Location{}, fmt.Errorf("Error getting cached item: %v", err)
		}
		return loc, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, fmt.Errorf("Error creating request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, fmt.Errorf("Error getting response: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Can't find area name: %s\n", *areaName)
		return Location{}, nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("Error reading data: %v", err)
	}

	c.pokeCache.Add(url, data)

	err = json.Unmarshal(data, &loc)
	if err != nil {
		return Location{}, fmt.Errorf("Error unmarshaling data: %v", err)
	}

	return loc, nil
}
