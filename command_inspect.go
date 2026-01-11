package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you have to provide a pokemon name")
	}

	val, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", val.Name)
	fmt.Printf("Weight: %v\n", val.Weight)
	fmt.Printf("Height: %v\n", val.Height)
	fmt.Println("Stats:")
	// for k,v := range val.Stats{
	// 	fmt.Printf("\t-")
	// }
	fmt.Println("Types:")
	for _, v := range val.Types {
		fmt.Printf("\t-%v\n", v.Type.Name)
	}

	return nil
}
