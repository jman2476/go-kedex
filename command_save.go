package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func commandSave(c *config, trainer string) error {
	filename := trainer + ".txt"
	if c.trainer != "" {
		filename = c.trainer + ".txt"
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
