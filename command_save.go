package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func commandSave(c *config, trainer string) error {
	if trainer == "" && c.Trainer == "" {
		return fmt.Errorf("Use the 'trainer' command to name your trainer before saving your progress")
	}

	filename := trainer + ".json"
	if c.Trainer != "" {
		filename = c.Trainer + ".json"
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}

	fmt.Println("Save successful")
	return nil
}
