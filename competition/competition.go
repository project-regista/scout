package competition

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"

	"github.com/project-regista/scout/configuration"
	"github.com/project-regista/scout/scout4neo"
)

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
	} `json:"data"`
}

// Uses socceramaAPI to get data for competition, including country data
func requestCompetitions() (Competition, error) {
	// The configurtion is loaded
	configuration.LoadConfig()
	a := configuration.APIAuth()
	a.GetAPI()

	// Makes a request to the API for competitons and country
	req, err := http.NewRequest("GET", "https://api.soccerama.pro/v1.2/competitions?api_token="+a.Key+"&include=country", nil)
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

	for i := 0; i < len(comp.Data); i++ {

		// Stores data from each competition
		compID := strconv.Itoa(comp.Data[i].ID)
		compName := comp.Data[i].Name
		countryID := strconv.Itoa(comp.Data[i].Country.ID)
		countryName := comp.Data[i].Country.Name

		// Construct a string that will store the competitons data in Neo4j
		str := "MERGE (comp:Competition{id:" + compID + "})\n" +
			"ON CREATE SET comp.name='" + compName + "'\n" +
			"ON MATCH SET comp.name='" + compName + "'\n" +
			"MERGE (country:Country{id:" + countryID + "})\n" +
			"ON CREATE SET country.name='" + countryName + "'\n" +
			"ON MATCH SET country.name='" + countryName + "'\n" +
			"MERGE (country)-[:ORGANISES]->(comp)"

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
	scout4neo.StoreData(prepareStatements(competitions))

	return competitions
}
