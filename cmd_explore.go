package main

import (
	"fmt"

	"github.com/agreen254/pokedexcli/internal/api"
)

func commandExplore(params cmdParams) error {
	path := "/location-area/" + params.arg
	fmt.Println("Exploring", params.arg, "...")

	data, err := requestOrCache(params.cache, path, api.ExploreRequest)
	if err != nil {
		return err
	}
	if len(data.PokemonEncounters) == 0 {
		return fmt.Errorf("no pokemon found")
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range data.PokemonEncounters {
		fmt.Println(" -", pokemon.Pokemon.Name)
	}

	return nil
}
