package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func commandLoad(c *config, trainer string) error {
	filename := trainer + ".json"
	if c.Trainer != "" {
		filename = c.Trainer + ".json"
	}

	if filename == ".txt" {
		return fmt.Errorf("What trainer do you want to load?")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully loaded %s\n", filename)
	return nil
}
