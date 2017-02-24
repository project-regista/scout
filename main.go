package main

import (
	"fmt"

	"github.com/project-regista/scout/competition"
)

func main() {
	competitions := competition.GetCompetitions()
	fmt.Println(competitions.Data)
}
