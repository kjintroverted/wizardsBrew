package races

import (
	"strconv"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// Race defines a DnD character race
type Race struct {
	ID       string             `json:"id" db:"id"`
	Name     string             `json:"name" db:"name"`
	Ability  []AbilityMod       `json:"ability" db:"ability"`
	Size     string             `json:"size" db:"size"`
	Speed    int                `json:"speed" db:"speed"`
	Age      string             `json:"age" db:"age"`
	Align    string             `json:"align" db:"align"`
	SizeDesc string             `json:"sizeDesc" db:"sizeDesc"`
	Traits   []psql.Description `json:"description" db:"description"`
}

// AbilityMod contains info for updating stats
type AbilityMod struct {
	Name string `json:"name" db:"name"`
	Mod  int    `json:"mod" db:"mod"`
}

// Scan is used to scan a record from the DB into a struct
func (a *AbilityMod) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}
	str := string(value.([]byte))
	vals := psql.ParseRow(str)

	a.Name = vals[0]
	a.Mod, _ = strconv.Atoi(vals[1])
	return
}
