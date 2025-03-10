package main

import (
	"errors"
	"fmt"

	"github.com/agreen254/pokedexcli/internal/api"
	"github.com/agreen254/pokedexcli/internal/pokecache"
)

func commandMap(cfg *config, cache *pokecache.Cache) error {
	// if a response has been returned with a valid next and nil prev
	// it means there are no more pages remaining
	if cfg.mapCommand.next == nil && cfg.mapCommand.prev != nil {
		return errors.New("last page reached")
	}

	var path string
	if cfg.mapCommand.next == nil && cfg.mapCommand.prev == nil {
		// if map command has not been used yet it will have neither prev nor next
		path = "location-area?offset=0&limit=20"
	} else {
		path = lastDir(*cfg.mapCommand.next)
	}

	data, err := requestAndCache(cache, path, api.MakeRequestMap)
	if err != nil {
		return err
	}

	for _, item := range data.Results {
		fmt.Println(item.Name)
	}

	// update the config with the next/prev returned via the response
	cfg.mapCommand.next = data.Next
	cfg.mapCommand.prev = data.Previous

	return nil
}
