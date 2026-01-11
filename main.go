package main

import (
	"time"

	"github.com/natnael-alemayehu/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeapiClient,
	}

	startRepl(cfg)
}
