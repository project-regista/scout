package scout4neo

import (
	"fmt"
	"log"

	"github.com/project-regista/scout/configuration"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func StoreData(statement []string) {

	configuration.LoadConfig()

	a := configuration.DbAuth()
	a.GetDB()

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
