package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// GenerateBackgroundInserts pulls data from json
// and converts to sql inserts
func GenerateBackgroundInserts() {
	fileData, _ := ioutil.ReadFile("data/json/backgrounds.json")
	var backgrounds []map[string]interface{}
	if err := json.Unmarshal(fileData, &backgrounds); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	f, err := os.Create("data/sql/srd/backgrounds_gen.pgsql")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	x := 0
	// BEGIN ITEM LOOP
	for _, background := range backgrounds {
		if s, ok := background["source"]; ok && s != "PHB" {
			continue
		}
		x++

		// SKILLS
		var skills []interface{}
		if skillArr, ok := background["skillProficiencies"].([]interface{}); ok {
			for sk := range skillArr[0].(map[string]interface{}) {
				skills = append(skills, sk)
			}
		}

		// TOOLS
		var tools []interface{}
		if toolArr, ok := background["toolProficiencies"].([]interface{}); ok {
			for t, v := range toolArr[0].(map[string]interface{}) {
				if t == "choose" {
					from := v.(map[string]interface{})["from"].([]interface{})
					tools = append(tools, join(from, " or "))
					continue
				}
				tools = append(tools, t)
			}
		}

		statement := fmt.Sprintf("INSERT INTO backgrounds (name, pro_skill, pro_tool, language, equipment, special_opts, character_opts) VALUES ('%s',%s,%s,%s,%s,%s,%s);\n",
			background["name"], simpleStrArray(skills), simpleStrArray(tools), "null", "null", "null", "null")

		f.WriteString(statement)
	}

	if err := f.Sync(); err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(x, "Backgrounds")
}
