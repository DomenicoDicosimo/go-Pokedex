package pokeapi

import (
	"net/http"
	"time"

	"github.com/DomenicoDicosimo/go-Pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
