package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
		classInfo := classArr[0].(map[string]interface{})
		if classInfo["source"] != "PHB" {
			continue
		}
		if statements, err := genFeatSQLString(classInfo["classFeatures"].([]interface{}), classInfo["name"], nil); err == nil {
			for _, s := range statements {
				x++
				f.WriteString(s.(string))
			}
			if err := f.Sync(); err != nil {
				fmt.Println("ERROR:", err)
			}
		}
	}

	fmt.Println(x, "Feats")
}

func genFeatSQLString(feats []interface{}, class, subclass interface{}) (statements []interface{}, err error) {
	for _, lvl := range feats {
		lvlArr := lvl.([]interface{})
		for _, d := range lvlArr {
			data := d.(map[string]interface{})
			if s, ok := data["source"]; ok && s != "PHB" {
				continue
			}

			statements = append(statements,
				fmt.Sprintf("INSERT into feats (name, ability, description, class, subclass, background, level, prereq) VALUES (%s, %s, %s, %s, %s, %s, %s, %s);\n",
					"", "", "", "", "", "", "", ""))
		}
	}

	return
}
