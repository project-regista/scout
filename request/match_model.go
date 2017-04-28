package request

// MatchSeason wraps information about a single match w/ season information
/*
{
  "id": 795379,
  "weather": null,
  "temperature": null,
  "commentary": true,
  "ht_score": null,
  "ft_score": null,
  "et_score": null,
  "home_team_id": 520,
  "away_team_id": 631,
  "home_score": 0,
  "away_score": 0,
  "home_score_penalties": 0,
  "away_score_penalties": 0,
  "formation": {
    "home": null,
    "away": null
  },
  "date_time_tba": 0,
  "spectators": null,
  "starting_date": "2017-04-28",
  "starting_time": "16:00:00",
  "status": "NS",
  "minute": 0,
  "injury_time": 0,
  "extra_time": 0,
  "competition_id": 43,
  "venue_id": 7640,
  "season_id": 662,
  "round_id": 10228,
  "stage_id": 1578,
  "referee_id": null,
  "aggregate": null,
  "placeholder": false,
  "deleted": false,
  "result_only": false,
  "season": {
    "id": 662,
    "competition_id": 43,
    "name": "2016/2017",
    "active": true
  }
}
*/
type MatchSeason struct {
	ID                 int         `json:"id"`
	Weather            interface{} `json:"weather"`
	Temperature        interface{} `json:"temperature"`
	Commentary         bool        `json:"commentary"`
	HtScore            interface{} `json:"ht_score"`
	FtScore            interface{} `json:"ft_score"`
	EtScore            interface{} `json:"et_score"`
	HomeTeamID         int         `json:"home_team_id"`
	AwayTeamID         int         `json:"away_team_id"`
	HomeScore          int         `json:"home_score"`
	AwayScore          int         `json:"away_score"`
	HomeScorePenalties int         `json:"home_score_penalties"`
	AwayScorePenalties int         `json:"away_score_penalties"`
	Formation          struct {
		Home interface{} `json:"home"`
		Away interface{} `json:"away"`
	} `json:"formation"`
	DateTimeTba   int         `json:"date_time_tba"`
	Spectators    interface{} `json:"spectators"`
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
	RefereeID     interface{} `json:"referee_id"`
	Aggregate     interface{} `json:"aggregate"`
	Placeholder   bool        `json:"placeholder"`
	Deleted       bool        `json:"deleted"`
	ResultOnly    bool        `json:"result_only"`
	Season        struct {
		ID            int    `json:"id"`
		CompetitionID int    `json:"competition_id"`
		Name          string `json:"name"`
		Active        bool   `json:"active"`
	} `json:"season"`
}

// MatchesSeason wraps all the matches for you plan that are played between the from and to date w/ season information
/*
{
  "data": [
    {
      "id": 795378,
      "weather": null,
      "temperature": "",
      "commentary": true,
      "ht_score": "2-0",
      "ft_score": "4-0",
      "et_score": null,
      "home_team_id": 629,
      "away_team_id": 173,
      "home_score": 4,
      "away_score": 0,
      "home_score_penalties": 0,
      "away_score_penalties": 0,
      "formation": {
        "home": "4-4-2",
        "away": "4-4-2"
      },
      "date_time_tba": 0,
      "spectators": 8547,
      "starting_date": "2017-04-24",
      "starting_time": "17:00:00",
      "status": "FT",
      "minute": 90,
      "injury_time": 0,
      "extra_time": 0,
      "competition_id": 43,
      "venue_id": 2424,
      "season_id": 662,
      "round_id": 10227,
      "stage_id": 1578,
      "referee_id": 374,
      "aggregate": null,
      "placeholder": false,
      "deleted": false,
      "result_only": false,
      "season": {
        "id": 662,
        "competition_id": 43,
        "name": "2016/2017",
        "active": true
      }
    }
  ],
  "meta": {
    "pagination": {
      "total": 1,
      "count": 1,
      "per_page": 500,
      "current_page": 1,
      "total_pages": 1,
      "links": []
    }
  }
}
*/
type MatchesSeason struct {
	Data []struct {
		ID                 int         `json:"id"`
		Weather            interface{} `json:"weather"`
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
		Season        struct {
			ID            int    `json:"id"`
			CompetitionID int    `json:"competition_id"`
			Name          string `json:"name"`
			Active        bool   `json:"active"`
		} `json:"season"`
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
