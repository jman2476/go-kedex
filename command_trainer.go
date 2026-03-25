package main

import "fmt"

func commandTrainer(c *config, name string) error {
	if name == "" {
		return fmt.Errorf("Give your trainer a name")
	}

	c.trainer = name

	fmt.Printf("Trainer set to %s\n", name)

	return nil
}
