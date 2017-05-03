package request

import (
	"encoding/json"
	"fmt"
	"log"

	scout "github.com/project-regista/scout/client"
)

// GetMatch retrieve information about a single match w/ season information
func GetMatch(client scout.Client, id string) (Match, error) {

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
		return Match{}, fmt.Errorf("Failed to make match request: %s", err)
	}

	// Decode the HTTP response into a Competition struct
	defer body.Close()
	dec := json.NewDecoder(body)

	var match Match

	if err := dec.Decode(&match); err != nil {
		return Match{}, fmt.Errorf("Failed to decode match-season response: %s", err)
	}
	log.Printf("%+v\n", match)
	return match, nil
}
