package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/agreen254/pokedexcli/internal/api"
)

func commandCatch(params cmdParams) error {
	_, ok := params.cfg.pokedex[params.arg]
	if ok {
		return fmt.Errorf("already caught %s", params.arg)
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", params.arg)

	path := "/pokemon/" + params.arg
	data, err := requestOrCache(params.cache, path, api.CatchRequest)
	if err != nil {
		return err
	}

	// build some suspense
	time.Sleep(500 * time.Millisecond)
	fmt.Println("...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("...")
	time.Sleep(500 * time.Millisecond)

	if willCatch(data.BaseExperience) {
		fmt.Println(params.arg, "was caught!")
		cfg.pokedex[params.arg] = data
	} else {
		fmt.Println(params.arg, "escaped!")
	}

	var _ = data
	return nil
}

// average base experience is roughly 100, so normal pokemon will be caught reliably
// and pseudolegendaries/legendaries will be more difficult
// e.g. pikachu has a base experience of 112 and moltres/articuno have a base experience of 290
func willCatch(baseExp int) bool {
	return rand.IntN(baseExp) <= 100
}
