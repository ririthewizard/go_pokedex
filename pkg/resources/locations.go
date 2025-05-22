package resources

type LocationArea struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	GameIndex int      `json:"game_index"`
	Location  Location `json:"location"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
