package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {

	pokemonlocation, err := cfg.pokeapiClient.ListPokemonInLocation(args...)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")

	for _, v := range pokemonlocation.PokemonEncounters {
		fmt.Printf("\t- %v \n", v.Pokemon.Name)
	}

	return nil
}
