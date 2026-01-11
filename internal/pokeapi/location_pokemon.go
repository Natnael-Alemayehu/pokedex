package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInLocation(args ...string) (RespLolcationPokemon, error) {

	url := baseURL + "/location-area/" + args[0]

	if data, ok := c.cache.Get(url); ok {
		var locationpokemon RespLolcationPokemon
		if err := json.Unmarshal(data, &locationpokemon); err != nil {
			return RespLolcationPokemon{}, err
		}
		return locationpokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLolcationPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLolcationPokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLolcationPokemon{}, err
	}

	c.cache.Add(url, data)

	var locationpokemon RespLolcationPokemon
	if err := json.Unmarshal(data, &locationpokemon); err != nil {
		return RespLolcationPokemon{}, err
	}

	return locationpokemon, nil

}
