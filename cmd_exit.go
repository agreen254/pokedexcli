package main

import (
	"fmt"
	"os"

	"github.com/agreen254/pokedexcli/internal/pokecache"
)

func commandExit(*config, *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
