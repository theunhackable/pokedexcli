package main

import (
	"github.com/theunhackable/pokedexcli/internal/cache"
	"github.com/theunhackable/pokedexcli/internal/client"
	"github.com/theunhackable/pokedexcli/internal/models"
	"time"
)

func main() {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	var NextUrl *string = &url
	var PrevUrl *string
	var ApiCache = cache.NewCache(10 * time.Second)
	var Pokemons = make(map[string]models.PokemonNameResponse)
	config := models.Config{
		NextUrl:  NextUrl,
		PrevUrl:  PrevUrl,
		ApiCache: ApiCache,
		Pokemons: Pokemons,
	}
	client.StartRepl(&config)
}
