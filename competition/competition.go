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

func requestCompetitions() Competition {

	configuration.LoadConfig()

	a := configuration.ApiAuth()
	a.GetAPI()

	req, err := http.NewRequest("GET", a.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	comp := Competition{}

	err = json.Unmarshal(content, &comp)
	if err != nil {
		log.Fatal(err)
	}

	return comp

}

func prepareStatements(comp Competition) []string {
	var statements []string

	for i := 0; i < len(comp.Data); i++ {

		compID := strconv.Itoa(comp.Data[i].ID)
		compName := comp.Data[i].Name
		countryID := strconv.Itoa(comp.Data[i].Country.ID)
		countryName := comp.Data[i].Country.Name

		str := "MERGE (comp:Competition{id:" + compID + "})\n" +
			"ON CREATE SET comp.name='" + compName + "'\n" +
			"ON MATCH SET comp.name='" + compName + "'\n" +
			"MERGE (country:Country{id:" + countryID + "})\n" +
			"ON CREATE SET country.name='" + countryName + "'\n" +
			"ON MATCH SET country.name='" + countryName + "'\n" +
			"MERGE (country)-[:ORGANISES]->(comp)"

		statements = append(statements, str)
	}

	return statements
}

func GetCompetitions() Competition {
	competitions := requestCompetitions()

	scout4neo.StoreData(prepareStatements(competitions))

	return competitions
}
