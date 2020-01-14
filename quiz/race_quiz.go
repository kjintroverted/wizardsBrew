package quiz

import (
	"fmt"

	"github.com/kjintroverted/wizardsBrew/db"
)

// BeginRaceQuiz starts a quiz to help choose a race
// for character building
func BeginRaceQuiz() {
	db, err := db.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	db.Close()
}
