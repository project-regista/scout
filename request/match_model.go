package request

// MatchSeason wraps information about a single match w/ season information
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
	Temperature        string      `json:"temperature"`
	Commentary         bool        `json:"commentary"`
	HtScore            string      `json:"ht_score"`
	FtScore            string      `json:"ft_score"`
	EtScore            interface{} `json:"et_score"`
	HomeTeamID         int         `json:"home_team_id"`
	AwayTeamID         int         `json:"away_team_id"`
	HomeScore          int         `json:"home_score"`
	AwayScore          int         `json:"away_score"`
	HomeScorePenalties int         `json:"home_score_penalties"`
	AwayScorePenalties int         `json:"away_score_penalties"`
	Formation          struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"formation"`
	DateTimeTba   int         `json:"date_time_tba"`
	Spectators    int         `json:"spectators"`
	StartingDate  string      `json:"starting_date"`
	StartingTime  string      `json:"starting_time"`
	Status        string      `json:"status"`
	Minute        int         `json:"minute"`
	InjuryTime    int         `json:"injury_time"`
	ExtraTime     int         `json:"extra_time"`
	CompetitionID int         `json:"competition_id"`
	VenueID       int         `json:"venue_id"`
	SeasonID      int         `json:"season_id"`
	RoundID       int         `json:"round_id"`
	StageID       int         `json:"stage_id"`
	RefereeID     int         `json:"referee_id"`
	Aggregate     interface{} `json:"aggregate"`
	Placeholder   bool        `json:"placeholder"`
	Deleted       bool        `json:"deleted"`
	ResultOnly    bool        `json:"result_only"`
	HomeTeam      struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Logo       string `json:"logo"`
		Twitter    string `json:"twitter"`
		VenueID    int    `json:"venue_id"`
		CoachID    int    `json:"coach_id"`
		ChairmanID int    `json:"chairman_id"`
	} `json:"homeTeam"`
	AwayTeam struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Logo       string `json:"logo"`
		Twitter    string `json:"twitter"`
		VenueID    int    `json:"venue_id"`
		CoachID    int    `json:"coach_id"`
		ChairmanID int    `json:"chairman_id"`
	} `json:"awayTeam"`
	Venue struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		City              string      `json:"city"`
		Address           interface{} `json:"address"`
		Phone             interface{} `json:"phone"`
		Fax               interface{} `json:"fax"`
		YearOfManufacture interface{} `json:"year_of_manufacture"`
		Seats             interface{} `json:"seats"`
		FieldType         interface{} `json:"field_type"`
	} `json:"venue"`
	Events struct {
		Data []struct {
			ID               int         `json:"id"`
			MatchID          int         `json:"match_id"`
			TeamID           int         `json:"team_id"`
			Minute           int         `json:"minute"`
			ExtraMin         interface{} `json:"extra_min"`
			Type             string      `json:"type"`
			PlayerInID       int         `json:"player_in_id,omitempty"`
			PlayerInName     string      `json:"player_in_name,omitempty"`
			PlayerOutID      int         `json:"player_out_id,omitempty"`
			PlayerOutName    string      `json:"player_out_name,omitempty"`
			PlayerID         int         `json:"player_id,omitempty"`
			PlayerName       string      `json:"player_name,omitempty"`
			AssistID         interface{} `json:"assist_id,omitempty"`
			AssistPlayerName interface{} `json:"assist_player_name,omitempty"`
		} `json:"data"`
	} `json:"events"`
	Lineup struct {
		Data []struct {
			MatchID         int    `json:"match_id"`
			TeamID          int    `json:"team_id"`
			PlayerID        int    `json:"player_id"`
			PlayerName      string `json:"player_name"`
			Position        string `json:"position"`
			ShirtNumber     int    `json:"shirt_number"`
			Assists         int    `json:"assists"`
			FoulsCommited   int    `json:"fouls_commited"`
			FoulsDrawn      int    `json:"fouls_drawn"`
			Goals           int    `json:"goals"`
			Offsides        int    `json:"offsides"`
			MissedPenalties int    `json:"missed_penalties"`
			ScoredPenalties int    `json:"scored_penalties"`
			Posx            int    `json:"posx"`
			Posy            int    `json:"posy"`
			Redcards        int    `json:"redcards"`
			Saves           int    `json:"saves"`
			ShotsOnGoal     int    `json:"shots_on_goal"`
			ShotsTotal      int    `json:"shots_total"`
			Yellowcards     int    `json:"yellowcards"`
			Type            string `json:"type"`
		} `json:"data"`
	} `json:"lineup"`
	HomeStats struct {
		ID             int         `json:"id"`
		TeamID         int         `json:"team_id"`
		ShotsOffTarget int         `json:"shots_off_target"`
		ShotsOnTarget  int         `json:"shots_on_target"`
		ShotsTotal     int         `json:"shots_total"`
		FoulsTotal     int         `json:"fouls_total"`
		CornersTotal   int         `json:"corners_total"`
		OffsidesTotal  interface{} `json:"offsides_total"`
		Possesion      int         `json:"possesion"`
		Yellowcards    int         `json:"yellowcards"`
		Redcards       int         `json:"redcards"`
		Saves          int         `json:"saves"`
		ThrowInTotal   interface{} `json:"throw_in_total"`
	} `json:"homeStats"`
	AwayStats struct {
		ID             int         `json:"id"`
		TeamID         int         `json:"team_id"`
		ShotsOffTarget int         `json:"shots_off_target"`
		ShotsOnTarget  int         `json:"shots_on_target"`
		ShotsTotal     int         `json:"shots_total"`
		FoulsTotal     int         `json:"fouls_total"`
		CornersTotal   int         `json:"corners_total"`
		OffsidesTotal  interface{} `json:"offsides_total"`
		Possesion      int         `json:"possesion"`
		Yellowcards    int         `json:"yellowcards"`
		Redcards       int         `json:"redcards"`
		Saves          int         `json:"saves"`
		ThrowInTotal   interface{} `json:"throw_in_total"`
	} `json:"awayStats"`
	Season struct {
		ID            int    `json:"id"`
		CompetitionID int    `json:"competition_id"`
		Name          string `json:"name"`
		Active        bool   `json:"active"`
	} `json:"season"`
	Commentaries struct {
		Data []struct {
			ID          int         `json:"id"`
			Body        string      `json:"body"`
			IsGoal      int         `json:"is_goal"`
			Important   int         `json:"important"`
			Minute      string      `json:"minute"`
			ExtraMinute interface{} `json:"extra_minute"`
			TeamID      int         `json:"team_id"`
		} `json:"data"`
	} `json:"commentaries"`
	Odds struct {
		Data []struct {
			BookmakerID int    `json:"bookmaker_id"`
			Name        string `json:"name"`
			Types       struct {
				Data []struct {
					Type string `json:"type"`
					Odds struct {
						Data []struct {
							Label      string `json:"label"`
							Value      string `json:"value"`
							PrevValue  string `json:"prev_value"`
							LastUpdate string `json:"last_update"`
							Winning    bool   `json:"winning"`
						} `json:"data"`
					} `json:"odds"`
				} `json:"data"`
			} `json:"types"`
		} `json:"data"`
	} `json:"odds"`
	Referee struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"referee"`
}

