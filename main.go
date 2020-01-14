package main

import (
	"fmt"
	"os"

	"github.com/kjintroverted/wizardsBrew/quiz"

	"github.com/kjintroverted/wizardsBrew/tasks"
)

func main() {
	var command string
	if command = os.Args[1]; command == "" {
		fmt.Println("Please enter a command.")
	}

	switch command {
	case "sql":
		tasks.GenerateItemInserts()
		tasks.GenerateSpellInserts()
	case "race":
		quiz.BeginRaceQuiz()
	}
}
