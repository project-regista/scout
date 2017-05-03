package request

// SeasonCompetition model wraps a single season w/ competition
type SeasonCompetition struct {
	ID            int    `json:"id"`
	CompetitionID int    `json:"competition_id"`
	Name          string `json:"name"`
	Active        bool   `json:"active"`
	Competition   struct {
		ID     int    `json:"id"`
		Cup    bool   `json:"cup"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"competition"`
}

// SeasonsCompetition model wraps a list of seasons w/ competition
type SeasonsCompetition struct {
	Data []struct {
		ID            int    `json:"id"`
		CompetitionID int    `json:"competition_id"`
		Name          string `json:"name"`
		Active        bool   `json:"active"`
		Competition   struct {
			ID     int    `json:"id"`
			Cup    bool   `json:"cup"`
			Name   string `json:"name"`
			Active bool   `json:"active"`
		} `json:"competition"`
	} `json:"data"`
}
