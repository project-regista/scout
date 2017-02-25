package competition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
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
	req, err := http.NewRequest("GET", "https://api.soccerama.pro/v1.2/competitions?api_token=IjtTiudWktGxEeAIXpTtJPSkG4RWaAVUdPZwSCeWU0aAmXtcN8n5ya7IUho7&include=country", nil)
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

func storeData(statement []string) {
	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo("bolt://neo4j:admin@localhost:7687")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for i := 0; i < len(statement); i++ {

		stmt, err := conn.PrepareNeo(statement[i])
		if err != nil {
			log.Fatal(err)
		}

		result, err := stmt.ExecNeo(map[string]interface{}{})

		if err != nil {
			log.Fatal(err)
		}
		numResult, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("CREATED ROWS: %d\n", numResult)

		stmt.Close()

	}
}

func GetCompetitions() Competition {
	competitions := requestCompetitions()

	storeData(prepareStatements(competitions))

	return competitions
}
