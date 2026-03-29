package main

import (
	"fmt"
	"slices"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func commandExplore(c *config, location string) error {

	if location == "" {
		if c.Location.Name == "" && len(c.MapArea) == 0 {
			return fmt.Errorf("Need location to explore")
		}
		location = exploreAreas(c)
		location = c.Location.Name
	}

	fmt.Printf("Exploring area %s\n", location)

	var encounters apicaller.PokemonEncounters
	encounters, err := c.client.GetPokemon(location)
	if err != nil {
		return fmt.Errorf("Error getting pokemon: %w", err)
	}
	c.Location.Name = location

	for _, encounter := range encounters.PokemonList {
		fmt.Printf("%v\n", encounter.Pokemon.Name)
		c.Location.Pokemon = append(c.Location.Pokemon, encounter.Pokemon.Name)
	}

	return nil
}

func exploreAreas(c *config) string {
	index := slices.Index(c.MapArea, c.Location.Name)
	if index == -1 {
		index = 0
	}
	// numAreas = len(c.MapArea)
	location, err := scrollRepl(c.MapArea, index)
	if err != nil {
		fmt.Printf("Error found: %w\n", err)
		return ""
	}

	// for idx, area := range c.MapArea {
	// 	margin := "  "
	// 	if idx == index {
	// 		margin = " >"
	// 	}
	// 	fmt.Printf("%s%s\n", margin, area)
	// }

	return location
}
