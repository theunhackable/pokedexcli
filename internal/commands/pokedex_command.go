package commands

import (
	"fmt"

	"github.com/theunhackable/pokedexcli/internal/models"
)

func commandPokedex(config *models.Config) error {
	if len(config.Pokemons) == 0 {
		return fmt.Errorf("You don't have any captured pokemon.")
	}
	fmt.Println("Your Pokedex:")
	for name, _ := range config.Pokemons {
		fmt.Printf("  - %s\n", name)
	}
	return nil

}
