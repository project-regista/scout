package competition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

func storeCompetitions(comp Competition) {
	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo("bolt://neo4j:admin@localhost:7687")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for i := 0; i < len(comp.Data); i++ {

		stmt, err := conn.PrepareNeo(
			"MERGE (comp:Competition{id:{comp_id}})" +
				"ON CREATE SET comp.name='" + comp.Data[i].Name + "'" +
				"ON MATCH SET comp.name='" + comp.Data[i].Name + "'" +
				"MERGE (country:Country{id:{country_id}})" +
				"ON CREATE SET country.name='" + comp.Data[i].Country.Name + "'" +
				"ON MATCH SET country.name='" + comp.Data[i].Country.Name + "'" +
				"MERGE (country)-[:ORGANISES]->(comp)")
		if err != nil {
			panic(err)
		}

		result, err := stmt.ExecNeo(map[string]interface{}{
			"comp_id": comp.Data[i].ID,
			// "comp_name":    comp.Data[i].Name,
			"country_id": comp.Data[i].Country.ID,
			// "country_name": comp.Data[i].Country.Name
		})

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

	storeCompetitions(competitions)

	return competitions
}
