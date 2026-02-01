package pokedexcommands

import (
	"fmt"

	"github.com/theunhackable/pokedexcli/internal/models"
	api "github.com/theunhackable/pokedexcli/internal/pokedexapi"
)

func getMap(config *models.Config) ([]models.Results, error) {
	if config.NextUrl == nil {
		return nil, fmt.Errorf("Your'e on the last page")
	}
	response, err := api.GetResponse[models.AreaResponse](*config.NextUrl)
	if err != nil {
		return nil, err
	}
	config.PrevUrl = response.Previous
	config.NextUrl = response.Next
	return response.Results, nil
}

func commandMap(config *models.Config) error {
	locations, err := getMap(config)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
