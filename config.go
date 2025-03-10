package main

type config struct {
	mapCommand mapCommandParams
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
}
