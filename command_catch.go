package main

import (
	"fmt"
	"math/rand"
	"slices"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func commandCatch(c *config, pokemon string) error {
	if c.Location.Name == "" {
		return fmt.Errorf("You have to explore a location first")
	}

	if pokemon == "" {
		return fmt.Errorf("What pokemon do you want to catch?")
	}

	if !slices.Contains(c.Location.Pokemon, pokemon) {
		return fmt.Errorf("%s isn't found in %s", pokemon, c.Location.Name)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	var pokemonInfo apicaller.PokemonInfo
	pokemonInfo, err := c.client.GetPokemonData(pokemon)
	if err != nil {
		return fmt.Errorf("Error catching pokemon: %w", err)
	}

	caught := capture(pokemonInfo)

	if caught {
		c.Pokedex[pokemon] = pokemonInfo
		fmt.Printf("Successfully caught %s!\n", pokemon)
	} else {
		fmt.Printf("%s got away...\n", pokemon)
	}

	return nil
}

func capture(pokemon apicaller.PokemonInfo) bool {
	scaleFactor := rand.Intn(10)
	return scaleFactor*pokemon.BaseExperience < 1000
}
