package request

import (
	"encoding/json"
	"fmt"

	scout "github.com/project-regista/scout/client"
)

// GetSeasonCompetition get information about a single season w/ competition
func GetSeasonCompetition(client scout.Client, id string) (SeasonCompetition, error) {

	requestURL := fmt.Sprintf("https://%s/%s/seasons/%s",
		client.APIHost, client.APIVersion, id)

	// Request options
	params := map[string]string{
		"api_token": client.APIToken,
		"include":   "competition",
	}

	// Make HTTP GET request
	body, err := Get(requestURL, params)
	if err != nil {
		return SeasonCompetition{}, fmt.Errorf("Failed to make season-competition request: %s", err)
	}

	// Decode the HTTP response into a Competition struct
	defer body.Close()
	dec := json.NewDecoder(body)

	var seasonCompetition SeasonCompetition

	if err := dec.Decode(&seasonCompetition); err != nil {
		return SeasonCompetition{}, fmt.Errorf("Failed to decode season-competition response: %s", err)
	}
	return seasonCompetition, nil
}

// GetSeasonsCompetition get all the seasons w/ competitions
func GetSeasonsCompetition(client scout.Client) (SeasonsCompetition, error) {

	requestURL := fmt.Sprintf("https://%s/%s/seasons",
		client.APIHost, client.APIVersion)

	params := map[string]string{
		"api_token": client.APIToken,
		"include":   "competition",
	}

	body, err := Get(requestURL, params)
	if err != nil {
		return SeasonsCompetition{}, fmt.Errorf("Failed to make competitions-country request: %s", err)
	}

	defer body.Close()
	dec := json.NewDecoder(body)

	var seasonsCompetition SeasonsCompetition

	if err := dec.Decode(&seasonsCompetition); err != nil {
		return SeasonsCompetition{}, fmt.Errorf("Failed to decode competitions-country response: %s", err)
	}
	return seasonsCompetition, nil
}
