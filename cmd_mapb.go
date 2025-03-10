package main

import (
	"errors"
	"fmt"

	"github.com/agreen254/pokedexcli/internal/api"
	"github.com/agreen254/pokedexcli/internal/pokecache"
)

func commandMapb(cfg *config, cache *pokecache.Cache) error {
	if cfg.mapCommand.prev == nil && cfg.mapCommand.next != nil {
		return errors.New("you're on the first page")
	}

	var path string
	if cfg.mapCommand.next == nil && cfg.mapCommand.prev == nil {
		return errors.New("no pages yet")
	} else {
		path = lastDir(*cfg.mapCommand.prev)
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
