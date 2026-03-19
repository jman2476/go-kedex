package apicaller

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int
	Next     *string
	Previous *string
	Results  []Area
}

type Area struct {
	Name string
	Url  string
}

func (c *Client) GetMaps(url *string) (LocationAreas, error) {
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	var locations LocationAreas
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationAreas{}, err
	}

	return locations, nil
}
