package main

import (
	"fmt"
	"log"

	scout "github.com/project-regista/scout/client"
	"github.com/project-regista/scout/request"
)

func main() {

	// Get Scout client
	client, err := scout.New("./client/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Competition by ID
	competition, err := request.GetCompetition(client, "43")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Competition by ID: %+v\n", competition)

	// List of Competitions
	competitions, err := request.GetCompetitions(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of Competitions: %+v\n", competitions)

	// List Competitions w/ Seasons
	competitionsSeasons, err := request.GetCompetitionsSeasons(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of Competitions w/ Seasons: %+v\n", competitionsSeasons)

	country, err := request.GetCountry(client, "13")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Country by ID: %+v\n", country)

	countries, err := request.GetCountries(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of Countries: %+v\n", countries)

	countriesCompetitions, err := request.GetCountriesCompetitions(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of Countries w/ Competitions: %+v\n", countriesCompetitions)
}
