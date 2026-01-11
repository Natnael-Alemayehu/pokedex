package main

import "fmt"

func commandPokedox(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.caughtPokemon {
		fmt.Printf("  -%v\n", v.Name)
	}
	return nil
}
