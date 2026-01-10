package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// List locations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	bte, ok := c.cache.Get(url)
	if ok {
		locationsResp := RespShallowLocations{}
		if err := json.Unmarshal(bte, &locationsResp); err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	locationsResp := RespShallowLocations{}
	if err = json.Unmarshal(dat, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
