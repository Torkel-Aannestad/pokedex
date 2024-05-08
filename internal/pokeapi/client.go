package pokeapi

import (
	"net/http"
	"time"

	"github.com/torkelaannestad/pokedex/internal/pokecache"
)

type Client struct {
	HttpClient http.Client
	Cache      pokecache.Cache
}

func NewClient(timeout time.Duration, cacheDuration time.Duration) Client {
	return Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(cacheDuration),
	}
}
