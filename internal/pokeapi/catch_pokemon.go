package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(args ...string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + args[0]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	var resppokemon RespPokemon
	if err := json.Unmarshal(data, &resppokemon); err != nil {
		return RespPokemon{}, err
	}

	return resppokemon, nil
}