// MatchesSeason wraps all the matches for you plan that are played between the from and to date w/ season information
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
		Temperature        string      `json:"temperature"`
		Commentary         bool        `json:"commentary"`
		HtScore            string      `json:"ht_score"`
		FtScore            string      `json:"ft_score"`
		EtScore            interface{} `json:"et_score"`
		HomeTeamID         int         `json:"home_team_id"`
		AwayTeamID         int         `json:"away_team_id"`
		HomeScore          int         `json:"home_score"`
		AwayScore          int         `json:"away_score"`
		HomeScorePenalties int         `json:"home_score_penalties"`
		AwayScorePenalties int         `json:"away_score_penalties"`
		Formation          struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"formation"`
		DateTimeTba   int         `json:"date_time_tba"`
		Spectators    int         `json:"spectators"`
		StartingDate  string      `json:"starting_date"`
		StartingTime  string      `json:"starting_time"`
		Status        string      `json:"status"`
		Minute        int         `json:"minute"`
		InjuryTime    int         `json:"injury_time"`
		ExtraTime     int         `json:"extra_time"`
		CompetitionID int         `json:"competition_id"`
		VenueID       int         `json:"venue_id"`
		SeasonID      int         `json:"season_id"`
		RoundID       int         `json:"round_id"`
		StageID       int         `json:"stage_id"`
		RefereeID     int         `json:"referee_id"`
		Aggregate     interface{} `json:"aggregate"`
		Placeholder   bool        `json:"placeholder"`
		Deleted       bool        `json:"deleted"`
		ResultOnly    bool        `json:"result_only"`
		HomeTeam      struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Logo       string `json:"logo"`
			Twitter    string `json:"twitter"`
			VenueID    int    `json:"venue_id"`
			CoachID    int    `json:"coach_id"`
			ChairmanID int    `json:"chairman_id"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Logo       string `json:"logo"`
			Twitter    string `json:"twitter"`
			VenueID    int    `json:"venue_id"`
			CoachID    int    `json:"coach_id"`
			ChairmanID int    `json:"chairman_id"`
		} `json:"awayTeam"`
		Venue struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			City              string      `json:"city"`
			Address           interface{} `json:"address"`
			Phone             interface{} `json:"phone"`
			Fax               interface{} `json:"fax"`
			YearOfManufacture interface{} `json:"year_of_manufacture"`
			Seats             interface{} `json:"seats"`
			FieldType         interface{} `json:"field_type"`
		} `json:"venue"`
		Events struct {
			Data []struct {
				ID               int         `json:"id"`
				MatchID          int         `json:"match_id"`
				TeamID           int         `json:"team_id"`
				Minute           int         `json:"minute"`
				ExtraMin         interface{} `json:"extra_min"`
				Type             string      `json:"type"`
				PlayerInID       int         `json:"player_in_id,omitempty"`
				PlayerInName     string      `json:"player_in_name,omitempty"`
				PlayerOutID      int         `json:"player_out_id,omitempty"`
				PlayerOutName    string      `json:"player_out_name,omitempty"`
				PlayerID         int         `json:"player_id,omitempty"`
				PlayerName       string      `json:"player_name,omitempty"`
				AssistID         interface{} `json:"assist_id,omitempty"`
				AssistPlayerName interface{} `json:"assist_player_name,omitempty"`
			} `json:"data"`
		} `json:"events"`
		Lineup struct {
			Data []struct {
				MatchID         int    `json:"match_id"`
				TeamID          int    `json:"team_id"`
				PlayerID        int    `json:"player_id"`
				PlayerName      string `json:"player_name"`
				Position        string `json:"position"`
				ShirtNumber     int    `json:"shirt_number"`
				Assists         int    `json:"assists"`
				FoulsCommited   int    `json:"fouls_commited"`
				FoulsDrawn      int    `json:"fouls_drawn"`
				Goals           int    `json:"goals"`
				Offsides        int    `json:"offsides"`
				MissedPenalties int    `json:"missed_penalties"`
				ScoredPenalties int    `json:"scored_penalties"`
				Posx            int    `json:"posx"`
				Posy            int    `json:"posy"`
				Redcards        int    `json:"redcards"`
				Saves           int    `json:"saves"`
				ShotsOnGoal     int    `json:"shots_on_goal"`
				ShotsTotal      int    `json:"shots_total"`
				Yellowcards     int    `json:"yellowcards"`
				Type            string `json:"type"`
			} `json:"data"`
		} `json:"lineup"`
		HomeStats struct {
			ID             int         `json:"id"`
			TeamID         int         `json:"team_id"`
			ShotsOffTarget int         `json:"shots_off_target"`
			ShotsOnTarget  int         `json:"shots_on_target"`
			ShotsTotal     int         `json:"shots_total"`
			FoulsTotal     int         `json:"fouls_total"`
			CornersTotal   int         `json:"corners_total"`
			OffsidesTotal  interface{} `json:"offsides_total"`
			Possesion      int         `json:"possesion"`
			Yellowcards    int         `json:"yellowcards"`
			Redcards       int         `json:"redcards"`
			Saves          int         `json:"saves"`
			ThrowInTotal   interface{} `json:"throw_in_total"`
		} `json:"homeStats,omitempty"`
		AwayStats struct {
			ID             int         `json:"id"`
			TeamID         int         `json:"team_id"`
			ShotsOffTarget int         `json:"shots_off_target"`
			ShotsOnTarget  int         `json:"shots_on_target"`
			ShotsTotal     int         `json:"shots_total"`
			FoulsTotal     int         `json:"fouls_total"`
			CornersTotal   int         `json:"corners_total"`
			OffsidesTotal  interface{} `json:"offsides_total"`
			Possesion      int         `json:"possesion"`
			Yellowcards    int         `json:"yellowcards"`
			Redcards       int         `json:"redcards"`
			Saves          int         `json:"saves"`
			ThrowInTotal   interface{} `json:"throw_in_total"`
		} `json:"awayStats,omitempty"`
		Season struct {
			ID            int    `json:"id"`
			CompetitionID int    `json:"competition_id"`
			Name          string `json:"name"`
			Active        bool   `json:"active"`
		} `json:"season"`
		Commentaries struct {
			Data []struct {
				ID          int         `json:"id"`
				Body        string      `json:"body"`
				IsGoal      int         `json:"is_goal"`
				Important   int         `json:"important"`
				Minute      string      `json:"minute"`
				ExtraMinute interface{} `json:"extra_minute"`
				TeamID      int         `json:"team_id"`
			} `json:"data"`
		} `json:"commentaries"`
		Odds struct {
			Data []struct {
				BookmakerID int    `json:"bookmaker_id"`
				Name        string `json:"name"`
				Types       struct {
					Data []struct {
						Type string `json:"type"`
						Odds struct {
							Data []struct {
								Label      string `json:"label"`
								Value      string `json:"value"`
								PrevValue  string `json:"prev_value"`
								LastUpdate string `json:"last_update"`
								Winning    bool   `json:"winning"`
							} `json:"data"`
						} `json:"odds"`
					} `json:"data"`
				} `json:"types"`
			} `json:"data"`
		} `json:"odds"`
		Referee struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"referee"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total       int           `json:"total"`
			Count       int           `json:"count"`
			PerPage     int           `json:"per_page"`
			CurrentPage int           `json:"current_page"`
			TotalPages  int           `json:"total_pages"`
			Links       []interface{} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}
