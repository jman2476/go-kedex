package apicaller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Next     string
	Previous string
	Results  []Area
}

type Area struct {
	Name string
	Url  string
}

func GetMaps(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("Error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("Error reading response: %w", err)
	}

	var locations LocationAreas
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreas{}, fmt.Errorf("Error unmarshaling data: %w", err)
	}

	return locations, nil
}
