package poke_api

import (
	"net/http"
	"time"

	"github.com/vemolista/pokedex-cli/internal/poke_cache"
)

const BaseUrl = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      poke_cache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: poke_cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
