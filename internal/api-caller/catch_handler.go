package apicaller

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(pokemon string) (PokemonInfo, error) {
	data, ok := c.cache.Get(pokemon)

	if !ok {
		url := baseURL + "/pokemon/" + pokemon
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonInfo{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonInfo{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonInfo{}, err
		}
	}

	var pokemonInfo PokemonInfo
	err := json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, nil
	}

	c.cache.Add(pokemon, data)
	return pokemonInfo, nil
}
