package request

// Country wraps a single country
/*
{
  "id": 13,
  "name": "Denmark",
  "iso_code": "DNK",
  "flag": "https://soccerama.pro/images/countries//13.png"
}
*/
type Country struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IsoCode string `json:"iso_code"`
	Flag    string `json:"flag"`
}

// Countries wraps a list of countries
/*
{
  "data": [
    {
      "id": 13,
      "name": "Denmark",
      "iso_code": "DNK",
      "flag": "https://soccerama.pro/images/countries//13.png"
    },
    {
      "id": 23,
      "name": "Scotland",
      "iso_code": "SCO",
      "flag": "https://soccerama.pro/images/countries//23.png"
    }
  ]
}
*/
type Countries struct {
	Data []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		IsoCode string `json:"iso_code"`
		Flag    string `json:"flag"`
	} `json:"data"`
}

// CountriesCompetitions holds supported countries and its competitions
/*
{
  "data": [
    {
      "id": 13,
      "name": "Denmark",
      "iso_code": "DNK",
      "flag": "https://soccerama.pro/images/countries//13.png"
    },
    {
      "id": 23,
      "name": "Scotland",
      "iso_code": "SCO",
      "flag": "https://soccerama.pro/images/countries//23.png"
    }
  ]
}
*/
type CountriesCompetitions struct {
	Data []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		IsoCode string `json:"iso_code"`
		Flag    string `json:"flag"`
	} `json:"data"`
}
