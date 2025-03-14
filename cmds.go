package main

import (
	"strings"

	"github.com/agreen254/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(params cmdParams) error
}

type cmdParams struct {
	cfg   *config
	cache *pokecache.Cache
	arg   string
}

func getCmd(input string) string {
	fields := strings.Fields(strings.ToLower(input))
	if len(fields) == 0 {
		return ""
	} else {
		return fields[0]
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the specified pokemon. Example: catch jolteon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "Returns the list of possible encounters within an area. Example: explore celadon-city-area",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Returns a list of properties for the given pokemon. Inspect not work unless you have caught the pokemon. Example: inspect pidgey",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of locations in the pokemon world. Successive calls will show more pages of names.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Short for map-back. Displays the previous page of location names.",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List out all pokemon you have caught. Gotta catch em' all!",
			callback:    commandPokedex,
		},
	}
}
