package commands

import (
	"fmt"
	"math/rand"

	"github.com/theunhackable/pokedexcli/internal/api"
	"github.com/theunhackable/pokedexcli/internal/models"
)

func getPokemonData(name string, config *models.Config) (models.PokemonNameResponse, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	res, err := api.GetResponse[models.PokemonNameResponse](url, config)
	if err != nil {
		return models.PokemonNameResponse{}, err
	}
	return res, nil
}
func commandCatch(config *models.Config) error {

	if config.Param == nil {
		return fmt.Errorf("Pokemon name NOT provided.")
	}

	name := *config.Param

	res, err := getPokemonData(name, config)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if res.BaseExperience <= rand.Intn(500) {
		fmt.Printf("%s was caught!\n", name)

		config.Pokemons[name] = res
		fmt.Println("You may now inspect it with the inspect command.")

	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil

}
