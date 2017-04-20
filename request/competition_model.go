package request

// Competition wraps single Competition
/*
{
  "id": 66,
  "cup": false,
  "name": "Premiership",
  "active": true
}
*/
type Competition struct {
	ID     int    `json:"id"`
	Cup    bool   `json:"cup"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// Competitions wraps a list of Competitions
/*
{
  "data": [
    {
      "id": 43,
      "cup": false,
      "name": "Superliga",
      "active": true
    },
    {
      "id": 66,
      "cup": false,
      "name": "Premiership",
      "active": true
    }
  ]
}
*/
type Competitions struct {
	Data []struct {
		ID     int    `json:"id"`
		Cup    bool   `json:"cup"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"data"`
}

// CompetitionsSeasons wraps a list of competitions with seasons (current season included)
/*
{
  "data": [
    {
      "id": 43,
      "cup": false,
      "name": "Superliga",
      "active": true,
      "currentSeason": {
        "id": 662,
        "competition_id": 43,
        "name": "2016/2017",
        "active": true
      },
      "seasons": {
        "data": [
          {
            "id": 350,
            "competition_id": 43,
            "name": "2015/2016",
            "active": true
          },
          {
            "id": 662,
            "competition_id": 43,
            "name": "2016/2017",
            "active": true
          }
        ]
      }
    },
    {
      "id": 66,
      "cup": false,
      "name": "Premiership",
      "active": true,
      "currentSeason": {
        "id": 741,
        "competition_id": 66,
        "name": "2016/2017",
        "active": true
      },
      "seasons": {
        "data": [
          {
            "id": 355,
            "competition_id": 66,
            "name": "2015/2016",
            "active": true
          },
          {
            "id": 741,
            "competition_id": 66,
            "name": "2016/2017",
            "active": true
          }
        ]
      }
    }
  ]
}
*/
type CompetitionsSeasons struct {
	Data []struct {
		ID            int    `json:"id"`
		Cup           bool   `json:"cup"`
		Name          string `json:"name"`
		Active        bool   `json:"active"`
		CurrentSeason struct {
			ID            int    `json:"id"`
			CompetitionID int    `json:"competition_id"`
			Name          string `json:"name"`
			Active        bool   `json:"active"`
		} `json:"currentSeason"`
		Seasons struct {
			Data []struct {
				ID            int    `json:"id"`
				CompetitionID int    `json:"competition_id"`
				Name          string `json:"name"`
				Active        bool   `json:"active"`
			} `json:"data"`
		} `json:"seasons"`
	} `json:"data"`
}
