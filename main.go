package main

import (
	"time"
	"fmt"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func main() {
	pokedexClient := apicaller.NewClient(5*time.Second, 5*time.Minute)

	cfg := &config{
		client:  pokedexClient,
		Pokedex: map[string]apicaller.PokemonInfo{},
	}

	fmt.Println("Loading...")
	fmt.Print("Getting Pokemons")
	time.Sleep(8 * time.Second)
	fmt.Println("\r\033[K\033[A\033[K\rReady!            ")

	startRepl(cfg)
}
