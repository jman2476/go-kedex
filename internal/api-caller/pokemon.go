package apicaller

type Pokemon struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}

type PokemonEncounters struct {
	PokemonList []Pokemon `json:"pokemon_encounters"`
	Location    Area      `json:"location"`
}
