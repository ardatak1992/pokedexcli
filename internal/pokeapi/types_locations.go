package pokeapi

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type LocationAreaResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonStatResponse struct {
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	BaseExperience int `json:"base_experience"`
	Height         int `json:"height"`
	Weight         int `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Pokemon struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	BaseExperience int
	Height         int `json:"height"`
	Weight         int `json:"weight"`
	Stats          []struct {
		BaseStat int
		Stat     struct {
			Name string
			URL  string
		}
	}
	Types []struct {
		Type struct {
			Name string
			URL  string
		}
	}
}
