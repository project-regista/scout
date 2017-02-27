package competition

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"

	"github.com/project-regista/scout/configuration"
	"github.com/project-regista/scout/scout4neo"
)

// Competition defines a structure to hold data from competition in the socceramaAPI
type Competition struct {
	Data []struct {
		ID      int    `json:"id"`
		Cup     bool   `json:"cup"`
		Name    string `json:"name"`
		Active  bool   `json:"active"`
		Country struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			IsoCode string `json:"iso_code"`
			Flag    string `json:"flag"`
		} `json:"country"`
		Seasons struct {
			Data []struct {
				ID            int    `json:"id"`
				CompetitionID int    `json:"competition_id"`
				Name          string `json:"name"`
				Active        bool   `json:"active"`
			} `json:"data"`
		} `json:"seasons"`
	} `json:"data"`
}

// Uses socceramaAPI to get data for competition, including country data
func requestCompetitions() (Competition, error) {
	// The configurtion is loaded
	configuration.LoadConfig()
	a := configuration.APIAuth()
	a.GetAPI()

	// Makes a request to the API for competitons and country
	req, err := http.NewRequest("GET", "https://api.soccerama.pro/v1.2/competitions?api_token="+a.Key+"&include=country,currentSeason,seasons", nil)
	if err != nil {
		return Competition{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Competition{}, err
	}

	// Read all responses into content
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Competition{}, err
	}

	defer resp.Body.Close()

	// Instantiates a new struct for Competitions
	comp := Competition{}

	// Fill the competitons with data inside JSON
	err = json.Unmarshal(content, &comp)
	if err != nil {
		return Competition{}, err
	}

	return comp, nil

}

// Prepares the statements to be sent to scout4neo
func prepareStatements(comp Competition) []string {
	var statements []string

	for _, comp := range comp.Data {
		// Stores data from each competition
		compID := strconv.Itoa(comp.ID)
		compName := comp.Name
		countryID := strconv.Itoa(comp.Country.ID)
		countryName := comp.Country.Name

		var buffer bytes.Buffer

		for i, season := range comp.Seasons.Data {
			seasonID := strconv.Itoa(season.ID)
			seasonName := season.Name

			buffer.WriteString(
				"MERGE(season" + strconv.Itoa(i) + ":Season{id:" + seasonID + "})\n" +
					"ON CREATE SET season" + strconv.Itoa(i) + ".name='" + seasonName + "'\n" +
					"ON MATCH SET season" + strconv.Itoa(i) + ".name='" + seasonName + "'\n" +
					"MERGE (comp)-[:HAS_SEASON]->(season" + strconv.Itoa(i) + ")\n",
			)
		}

		// Construct a string that will store the competitons data in Neo4j
		str := "MERGE (comp:Competition{id:" + compID + "})\n" +
			"ON CREATE SET comp.name='" + compName + "'\n" +
			"ON MATCH SET comp.name='" + compName + "'\n" +
			"MERGE (country:Country{id:" + countryID + "})\n" +
			"ON CREATE SET country.name='" + countryName + "'\n" +
			"ON MATCH SET country.name='" + countryName + "'\n" +
			"MERGE (country)-[:ORGANISES]->(comp)\n" +
			buffer.String()

		// Array with all cypher statements
		statements = append(statements, str)
	}

	return statements
}

// GetCompetitions is the main function to access all methos in this module
func GetCompetitions() Competition {
	competitions, err := requestCompetitions()
	if err != nil {
		log.Fatal(err)
	}

	// prepareStatements(competitions)
	scout4neo.StoreData(prepareStatements(competitions))

	return competitions
}
