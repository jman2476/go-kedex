package main

import (
	"fmt"

	apicaller "github.com/jman2476/go-kedex/internal/api-caller"
)

func commandMap(c *config) error {
	// get next 20 map locations
	var locationAreas apicaller.LocationAreas
	locationAreas, err := c.client.GetMaps(c.Next)

	if err != nil {
		return fmt.Errorf("Error getting maps: %w", err)
	}

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(c *config) error {
	// get previous 20 map locations
	if c.Previous == nil {
		fmt.Println("You are on the first page of results")
		return nil
	}
	var locationAreas apicaller.LocationAreas
	locationAreas, err := c.client.GetMaps(c.Previous)

	if err != nil {
		return fmt.Errorf("Error getting maps: %w", err)
	}

	c.Next = locationAreas.Next
	c.Previous = locationAreas.Previous

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
