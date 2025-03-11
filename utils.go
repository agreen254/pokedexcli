package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/agreen254/pokedexcli/internal/api"
	"github.com/agreen254/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	fields := strings.Fields(text)

	lowered := make([]string, len(fields))
	for i, field := range fields {
		lowered[i] = strings.ToLower(field)
	}
	return lowered
}

func lastDir(path string) string {
	slice := strings.Split(path, "/")
	return slice[len(slice)-1]
}

func requestOrCache[T any](cache *pokecache.Cache, path string, handlerFunc func(string) (T, error)) (T, error) {
	var data T
	entry, ok := cache.Get(api.BASE_PATH + path)

	if ok {
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return data, fmt.Errorf("error accessing cache: %w", err)
		}
	} else {
		res, err := handlerFunc(path)
		if err != nil {
			return data, err
		}

		cacheData, err := json.Marshal(res)
		if err != nil {
			return data, fmt.Errorf("error adding data to cache: %w", err)
		}

		cache.Add(api.BASE_PATH+path, cacheData)
		data = res
	}

	return data, nil
}
