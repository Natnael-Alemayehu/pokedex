package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"

	fmt.Println(url)

	if c.Next != "" {
		url = c.Next
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error occured trying to fetch location-are: %v", err)
	}

	defer resp.Body.Close()

	var locationarea locationArea
	if err := json.NewDecoder(resp.Body).Decode(&locationarea); err != nil {
		return fmt.Errorf("error while decoding to locationarea: %v", err)
	}

	c.Next = locationarea.Next
	c.Previous = locationarea.Previous

	for _, v := range locationarea.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapPrevious(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"

	fmt.Println(url)

	if c.Previous != "" {
		url = c.Previous
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error occured trying to fetch location-are: %v", err)
	}

	defer resp.Body.Close()

	var locationarea locationArea
	if err := json.NewDecoder(resp.Body).Decode(&locationarea); err != nil {
		return fmt.Errorf("error while decoding to locationarea: %v", err)
	}

	c.Next = locationarea.Next
	c.Previous = locationarea.Previous

	for _, v := range locationarea.Results {
		fmt.Println(v.Name)
	}

	return nil
}
