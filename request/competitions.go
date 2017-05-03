package request

import (
	"encoding/json"
	"log"

	"fmt"

	scout "github.com/project-regista/scout/client"
)

// GetCompetitionCountry get information about a single competition w/ country
func GetCompetitionCountry(client scout.Client, id string) (CompetitionCountry, error) {

	requestURL := fmt.Sprintf("https://%s/%s/competitions/%s",
		client.APIHost, client.APIVersion, id)

	// Request options
	params := map[string]string{
		"api_token": client.APIToken,
		"include":   "country",
	}

	// Make HTTP GET request
	body, err := Get(requestURL, params)
	if err != nil {
		return CompetitionCountry{}, fmt.Errorf("Failed to make competition-country request: %s", err)
	}

	// Decode the HTTP response into a Competition struct
	defer body.Close()
	dec := json.NewDecoder(body)

	var competitionCountry CompetitionCountry

	if err := dec.Decode(&competitionCountry); err != nil {
		return CompetitionCountry{}, fmt.Errorf("Failed to decode competition-country response: %s", err)
	}
	log.Printf("%+v\n", competitionCountry)
	return competitionCountry, nil
}

// GetCompetitionsCountry get all the competitions w/ countries
func GetCompetitionsCountry(client scout.Client) (CompetitionsCountry, error) {

	requestURL := fmt.Sprintf("https://%s/%s/competitions",
		client.APIHost, client.APIVersion)

	params := map[string]string{
		"api_token": client.APIToken,
		"include":   "country",
	}

	body, err := Get(requestURL, params)
	if err != nil {
		return CompetitionsCountry{}, fmt.Errorf("Failed to make competitions-country request: %s", err)
	}

	defer body.Close()
	dec := json.NewDecoder(body)

	var competitionsCountry CompetitionsCountry

	if err := dec.Decode(&competitionsCountry); err != nil {
		return CompetitionsCountry{}, fmt.Errorf("Failed to decode competitions-country response: %s", err)
	}
	log.Printf("%+v\n", competitionsCountry)
	return competitionsCountry, nil
}
