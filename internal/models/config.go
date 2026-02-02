package models

import "github.com/theunhackable/pokedexcli/internal/cache"

type Config struct {
	NextUrl  *string
	PrevUrl  *string
	ApiCache *cache.Cache
	Param    *string
	Pokemons map[string]PokemonNameResponse
}
