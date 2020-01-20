package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var dmgTypes map[string]string = map[string]string{
	"B": "bludgeoning",
	"N": "necrotic",
	"P": "piercing",
	"R": "radiant",
	"S": "slashing",
}

var itemTypes map[string]string = map[string]string{
	"AT":  "Artisan Tools",
	"A":   "Ammunition",
	"TAH": "Tack and Harness",
	"LA":  "Light Armor",
	"MA":  "Medium Armor",
	"HA":  "Heavy Armor",
	"P":   "Potion",
	"EXP": "Explosive",
	"RD":  "Rod",
	"R":   "Ranged Weapon",
	"AF":  "Futuristic",
	"S":   "Shield",
	"SHP": "Water Vehicle",
	"GS":  "Gaming Set",
	"G":   "Adventuring Gear",
	"MNT": "Mount",
	"VEH": "Land Vehicle",
	"T":   "Tool",
	"M":   "Melee Weapon",
	"SCF": "Spellcasting Focus",
	"WD":  "Wand",
	"SC":  "Scroll",
	"$":   "Treasure",
	"INS": "Instrument",
	"RG":  "Ring",
	"TG":  "Trade Good",
	"AIR": "Air Vehicle",
	"OTH": "Other",
}

// GenerateItemInserts will pull in a json file of
// items and generate SQL insert statements
func GenerateItemInserts() {
	fileData, _ := ioutil.ReadFile("data/json/ref.json")
	var ref map[string]interface{}
	if err := json.Unmarshal(fileData, &ref); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fileData, _ = ioutil.ReadFile("data/json/items.json")
	var items []map[string]interface{}
	if err := json.Unmarshal(fileData, &items); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	f, err := os.Create("data/sql/srd/items_gen.pgsql")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	x := 0
	// BEGIN ITEM LOOP
	for _, item := range items {
		// FILTER NON STANDARD ITEMS
		if item["source"].(string) != "PHB" && item["source"].(string) != "DMG" {
			continue
		}
		x++

		// GET ITEM TYPE
		itemType := "Wondrous Item"
		if item["type"] != nil {
			itemType = itemTypes[item["type"].(string)]
		}

		// GET ITEM COST
		cost := "null"
		if item["value"] != nil {
			cost = strconv.FormatFloat(item["value"].(float64)/100, 'f', -1, 64)
		}

		// GET ITEM WEIGHT
		weight := "null"
		if item["weight"] != nil {
			weight = strconv.FormatFloat(item["weight"].(float64), 'f', -1, 64)
		}

		// GET ITEM WEIGHT
		reqAttune := "false"
		if item["reqAttune"] != nil {
			var ok bool
			reqAttune, ok = item["reqAttune"].(string)
			if !ok {
				reqAttune = strconv.FormatBool(item["reqAttune"].(bool))
			}
		}

		// GET ITEM RARITY
		rarity := "Common"
		if item["rarity"] != nil {
			rarity = item["rarity"].(string)
			if rarity == "None" {
				rarity = "Common"
			}
		}

		// GET ITEM WEAPON INFO
		weapon := "null"
		if item["weaponCategory"] != nil {
			dmgType := "null"
			if item["dmgType"] != nil {
				dmgType = fmt.Sprintf("'%s'", dmgTypes[item["dmgType"].(string)])
			}
			dmg := "null"
			if item["dmg1"] != nil {
				dmg = fmt.Sprintf("'%s'", item["dmg1"])
			}
			weapon = fmt.Sprintf("row('%s', %s, %s)::weapon_info", item["weaponCategory"], dmg, dmgType)
		}

		// GET ITEM AC
		ac := "null"
		if item["ac"] != nil {
			ac = strconv.FormatFloat(item["ac"].(float64), 'f', -1, 64)
		}

		var info []section
		// LOAD PROPERTIES TO INFO
		if item["property"] != nil {
			for _, p := range item["property"].([]interface{}) {
				n, d, ok := getProp(p.(string), ref)
				if !ok {
					continue
				}
				if p.(string) == "V" {
					n += " (" + item["dmg2"].(string) + ")"
				}
				if p.(string) == "T" || p.(string) == "A" {
					n += " (" + item["range"].(string) + " ft)"
				}
				info = append(info, section{Title: n, Body: d})
			}
		}

		// LOAD ENTRIES INTO INFO
		if item["entries"] != nil {
			var paragraph []interface{}
			for _, e := range item["entries"].([]interface{}) {
				if str, ok := e.(string); ok {
					paragraph = append(paragraph, san(str))
				} else {
					entryMap := e.(map[string]interface{})
					if entryMap["type"].(string) == "entries" {
						info = append(info, section{Title: san(entryMap["name"].(string)), Body: sanAll(entryMap["entries"].([]interface{}))})
					}
					if entryMap["type"].(string) == "table" {
						title := createRow(entryMap["colLabels"].([]interface{}))
						var body []interface{}
						for _, row := range entryMap["rows"].([]interface{}) {
							rowStr := createRow(row.([]interface{}))
							body = append(body, san(rowStr))
						}
						info = append(info, section{Title: san(title), Body: body})
					}
				}
			}
			if len(paragraph) > 0 {
				info = append(info, section{Title: "", Body: paragraph})
			}
		}

		infoInsert := "null"
		if len(info) > 0 {
			infoInsert = "array["
			for _, d := range info {
				infoInsert += fmt.Sprintf("row('%s', array[", d.Title)
				for _, b := range d.Body {
					infoInsert += fmt.Sprintf("'%s',", b)
				}
				infoInsert = strings.Trim(infoInsert, ",") + "])::section,"
			}
			infoInsert = strings.Trim(infoInsert, ",") + "]"
		}

		statement := fmt.Sprintf("INSERT INTO items (name, type, cost, weight, attune, rarity, weapon, armor_class, info) VALUES ('%s', '%s', %s, %s, '%s', '%s', %s, %s, %s);\n",
			san(item["name"].(string)), san(itemType), cost, weight, san(reqAttune), rarity, weapon, ac, infoInsert)

		f.WriteString(statement)
	}

	if err := f.Sync(); err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(x, "Items")
}

func getProp(prop string, ref map[string]interface{}) (name string, desc []interface{}, ok bool) {
	props := ref["itemProperty"].([]interface{})
	for _, propRef := range props {
		refMap := propRef.(map[string]interface{})
		if refMap["entries"] == nil {
			continue
		}
		if prop == refMap["abbreviation"].(string) {
			for _, info := range refMap["entries"].([]interface{}) {
				infoMap := info.(map[string]interface{})
				return infoMap["name"].(string), sanAll(infoMap["entries"].([]interface{})), true
			}
		}
	}
	return "", nil, false
}
