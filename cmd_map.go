package main

import (
	"errors"
	"fmt"

	"github.com/agreen254/pokedexcli/internal/api"
)

func commandMap(params cmdParams) error {
	next := params.cfg.mapCommand.next
	prev := params.cfg.mapCommand.prev

	// if a response has been returned with a valid next and nil prev
	// it means there are no more pages remaining
	if next == nil && prev != nil {
		return errors.New("last page reached")
	}

	var path string
	if next == nil && prev == nil {
		// if map command has not been used yet it will have neither prev nor next
		path = "location-area?offset=0&limit=20"
	} else {
		path = lastDir(*cfg.mapCommand.next)
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
