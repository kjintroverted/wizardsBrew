package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

		// LANGUAGES
		var lang []interface{}
		if langArr, ok := background["languageProficiencies"].([]interface{}); ok {
			for _, vals := range langArr {
				for t, v := range vals.(map[string]interface{}) {
					if t == "choose" {
						from := v.(map[string]interface{})["from"].([]interface{})
						lang = append(lang, join(from, " or "))
						continue
					}
					if t == "anyStandard" {
						lang = append(lang, fmt.Sprintf("Choose %d", int(v.(float64))))
						continue
					}
					lang = append(lang, t)
				}
			}
		}

		var equip []string
		var special []interface{}
		if entries, ok := background["entries"].([]interface{}); ok {
			for _, e := range entries {
				entry := e.(map[string]interface{})
				// EQUIPMENT
				switch {
				case entry["type"] == "list":
					items := entry["items"].([]interface{})
					for _, i := range items {
						item := i.(map[string]interface{})
						if item["name"] == "Equipment" {
							equip = strings.Split(item["entry"].(string), ", ")
						}
					}
				case entry["name"] == "Favorite Schemes" || entry["name"] == "Specialty":
					items := entry["entries"].([]interface{})
					for _, i := range items {
						if item, ok := i.(map[string]interface{}); ok {
							if item["type"] == "table" {
								rows := item["rows"].([]interface{})
								for _, r := range rows {
									special = append(special, r.([]interface{})[1])
								}
							}
						}
					}
				}
			}
		}

		statement := fmt.Sprintf("INSERT INTO backgrounds (name, pro_skill, pro_tool, language, equipment, special_opts) VALUES ('%s',%s,%s,%s,%s,%s);\n",
			background["name"], simpleStrArray(skills), simpleStrArray(tools), simpleStrArray(lang), simplerStrArray(equip), simpleStrArray(special))

		f.WriteString(statement)
	}

	if err := f.Sync(); err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(x, "Backgrounds")
}
