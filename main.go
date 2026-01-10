package main

import (
	"time"

	"github.com/natnael-alemayehu/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeapiClient,
	}

	startRepl(cfg)
}
