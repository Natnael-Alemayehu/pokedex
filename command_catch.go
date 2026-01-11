package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you have to provide a pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%v was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
