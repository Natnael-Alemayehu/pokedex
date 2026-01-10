package pokeapi

import (
	"net/http"
	"time"
)

// Client
type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		http.Client{Timeout: timeout},
	}
}
