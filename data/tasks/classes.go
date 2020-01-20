package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GenerateClassInserts will pull in all json class files
// and generate SQL insert statements
func GenerateClassInserts() {
	f, err := os.Create("data/sql/srd/class_gen.pgsql")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	dir := "data/json/class"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	x := 0
	// BEGIN ITEM LOOP
	for _, fileInfo := range files {
		fileName := fmt.Sprintf("%s/%s", dir, fileInfo.Name())
		fileData, _ := ioutil.ReadFile(fileName)
		var classData map[string]interface{}
		if err := json.Unmarshal(fileData, &classData); err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		classArr := classData["class"].([]interface{})
		if statement, err := genSQLString(classArr[0].(map[string]interface{})); err == nil {
			x++
			f.WriteString(statement)
			if err := f.Sync(); err != nil {
				fmt.Println("ERROR:", err)
			}
		}
	}

	fmt.Println(x, "Classes")
}

func genSQLString(data map[string]interface{}) (statement string, err error) {
	if data["source"] != "PHB" {
		return "", fmt.Errorf("Not standard class: %s", data["name"])
	}

	// HIT DICE
	hd := data["hd"].(map[string]interface{})
	hitDice := fmt.Sprintf("%dd%d", int(hd["number"].(float64)), int(hd["faces"].(float64)))

	// PROFICIENCIES
	pro := data["startingProficiencies"].(map[string]interface{})
	// 		ARMOR
	var armor []interface{}
	if v, ok := pro["armor"]; ok {
		armor = v.([]interface{})
	}
	// 		WEAPONS
	var weapons []interface{}
	if v, ok := pro["weapons"]; ok {
		weapons = v.([]interface{})
	}
	// 		TOOLS
	tools := "null"
	if v, ok := pro["tools"]; ok {
		tools = fmt.Sprintf("'%s'", san(v.([]interface{})[0].(string)))
	}
	// 		SKILLS
	skills := "null"
	if v, ok := pro["skills"]; ok {
		skillData := v.([]interface{})[0].(map[string]interface{})
		choose := skillData["choose"].(map[string]interface{})
		skills = fmt.Sprintf("'Choose %d: %s'", int(choose["count"].(float64)), join(choose["from"].([]interface{}), ", "))
	}

	// SAVING THROWS
	stats := data["proficiency"].([]interface{})
	for i, s := range stats {
		stats[i] = strings.ToUpper(s.(string))
	}

	// STARTING EQUIP
	var equip []interface{}
	if v, ok := data["startingEquipment"]; ok {
		equipInfo := v.(map[string]interface{})
		equip = equipInfo["default"].([]interface{})
	}

	// DESC
	fluff := data["fluff"].([]interface{})
	entries := fluff[0].(map[string]interface{})["entries"].([]interface{})
	var descArr []section
	var p []interface{}
	for _, v := range entries {
		if s, ok := v.(string); ok {
			p = append(p, s)
			continue
		}
		valMap := v.(map[string]interface{})
		descArr = append(descArr, section{Title: valMap["name"].(string), Body: valMap["entries"].([]interface{})})
	}

	var descRows []interface{}
	descRows = append(descRows, rowString("section", "''", simpleStrArray(p)))
	for _, sec := range descArr {
		descRows = append(descRows, rowString("section", fmt.Sprintf("'%s'", san(sec.Title)), simpleStrArray(sec.Body)))
	}

	statement = fmt.Sprintf("INSERT into classes (name, hit_dice, pro_armor, pro_weapon, pro_tool, pro_save, skills, init_equip, description, progress) VALUES ('%s', '%s', %s, %s, %s, %s, %s, %s, %s, %s);\n",
		data["name"], hitDice, simpleStrArray(armor), simpleStrArray(weapons), tools, simpleStrArray(stats), skills, simpleStrArray(equip), simpleArray(descRows), "")
	return
}
