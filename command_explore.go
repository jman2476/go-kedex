package main

import (
	"fmt"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func commandExplore(c *config, location string) error {
	if location == "" {
		return fmt.Errorf("Need location to explore")
	}

	fmt.Printf("Exploring area %s\n", location)

	var encounters apicaller.PokemonEncounters
	encounters, err := c.client.GetPokemon(location)
	if err != nil {
		return fmt.Errorf("Error getting pokemon: %w", err)
	}

	for _, encounter := range encounters.PokemonList {
		fmt.Printf("%v\n", encounter.Pokemon.Name)
	}

	return nil
}
