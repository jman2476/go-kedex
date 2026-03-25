package main

import "fmt"

func commandTrainer(c *config, name string) error {
	if name == "" {
		if c.Trainer != "" {
			return fmt.Errorf("Current trainer is %s", c.Trainer)
		}
		return fmt.Errorf("Give your trainer a name")
	}

	c.Trainer = name

	fmt.Printf("Trainer set to %s\n", name)

	return nil
}
