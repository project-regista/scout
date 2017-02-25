package competition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/project-regista/scout/configuration"
	"github.com/spf13/viper"
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

type Auth struct {
	user     string
	password string
	host     string
	port     string
	apitkn   string
	URL      string
}

func apiAuth() *Auth {
	return &Auth{
		apitkn: viper.GetString("soccerama.apitkn"),
	}
}

func (a *Auth) getAPI() {
	a.URL = fmt.Sprintf("https://api.soccerama.pro/v1.2/competitions?api_token=%s&include=country",
		a.apitkn)
}

func dbAuth() *Auth {
	return &Auth{
		user:     viper.GetString("database.user"),
		password: viper.GetString("database.password"),
		host:     viper.GetString("database.host"),
		port:     viper.GetString("database.port"),
	}
}

func (a *Auth) getURL() {
	a.URL = fmt.Sprintf("bolt://%s:%s@%s:%s/db/data",
		a.user, a.password, a.host, a.port)
}

func requestCompetitions() Competition {

	configuration.LoadConfig()

	a := apiAuth()
	a.getAPI()

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

func storeData(statement []string) {

	configuration.LoadConfig()

	a := dbAuth()
	a.getURL()

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(a.URL)

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
