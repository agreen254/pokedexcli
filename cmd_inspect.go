package main

import (
	"fmt"
)

func commandInspect(params cmdParams) error {
	pokemon, ok := params.cfg.pokedex[params.arg]
	if !ok {
		return fmt.Errorf("you have not caught %s", params.arg)
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		str := fmt.Sprintf(" -%s: %d", stat.Stat.Name, stat.BaseStat)
		fmt.Println(str)
	}

	fmt.Println("Types:")
	for _, tp := range pokemon.Types {
		fmt.Println(" -", tp.Type.Name)
	}

	return nil
}
