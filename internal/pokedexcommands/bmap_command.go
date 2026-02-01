package pokedexcommands

import (
	"fmt"

	models "github.com/theunhackable/pokedexcli/internal/models"
	api "github.com/theunhackable/pokedexcli/internal/pokedexapi"
)

func getBMap(config *models.Config) ([]models.Results, error) {
	if config.PrevUrl == nil {
		return nil, fmt.Errorf("You're on the first page")
	}
	response, err := api.GetResponse[models.AreaResponse](*config.PrevUrl, config)

	if err != nil {
		return nil, err
	}
	config.PrevUrl = response.Previous
	config.NextUrl = response.Next
	return response.Results, nil
}

func commandBMap(config *models.Config) error {

	locations, err := getBMap(config)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
