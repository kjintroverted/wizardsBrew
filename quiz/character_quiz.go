package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kjintroverted/wizardsBrew/api/nodes"
	"github.com/kjintroverted/wizardsBrew/psql"
)

// BeginQuiz starts a quiz to help choose a
// race and class for character building
func BeginQuiz() {
	db, err := psql.NewPostgresConnection()
	if err != nil {
		fmt.Println("ERR", err)
	}
	defer db.Close()
	service := nodes.NewNodeService(nodes.NewNodeRepo(db))

	reader := bufio.NewReader(os.Stdin)

	// GET START NODE
	raceNode, err := service.FindByID(`16`)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	// START INTERACTION LOOP
	for {
		fmt.Println()
		if raceNode.Type == "terminus" {
			break
		}

		// PROMPT
		fmt.Println(raceNode.Value)
		for i, c := range raceNode.Paths {
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
		if err != nil || choice > len(raceNode.Paths) || choice < 1 {
			fmt.Printf("Please enter a valid option (1-%d)", len(raceNode.Paths))
			continue
		}

		// LOAD NEXT NODE
		raceNode, err = service.FindByID(raceNode.Paths[choice-1].Value)
	}

	// GET START NODE
	classNode, err := service.FindByID(`35`)
	if err != nil {
		fmt.Println("ERR", err)
		return
	}

	// START INTERACTION LOOP
	for {
		fmt.Println()
		if classNode.Type == "terminus" {
			fmt.Printf("Your result: %s\n", classNode.Value)
			break
		}

		// PROMPT
		fmt.Println(classNode.Value)
		for i, c := range classNode.Paths {
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
		if err != nil || choice > len(classNode.Paths) || choice < 1 {
			fmt.Printf("Please enter a valid option (1-%d)", len(classNode.Paths))
			continue
		}

		// LOAD NEXT NODE
		classNode, err = service.FindByID(classNode.Paths[choice-1].Value)
	}

	fmt.Printf("Your Result: %s %s\n", raceNode.Value, classNode.Value)
}
