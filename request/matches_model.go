package request

// MatchSeason model wraps a single match w/ season
type MatchSeason struct {
	ID      int `json:"id"`
	Weather struct {
		Code        string `json:"code"`
		Type        string `json:"type"`
		Icon        string `json:"icon"`
		Temperature struct {
			Temp float64 `json:"temp"`
			Unit string  `json:"unit"`
		} `json:"temperature"`
		Clouds   string `json:"clouds"`
		Humidity string `json:"humidity"`
		Wind     struct {
			Speed  string `json:"speed"`
			Degree int    `json:"degree"`
		} `json:"wind"`
	} `json:"weather"`
	Temperature        string `json:"temperature"`
	Commentary         bool   `json:"commentary"`
	HtScore            string `json:"ht_score"`
	FtScore            string `json:"ft_score"`
	EtScore            string `json:"et_score"`
	HomeTeamID         int    `json:"home_team_id"`
	AwayTeamID         int    `json:"away_team_id"`
	HomeScore          int    `json:"home_score"`
	AwayScore          int    `json:"away_score"`
	HomeScorePenalties int    `json:"home_score_penalties"`
	AwayScorePenalties int    `json:"away_score_penalties"`
	Formation          struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"formation"`
	DateTimeTba   int    `json:"date_time_tba"`
	Spectators    int    `json:"spectators"`
	StartingDate  string `json:"starting_date"`
	StartingTime  string `json:"starting_time"`
	Status        string `json:"status"`
	Minute        int    `json:"minute"`
	InjuryTime    int    `json:"injury_time"`
	ExtraTime     int    `json:"extra_time"`
	CompetitionID int    `json:"competition_id"`
	VenueID       int    `json:"venue_id"`
	SeasonID      int    `json:"season_id"`
	RoundID       int    `json:"round_id"`
	StageID       int    `json:"stage_id"`
	Placeholder   bool   `json:"placeholder"`
	Deleted       bool   `json:"deleted"`
	ResultOnly    bool   `json:"result_only"`
	Season        struct {
		ID            int    `json:"id"`
		CompetitionID int    `json:"competition_id"`
		Name          string `json:"name"`
		Active        bool   `json:"active"`
	} `json:"season"`
}

// MatchesSeason model wraps a list of matches w/ season
type MatchesSeason struct {
	Data []struct {
		ID      int `json:"id"`
		Weather struct {
			Code        string `json:"code"`
			Type        string `json:"type"`
			Icon        string `json:"icon"`
			Temperature struct {
				Temp float64 `json:"temp"`
				Unit string  `json:"unit"`
			} `json:"temperature"`
			Clouds   string `json:"clouds"`
			Humidity string `json:"humidity"`
			Wind     struct {
				Speed  string `json:"speed"`
				Degree int    `json:"degree"`
			} `json:"wind"`
		} `json:"weather"`
		Temperature        string `json:"temperature"`
		Commentary         bool   `json:"commentary"`
		HtScore            string `json:"ht_score"`
		FtScore            string `json:"ft_score"`
		EtScore            string `json:"et_score"`
		HomeTeamID         int    `json:"home_team_id"`
		AwayTeamID         int    `json:"away_team_id"`
		HomeScore          int    `json:"home_score"`
		AwayScore          int    `json:"away_score"`
		HomeScorePenalties int    `json:"home_score_penalties"`
		AwayScorePenalties int    `json:"away_score_penalties"`
		Formation          struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"formation"`
		DateTimeTba   int    `json:"date_time_tba"`
		Spectators    int    `json:"spectators"`
		StartingDate  string `json:"starting_date"`
		StartingTime  string `json:"starting_time"`
		Status        string `json:"status"`
		Minute        int    `json:"minute"`
		InjuryTime    int    `json:"injury_time"`
		ExtraTime     int    `json:"extra_time"`
		CompetitionID int    `json:"competition_id"`
		VenueID       int    `json:"venue_id"`
		SeasonID      int    `json:"season_id"`
		RoundID       int    `json:"round_id"`
		StageID       int    `json:"stage_id"`
		RefereeID     int    `json:"referee_id"`
		Aggregate     int    `json:"aggregate"`
		Placeholder   bool   `json:"placeholder"`
		Deleted       bool   `json:"deleted"`
		ResultOnly    bool   `json:"result_only"`
		Season        struct {
			ID            int    `json:"id"`
			CompetitionID int    `json:"competition_id"`
			Name          string `json:"name"`
			Active        bool   `json:"active"`
		} `json:"season"`
	} `json:"data"`
}
