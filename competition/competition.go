package competition

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func GetCompetitions() Competition {
	return requestCompetitions()
}
