package main

import "fmt"

func commandPokedex(c *config, trainer string) error {
	if len(c.Pokedex) == 0 {
		return fmt.Errorf("You haven't caught any pokemon! You gotta catch 'em all!")
	}

	for _, pokemon := range c.Pokedex {
		entry := fmt.Sprintf(" - %s", pokemon.Name)
		fmt.Println(entry)
	}

	return nil
}
