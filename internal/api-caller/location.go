package apicaller

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Area
}

type Area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
