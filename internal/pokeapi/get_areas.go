package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreas(pageURL *string) (LocationResponse, error) {

	url := apiURL + locationEndpoint
	if pageURL != nil {
		url = *pageURL
	}
	var locations LocationResponse
	cachedItem, ok := c.pokeCache.Get(url)
	if ok {
		err := json.Unmarshal(cachedItem, &locations)
		if err != nil {
			return LocationResponse{}, fmt.Errorf("Error getting cached item: %v", err)
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("Error sending request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("Error getting response: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("Error reading data: %v", err)
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("Error unmarshaling data: %v", err)
	}

	return locations, nil

}
