package main

import (
	"time"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func main() {
	pokedexClient := apicaller.NewClient(5 * time.Second)

	cfg := &config{
		client: pokedexClient,
	}

	startRepl(cfg)
}
