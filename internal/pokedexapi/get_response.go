package pokedexapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetResponse[T any](url string) (T, error) {
	var response T

	res, err := http.Get(url)
	if err != nil {
		return response, fmt.Errorf("Error fetching map: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return response, fmt.Errorf("returned status %d", res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return response, fmt.Errorf("Error while decoding the body %w", err)
	}
	return response, nil
}
