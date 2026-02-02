package commands

import (
	"fmt"

	"github.com/theunhackable/pokedexcli/internal/models"
)

func commandInspect(config *models.Config) error {

	if config.Param == nil {
		return fmt.Errorf("Pokemon name not provided.")
	}
	name := *config.Param

	entry, ok := config.Pokemons[name]

	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Height: %v\n", entry.Height)
	fmt.Printf("Weight: %v\n", entry.Weight)
	fmt.Println("Stats:")
	for _, stat := range entry.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, entry := range entry.Types {
		fmt.Printf("  - %s\n", entry.Type.Name)
	}

	return nil
}
