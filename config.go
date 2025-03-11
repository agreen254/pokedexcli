package main

type config struct {
	mapCommand mapCommandParams
	pokedex    Pokedex
}

type mapCommandParams struct {
	next *string
	prev *string
}

var cfg = config{
	mapCommand: mapCommandParams{
		next: nil,
		prev: nil,
	},
	pokedex: Pokedex{},
}
