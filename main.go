package main

import (
	"time"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func main() {
	pokedexClient := apicaller.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		client: pokedexClient,
	}

	startRepl(cfg)
}
