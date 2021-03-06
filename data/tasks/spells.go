package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var schoolMap map[interface{}]string = map[interface{}]string{
	"T": "transmutation",
	"N": "necromancy",
	"I": "illusion",
	"V": "evocation",
	"E": "enchantment",
	"D": "divination",
	"C": "conjuration",
	"A": "abjuration",
}

// GenerateSpellInserts pulls data from json
// and converts to sql inserts
func GenerateSpellInserts() {
	f, err := os.Create("data/sql/srd/spells_gen.pgsql")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	dir := "data/json/spells"
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
		var spellData map[string]interface{}
		if err := json.Unmarshal(fileData, &spellData); err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		spellArr := spellData["spell"].([]interface{})
		for _, s := range spellArr {
			if statement, err := genSpellSQLString(s.(map[string]interface{})); err == nil {
				x++
				f.WriteString(stripFilters(statement))
				if err := f.Sync(); err != nil {
					fmt.Println("ERROR:", err)
				}
			}
		}
	}

	if err := f.Sync(); err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(x, "Spells")
}

func genSpellSQLString(spell map[string]interface{}) (string, error) {
	if spell["source"] != "PHB" && spell["source"] != "SCAG" && spell["source"] != "GGR" && spell["source"] != "XGE" {
		return "", fmt.Errorf("Invalid source")
	}

	// GET TIME
	timeArr := spell["time"].([]interface{})
	timeObj := timeArr[0].(map[string]interface{})
	time := fmt.Sprintf("%s %s", strconv.FormatFloat(timeObj["number"].(float64), 'f', -1, 64), timeObj["unit"])
	condition := "null"
	if timeObj["condition"] != nil {
		condition = fmt.Sprintf("'%s'", escape(timeObj["condition"].(string)))
	}
	spellTime := fmt.Sprintf("row('%s', %s)::spell_time", time, condition)

	//GET DURATION
	durArr := spell["duration"].([]interface{})
	durObject := durArr[0].(map[string]interface{})
	duration := durObject["type"].(string)
	if duration == "timed" {
		timeInfo := durObject["duration"].(map[string]interface{})
		duration = fmt.Sprintf("%s %s", strconv.FormatFloat(timeInfo["amount"].(float64), 'f', -1, 64), timeInfo["type"])
	}

	// GET CONCENTRATION
	_, concentrate := durObject["concentration"]

	// GET COMPONENTS
	comp := spell["components"].(map[string]interface{})
	componentStr := "array["
	if _, ok := comp["v"]; ok {
		componentStr += "row('verbal', 'A verbal component is a spoken incantation. To provide a verbal component, you must be able to speak in a strong voice.', null, null)::spell_comp,"
	}
	if _, ok := comp["s"]; ok {
		componentStr += "row('somatic', 'A somatic component is a measured and precise movement of the hand. You must have at least one hand free to provide a somatic component.', null, null)::spell_comp,"
	}
	if info, ok := comp["m"]; ok {
		if infoStr, ok := info.(string); ok {
			componentStr += fmt.Sprintf("row('material', '%s', null, null)::spell_comp", escape(infoStr))
		} else {
			infoObj := info.(map[string]interface{})
			cost := "0"
			if infoObj["cost"] != nil {
				cost = strconv.FormatFloat(infoObj["cost"].(float64), 'f', -1, 64)
			}
			_, consume := infoObj["consume"]
			componentStr += fmt.Sprintf("row('material', '%s', %s, %s)::spell_comp", escape(infoObj["text"].(string)), cost, strconv.FormatBool(consume))
		}
	}
	componentStr = strings.Trim(componentStr, ",")
	componentStr += "]"

	// GET RANGE
	rangeObj := spell["range"].(map[string]interface{})
	distanceStr := ""
	if distance, ok := rangeObj["distance"].(map[string]interface{}); ok {
		disType := distance["type"].(string)
		if disType == "miles" || disType == "feet" {
			distanceStr = strconv.FormatFloat(distance["amount"].(float64), 'f', -1, 64) + " "
		}
		distanceStr += disType
	} else {
		distanceStr = "special"
	}

	// GET LEVEL
	level := strconv.FormatFloat(spell["level"].(float64), 'f', -1, 64)

	// GET CLASSES
	classStr := "null"
	if classObj, ok := spell["classes"].(map[string]interface{}); ok {
		classArr := classObj["fromClassList"].([]interface{})
		classStr = "array["
		for _, c := range classArr {
			classInfo := c.(map[string]interface{})
			if classInfo["source"] != "PHB" && classInfo["source"] != "ERLW" {
				continue
			}
			classStr += fmt.Sprintf("'%s',", classInfo["name"].(string))
		}
		classStr = strings.Trim(classStr, ",") + "]"
	}

	// GET DESC
	entries := spell["entries"].([]interface{})
	if higher, ok := spell["entriesHigherLevel"]; ok {
		entries = append(entries, higher.([]interface{})...)
	}

	statement := fmt.Sprintf("INSERT INTO spells (name, school, time, duration, comp, concentrate, range, level, class, description) VALUES ('%s', '%s', %s, '%s', %s, %s, '%s', %s, %s, %s);\n",
		escape(spell["name"].(string)), schoolMap[spell["school"]], spellTime, duration, componentStr, strconv.FormatBool(concentrate), distanceStr, level, classStr, parseEntries(entries))

	return statement, nil
}
