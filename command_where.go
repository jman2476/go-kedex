package main

import "fmt"

func commandWhere(c *config, _ string) error {
	if c.Location.Name == "" {
		return fmt.Errorf("You haven't explored anywhere yet")
	}

	fmt.Printf("You are in %s\n", c.Location.Name)

	return nil
}
