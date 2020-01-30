package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// GenerateFeatInserts will pull in all json class files
// and generate SQL insert statements
func GenerateFeatInserts() {
	f, err := os.Create("data/sql/srd/feat_gen.pgsql")
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
		for _, c := range classArr {
			classInfo := c.(map[string]interface{})
			if classInfo["source"] != "PHB" && classInfo["source"] != "ERLW" {
				continue
			}
			if statements, err := genFeatSQLString(classInfo["classFeatures"].([]interface{}), classInfo["name"]); err == nil {
				for _, s := range statements {
					x++
					f.WriteString(s.(string))
				}
				if err := f.Sync(); err != nil {
					fmt.Println("ERROR:", err)
				}
			}

			if subArr, ok := classInfo["subclasses"].([]interface{}); ok {
				for _, sub := range subArr {
					subClass := sub.(map[string]interface{})
					if _, ok := subClass["source"]; ok && subClass["source"] != "PHB" && subClass["source"] != "ERLW" {
						continue
					}
					if statements, err := genSubClassFeatSQLString(subClass["subclassFeatures"].([]interface{}), classInfo["name"], subClass["name"]); err == nil {
						for _, s := range statements {
							x++
							f.WriteString(s.(string))
						}
						if err := f.Sync(); err != nil {
							fmt.Println("ERROR:", err)
						}
					}
				}
			}
		}

	}

	fileData, _ := ioutil.ReadFile("data/json/backgrounds.json")
	var backgrounds []map[string]interface{}
	if err := json.Unmarshal(fileData, &backgrounds); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for _, background := range backgrounds {
		if s, ok := background["source"].(string); ok && (strings.Contains(s, "AL") || strings.Contains(s, "PS") || strings.Contains(s, "BG")) {
			continue
		}
		if _, ok := background["_copy"]; ok {
			continue
		}
		// BACKGROUND FEATS
		var statements []string
		if entries, ok := background["entries"].([]interface{}); ok {
			for _, e := range entries {
				entry := e.(map[string]interface{})
				if v, ok := entry["name"]; ok {
					if str, ok := v.(string); ok && strings.Contains(str, "Feature") {
						statement := fmt.Sprintf("INSERT into feats (name, description, background) VALUES ('%s', %s, '%s');\n",
							escape(string(str[9:])), parseEntries(entry["entries"].([]interface{})), escape(background["name"].(string)))
						statements = append(statements, statement)
					}
				}
			}
		}
		for _, s := range statements {
			x++
			f.WriteString(s)
		}
		if err := f.Sync(); err != nil {
			fmt.Println("ERROR:", err)
		}
	}

	fmt.Println(x, "Feats")
}

func genFeatSQLString(feats []interface{}, class interface{}) (statements []interface{}, err error) {
	for _, lvl := range feats {
		lvlArr := lvl.([]interface{})
		for _, d := range lvlArr {
			data := d.(map[string]interface{})
			if s, ok := data["source"]; ok && s != "PHB" {
				continue
			}

			// ENTRIES
			entries := data["entries"].([]interface{})
			desc := parseEntries(entries)

			// LEVEL
			level := 1
			if x, err := findLevel(desc); err == nil {
				level = x
			}

			statements = append(statements,
				fmt.Sprintf("INSERT into feats (name, description, class, level) VALUES ('%s', %s, '%s', %d);\n",
					escape(data["name"].(string)), desc, class, level))
		}
	}
	return
}

func genSubClassFeatSQLString(feats []interface{}, class, subclass interface{}) (statements []interface{}, err error) {
	for _, lvl := range feats {
		lvlArr := lvl.([]interface{})
		lvlEntries := lvlArr[0].(map[string]interface{})["entries"].([]interface{})
		level := 1
		for _, d := range lvlEntries {
			data, ok := d.(map[string]interface{})
			if !ok {
				continue
			}
			if s, ok := data["source"]; ok && s != "PHB" {
				continue
			}
			if _, ok := data["entries"]; !ok {
				continue
			}

			// ENTRIES
			entries := data["entries"].([]interface{})
			desc := parseEntries(entries)

			// SUBCLASS
			sub := "null"
			if subclass != nil {
				sub = fmt.Sprintf("'%s'", subclass)
			}

			// LEVEL
			if x, err := findLevel(desc); err == nil {
				level = x
			}

			statements = append(statements,
				fmt.Sprintf("INSERT into feats (name, description, class, subclass, level) VALUES ('%s', %s, '%s', %s, %d);\n",
					escape(data["name"].(string)), desc, class, sub, level))
		}
	}
	return
}

func findLevel(desc string) (lvl int, err error) {
	if i := strings.Index(desc, "level"); i != -1 {
		lvlStr := string(desc[i-5 : i-3])
		x, err := strconv.Atoi(strings.Trim(lvlStr, " "))
		if err != nil {
			return -1, err
		}
		return x, nil
	}
	return -1, fmt.Errorf("level not found in description entries")
}
