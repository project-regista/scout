package main

import (
	"fmt"
	"log"

	scout "github.com/project-regista/scout/client"
	"github.com/project-regista/scout/competition"
)

func main() {

	// Get Scout client
	client, err := scout.New("./client/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Competition by ID
	comp, err := competition.GetCompetition(client, "43")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Competition by ID: %+v\n", comp)

	// List of Competitions
	comps, err := competition.GetCompetitions(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of Competitions: %+v\n", comps)

	// List Competitions w/ Seasons
	compsSeasons, err := competition.GetCompetitionsSeasons(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List of Competitions w/ Seasons: %+v\n", compsSeasons)
}
