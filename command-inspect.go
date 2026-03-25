package main

import (
	"fmt"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func commandInspect(c *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("What pokemon do you want to inspect?")
	}

	data, ok := c.Pokedex[pokemon]
	if !ok {
		fmt.Printf("You have not caught %s\n", pokemon)
		return nil
	}

	printStats(data)
	return nil
}

func printStats(pokemon apicaller.PokemonInfo) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Base experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Stats:\n")
	for idx, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d", stat.Stat.Name, stat.BaseStat)
		if idx%3 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
	fmt.Printf("Abilities:\n")
	for idx, ability := range pokemon.Abilities {
		fmt.Printf(" -%s", ability.Ability.Name)
		if idx%3 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
	fmt.Printf("Moves:\n")
	for idx, move := range pokemon.Moves {
		fmt.Printf(" -%s", move.Move.Name)
		if idx%3 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
	fmt.Printf("Types:\n")
	for idx, pokeType := range pokemon.Types {
		fmt.Printf(" -%s", pokeType.Type.Name)
		if idx%3 == 0 {
			fmt.Print("\n")
		}
	}
}
