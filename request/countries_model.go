package request

// Country model wraps a single country
type Country struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IsoCode string `json:"iso_code"`
	Flag    string `json:"flag"`
}

// Countries model wraps a list of countries
type Countries struct {
	Data []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		IsoCode string `json:"iso_code"`
		Flag    string `json:"flag"`
	} `json:"data"`
}
