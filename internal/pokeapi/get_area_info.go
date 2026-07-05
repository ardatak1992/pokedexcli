package pokeapi

func (c *Client) GetAreaInfo(areaName *string) (Location, error) {
	url := apiURL + locationEndpoint + *areaName
	if val, ok := c.pokeCache.Get(url); ok {
		c.pokeCache.Add()
	}
}
