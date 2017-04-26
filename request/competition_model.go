package request

// CompetitionCountry wraps single Competition w/ Country
/*
{
  "id": 66,
  "cup": false,
  "name": "Premiership",
  "active": true,
  "country": {
    "id": 23,
    "name": "Scotland",
    "iso_code": "SCO",
    "flag": "https://soccerama.pro/images/countries//23.png"
  }
}
*/
type CompetitionCountry struct {
	ID     int    `json:"id"`
	Cup    bool   `json:"cup"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// CompetitionsCountry wraps a list of Competitions w/ Country
/*
{
  "data": [
    {
      "id": 43,
      "cup": false,
      "name": "Superliga",
      "active": true,
      "country": {
        "id": 13,
        "name": "Denmark",
        "iso_code": "DNK",
        "flag": "https://soccerama.pro/images/countries//13.png"
      }
    },
    {
      "id": 66,
      "cup": false,
      "name": "Premiership",
      "active": true,
      "country": {
        "id": 23,
        "name": "Scotland",
        "iso_code": "SCO",
        "flag": "https://soccerama.pro/images/countries//23.png"
      }
    }
  ]
}
*/
type CompetitionsCountry struct {
	Data []struct {
		ID     int    `json:"id"`
		Cup    bool   `json:"cup"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"data"`
}
