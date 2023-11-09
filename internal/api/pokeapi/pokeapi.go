package pokeapi

import (
	"net/http"
	"time"

	"github.com/bkohler93/pokedexcli/internal/cache"
)

const (
	baseUrl = "https://pokeapi.co/api/v2/"
)

type Client struct {
	client *http.Client
	cache  *cache.Cache
}

func NewPokeApiClient(timeout time.Duration, reapInterval time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(reapInterval),
	}
}
