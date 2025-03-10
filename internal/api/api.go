package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_PATH = "https://pokeapi.co/api/v2/"

// MakeRequest is a wrapper for http GET requests to the PokeAPI.
// It adds the path arg to 'https://pokeapi.co/api/v2/' and will attempt
// to decode any response into the expected arg.
func MakeRequest(path string, expected any) error {
	fullPath := BASE_PATH + path
	res, err := http.Get(fullPath)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&expected)
	if err != nil {
		return fmt.Errorf("failed to decode body: %w", err)
	}

	return nil
}
