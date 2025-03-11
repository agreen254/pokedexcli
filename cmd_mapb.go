package main

import (
	"errors"
	"fmt"

	"github.com/agreen254/pokedexcli/internal/api"
)

func commandMapb(params cmdParams) error {
	next := params.cfg.mapCommand.next
	prev := params.cfg.mapCommand.prev

	if prev == nil && next != nil {
		return errors.New("you're on the first page")
	}

	var path string
	if next == nil && prev == nil {
		return errors.New("no pages yet")
	} else {
		path = lastDir(*cfg.mapCommand.prev)
	}

	data, err := requestOrCache(params.cache, path, api.MapRequest)
	if err != nil {
		return err
	}

	for _, item := range data.Results {
		fmt.Println(item.Name)
	}

	// update the config with the next/prev returned via the response
	params.cfg.mapCommand.next = data.Next
	params.cfg.mapCommand.prev = data.Previous

	return nil
}
