package pokedexcommands

import (
	"fmt"

	"github.com/theunhackable/pokedexcli/internal/models"
	api "github.com/theunhackable/pokedexcli/internal/pokedexapi"
)

func getPokemonsInArea(url string) ([]string, error) {

	res, err := api.GetResponse[models.PokemonInArea](url)
	if err != nil {
		return nil, err
	}
	pokemons := make([]string, 0)
	for _, entry := range res.PokemonEncounters {
		pokemons = append(pokemons, entry.Pokemon.Name)
	}
	return pokemons, nil
}

func commandExplore(config *models.Config) error {
	if config.Param == nil {
		return fmt.Errorf("Invalid params")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + *config.Param
	pokemons, err := getPokemonsInArea(url)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s area...\n", *config.Param)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}
