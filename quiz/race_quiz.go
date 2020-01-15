package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kjintroverted/wizardsBrew/nodes"
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
	service := nodes.NewNodeService(nodes.NewNodeRepo(db))

	// GET START NODE
	node, err := service.FindByID(`16`)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	// START INTERACTION LOOP
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		if node.Type == "terminus" {
			fmt.Printf("Your result: %s\n", node.Value)
			break
		}

		// PROMPT
		fmt.Println(node.Value)
		for i, c := range node.Paths {
			fmt.Printf("%d. %s\n", i+1, c.Prompt)
		}
		fmt.Print(">> ")

		// PARSE ANSWER
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdString = strings.Trim(strings.TrimSuffix(cmdString, "\n"), " ")
		choice, err := strconv.Atoi(cmdString)
		if err != nil || choice > len(node.Paths) || choice < 1 {
			fmt.Printf("Please enter a valid option (1-%d)", len(node.Paths))
			continue
		}

		// LOAD NEXT NODE
		node, err = service.FindByID(node.Paths[choice-1].Value)
	}
}
