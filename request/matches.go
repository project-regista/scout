package request

import (
	"encoding/json"
	"fmt"

	scout "github.com/project-regista/scout/client"
)

// GetMatchSeason retrieve information about a single match w/ season information
func GetMatchSeason(client scout.Client, id string) (MatchSeason, error) {

	requestURL := fmt.Sprintf("https://%s/%s/matches/%s",
		client.APIHost, client.APIVersion, id)

	// Request options
	params := map[string]string{
		"api_token": client.APIToken,
		"include":   "season,homeTeam,awayTeam,venue,events,lineup,homeStats,awayStats,commentaries,odds,referee",
	}

	// Make HTTP GET request
	body, err := Get(requestURL, params)
	if err != nil {
		return MatchSeason{}, fmt.Errorf("Failed to make match-season request: %s", err)
	}

	// Decode the HTTP response into a Competition struct
	defer body.Close()
	dec := json.NewDecoder(body)

	var matchSeason MatchSeason

	if err := dec.Decode(&matchSeason); err != nil {
		return MatchSeason{}, fmt.Errorf("Failed to decode match-season response: %s", err)
	}
	return matchSeason, nil
}
