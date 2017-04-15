// Package competition contains functions to make
// use of the Soccerama 'Competitions' endpoint
package competition

import (
	"encoding/json"

	"fmt"

	scout "github.com/project-regista/scout/client"
	"github.com/project-regista/scout/request"
)

// GetCompetition get a competition by ID
func GetCompetition(client scout.Client, id string) (Competition, error) {

	requestURL := fmt.Sprintf("https://%s/%s/competitions/%s",
		client.APIHost, client.APIVersion, id)

	// Request options
	params := map[string]string{"api_token": client.APIToken}

	// Make HTTP GET request
	body, err := request.Get(requestURL, params)
	if err != nil {
		return Competition{}, fmt.Errorf("Failed to make competition request: %s", err)
	}

	// Decode the HTTP reponse into a Competition struct
	defer body.Close()
	dec := json.NewDecoder(body)

	var competition Competition

	if err := dec.Decode(&competition); err != nil {
		return Competition{}, fmt.Errorf("Failed to decode competition response: %s", err)
	}
	return competition, nil
}

// GetCompetitions get a list of competitions
func GetCompetitions(client scout.Client) (Competitions, error) {

	requestURL := fmt.Sprintf("https://%s/%s/competitions",
		client.APIHost, client.APIVersion)

	params := map[string]string{"api_token": client.APIToken}

	body, err := request.Get(requestURL, params)
	if err != nil {
		return Competitions{}, fmt.Errorf("Failed to make competitions request: %s", err)
	}

	defer body.Close()
	dec := json.NewDecoder(body)

	var competitions Competitions

	if err := dec.Decode(&competitions); err != nil {
		return Competitions{}, fmt.Errorf("Failed to decode competitions response: %s", err)
	}
	return competitions, nil
}

// GetCompetitionsSeasons get a list of competitions w/ seasons
func GetCompetitionsSeasons(client scout.Client) (CompetitionsSeasons, error) {

	requestURL := fmt.Sprintf("https://%s/%s/competitions",
		client.APIHost, client.APIVersion)

	params := map[string]string{
		"api_token": client.APIToken,
		"include":   "currentSeason,seasons",
	}

	body, err := request.Get(requestURL, params)
	if err != nil {
		return CompetitionsSeasons{}, fmt.Errorf("Failed to make competitions request: %s", err)
	}

	defer body.Close()
	dec := json.NewDecoder(body)

	var competitionsSeasons CompetitionsSeasons

	if err := dec.Decode(&competitionsSeasons); err != nil {
		return CompetitionsSeasons{}, fmt.Errorf("Failed to decode competitions response: %s", err)
	}
	return competitionsSeasons, nil
}
