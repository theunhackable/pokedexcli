package models

import cache "github.com/theunhackable/pokedexcli/internal/pokedexcache"

type Config struct {
	NextUrl  *string
	PrevUrl  *string
	ApiCache *cache.Cache
	Param    *string
}
