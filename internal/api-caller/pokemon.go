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
	ID             int    `json:"id"`
	BaseExperience int    `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	} `json:"stats"`
	Height     int  `json:"height"`
	Weight     int  `json:"weight"`
	Is_default bool `json:"is_default"`
	Abilities  []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
