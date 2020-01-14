package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func createRow(cells []interface{}) string {
	s := "|"
	for _, c := range cells {
		s += fmt.Sprintf(" %s |", c)
	}
	return s
}

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
	fileData, _ := ioutil.ReadFile("data/spells.json")
	var spells []map[string]interface{}
	if err := json.Unmarshal(fileData, &spells); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	f, err := os.Create("sql/spells.pgsql")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	x := 0
	// var collection = make(map[interface{}][]string)
	// BEGIN ITEM LOOP
	for _, spell := range spells {
		x++

		// GET TIME
		timeArr := spell["time"].([]interface{})
		timeObj := timeArr[0].(map[string]interface{})
		time := fmt.Sprintf("%s %s", strconv.FormatFloat(timeObj["number"].(float64), 'f', -1, 64), timeObj["unit"])
		condition := "null"
		if timeObj["condition"] != nil {
			condition = fmt.Sprintf("'%s'", san(timeObj["condition"].(string)))
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
				componentStr += fmt.Sprintf("row('material', '%s', null, null)::spell_comp", san(infoStr))
			} else {
				infoObj := info.(map[string]interface{})
				cost := "0"
				if infoObj["cost"] != nil {
					cost = strconv.FormatFloat(infoObj["cost"].(float64), 'f', -1, 64)
				}
				_, consume := infoObj["consume"]
				componentStr += fmt.Sprintf("row('material', '%s', %s, %s)::spell_comp", san(infoObj["text"].(string)), cost, strconv.FormatBool(consume))
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
		classObj := spell["classes"].(map[string]interface{})
		classArr := classObj["fromClassList"].([]interface{})
		classStr := "array["
		for _, c := range classArr {
			classInfo := c.(map[string]interface{})
			if classInfo["source"].(string) != "PHB" {
				continue
			}
			classStr += fmt.Sprintf("'%s',", classInfo["name"].(string))
		}
		classStr = strings.Trim(classStr, ",") + "]"

		// GET DESC
		entries := spell["entries"].([]interface{})
		if higher, ok := spell["entriesHigherLevel"]; ok {
			entries = append(entries, higher.([]interface{})...)
		}

		var description []section
		description = append(description, section{Title: ""})
		var body []interface{}
		for _, e := range entries {
			if p, ok := e.(string); ok {
				body = append(body, san(p))
			} else {
				info := e.(map[string]interface{})
				if info["type"].(string) == "table" {
					title := createRow(info["colLabels"].([]interface{}))
					var rows []interface{}
					for _, row := range info["rows"].([]interface{}) {
						rows = append(rows, createRow(row.([]interface{})))
					}
					description = append(description, section{Title: title, Body: body})
				} else if info["type"].(string) == "list" {
					description = append(description, section{Title: "choices", Body: sanAll(info["items"].([]interface{}))})
				} else {
					description = append(description, section{Title: san(info["name"].(string)), Body: sanAll(info["entries"].([]interface{}))})
				}
			}
		}
		description[0] = section{Title: "", Body: body}

		descStr := "array["
		for _, d := range description {
			descStr += fmt.Sprintf("row('%s', array[", d.Title)
			for _, b := range d.Body {
				descStr += fmt.Sprintf("'%s',", b)
			}
			descStr = strings.Trim(descStr, ",") + "])::section,"
		}
		descStr = strings.Trim(descStr, ",") + "]"

		statement := fmt.Sprintf("INSERT INTO spells (name, school, time, duration, comp, concentrate, range, level, class, description) VALUES ('%s', '%s', %s, '%s', %s, %s, '%s', %s, %s, %s);\n",
			san(spell["name"].(string)), schoolMap[spell["school"]], spellTime, duration, componentStr, strconv.FormatBool(concentrate), distanceStr, level, classStr, descStr)

		f.WriteString(statement)
	}

	if err := f.Sync(); err != nil {
		fmt.Println("ERROR:", err)
	}

	// for d, arr := range collection {
	// 	fmt.Println(d, len(arr), arr[0])
	// }

	fmt.Println(x, "Spells")
}
