package request

// CompetitionCountry model wraps a single competition w/ country
type CompetitionCountry struct {
	ID      int    `json:"id"`
	Cup     bool   `json:"cup"`
	Name    string `json:"name"`
	Active  bool   `json:"active"`
	Country struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		IsoCode string `json:"iso_code"`
		Flag    string `json:"flag"`
	} `json:"country"`
}

// CompetitionsCountry model wraps a list of competitions w/ country
type CompetitionsCountry struct {
	Data []struct {
		ID      int    `json:"id"`
		Cup     bool   `json:"cup"`
		Name    string `json:"name"`
		Active  bool   `json:"active"`
		Country struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			IsoCode string `json:"iso_code"`
			Flag    string `json:"flag"`
		} `json:"country"`
	} `json:"data"`
}
