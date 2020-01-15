package quiz

import (
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// BeginRaceQuiz starts a quiz to help choose a race
// for character building
func BeginRaceQuiz() {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	repo := psql.NewNodeRepo(db)

	// GET START NODE
	node, err := repo.FindByID(16)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}
	fmt.Println(node.Value)
}
