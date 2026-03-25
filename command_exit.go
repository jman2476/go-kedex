package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, _ string) error {
	if c.Trainer != "" {
		fmt.Printf("Saving trainer %s's data...", c.Trainer)
		commandSave(c, c.Trainer)
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
