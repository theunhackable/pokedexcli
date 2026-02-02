package commands

import (
	"fmt"
	"os"

	"github.com/theunhackable/pokedexcli/internal/models"
)

func commandExit(config *models.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
