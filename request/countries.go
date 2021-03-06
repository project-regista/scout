package request

import (
	"encoding/json"
	"fmt"
	"log"

	scout "github.com/project-regista/scout/client"
)

// GetCountry get information about a single Country
func GetCountry(client scout.Client, id string) (Country, error) {

	requestURL := fmt.Sprintf("https://%s/%s/countries/%s",
		client.APIHost, client.APIVersion, id)

	params := map[string]string{"api_token": client.APIToken}

	body, err := Get(requestURL, params)
	if err != nil {
		return Country{}, fmt.Errorf("Failed to make country request: %s", err)
	}

	defer body.Close()
	dec := json.NewDecoder(body)

	var country Country

	if err := dec.Decode(&country); err != nil {
		return Country{}, fmt.Errorf("Failed to decode country response: %s", err)
	}
	log.Printf("%+v\n", country)
	return country, nil
}

// GetCountries get all the countries
func GetCountries(client scout.Client) (Countries, error) {

	requestURL := fmt.Sprintf("https://%s/%s/countries",
		client.APIHost, client.APIVersion)

	params := map[string]string{"api_token": client.APIToken}

	body, err := Get(requestURL, params)
	if err != nil {
		return Countries{}, fmt.Errorf("Failed to make countries request: %s", err)
	}

	// Decode the HTTP response into a Competition struct
	defer body.Close()
	dec := json.NewDecoder(body)

	var countries Countries

	if err := dec.Decode(&countries); err != nil {
		return Countries{}, fmt.Errorf("Failed to decode countries response: %s", err)
	}
	log.Printf("%+v\n", countries)
	return countries, nil
}
