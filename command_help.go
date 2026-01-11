package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, v := range getCommands() {
		fmt.Printf("%v: \t %v \n", v.name, v.description)
	}

	return nil
}
