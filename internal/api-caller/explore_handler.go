package apicaller

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(location string) (PokemonEncounters, error) {
	data, ok := c.cache.Get(location)

	if !ok {
		url := baseURL + "/location-area/" + location
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonEncounters{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonEncounters{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonEncounters{}, err
		}
	}

	var encounters PokemonEncounters
	err := json.Unmarshal(data, &encounters)
	if err != nil {
		return PokemonEncounters{}, err
	}
	// fmt.Printf("%v\n", encounters)

	c.cache.Add(location, data)
	return encounters, nil
}
