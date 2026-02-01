package pokedexapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/theunhackable/pokedexcli/internal/models"
)

func GetResponse[T any](url string, config *models.Config) (T, error) {
	var response T
	if config == nil {
		return response, fmt.Errorf("No configuration is provided.")
	}

	if val, ok := config.ApiCache.Get(url); ok {
		err := json.Unmarshal(val, &response)
		return response, err
	}

	res, err := http.Get(url)
	if err != nil {
		return response, fmt.Errorf("Error fetching map: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return response, fmt.Errorf("returned status %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return response, err
	}

	config.ApiCache.Add(url, body)

	err = json.Unmarshal(body, &response)

	return response, err
}
