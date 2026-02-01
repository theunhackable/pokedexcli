package main

import (
	models "github.com/theunhackable/pokedexcli/internal/models"
	cache "github.com/theunhackable/pokedexcli/internal/pokedexcache"
	client "github.com/theunhackable/pokedexcli/internal/pokedexclient"
	"time"
)

func main() {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	var NextUrl *string = &url
	var PrevUrl *string
	var ApiCache = cache.NewCache(10 * time.Second)

	config := models.Config{
		NextUrl:  NextUrl,
		PrevUrl:  PrevUrl,
		ApiCache: ApiCache,
	}

	client.StartRepl(&config)
}
