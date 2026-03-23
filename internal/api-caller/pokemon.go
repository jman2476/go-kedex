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

type PokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Stats          []Stat `json:"stats"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}
