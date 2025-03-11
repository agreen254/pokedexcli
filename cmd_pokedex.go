package main

import (
	"fmt"
)

func commandPokedex(params cmdParams) error {
	if len(params.cfg.pokedex) == 0 {
		fmt.Println("Your pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range params.cfg.pokedex {
		fmt.Println(" -", pokemon)
	}
	return nil
}
