package request

// SeasonCompetition get information about a single season w/ competition
/*
{
  "id": 350,
  "competition_id": 43,
  "name": "2015/2016",
  "active": true,
  "competition": {
    "id": 43,
    "cup": false,
    "name": "Superliga",
    "active": true
  }
}
*/
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

// SeasonsCompetition get all the seasons w/ competitions
/*
{
  "data": [
    {
      "id": 350,
      "competition_id": 43,
      "name": "2015/2016",
      "active": true,
      "competition": {
        "id": 43,
        "cup": false,
        "name": "Superliga",
        "active": true
      }
    },
    {
      "id": 662,
      "competition_id": 43,
      "name": "2016/2017",
      "active": true,
      "competition": {
        "id": 43,
        "cup": false,
        "name": "Superliga",
        "active": true
      }
    },
    {
      "id": 355,
      "competition_id": 66,
      "name": "2015/2016",
      "active": true,
      "competition": {
        "id": 66,
        "cup": false,
        "name": "Premiership",
        "active": true
      }
    },
    {
      "id": 741,
      "competition_id": 66,
      "name": "2016/2017",
      "active": true,
      "competition": {
        "id": 66,
        "cup": false,
        "name": "Premiership",
        "active": true
      }
    }
  ],
  "meta": {
    "pagination": {
      "total": 4,
      "count": 4,
      "per_page": 50,
      "current_page": 1,
      "total_pages": 1,
      "links": []
    }
  }
}
*/
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
