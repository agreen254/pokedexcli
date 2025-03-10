package main

import (
	"fmt"

	"github.com/agreen254/pokedexcli/internal/pokecache"
)

func commandHelp(*config, *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
