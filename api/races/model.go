package races

import (
	"strconv"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// Race defines a DnD character race
type Race struct {
	ID       string             `json:"id" db:"id"`
	Name     string             `json:"name" db:"name"`
	Ability  []abilityMod       `json:"ability" db:"ability"`
	Size     string             `json:"size" db:"size"`
	Speed    int                `json:"speed" db:"speed"`
	Age      string             `json:"age" db:"age"`
	Align    string             `json:"align" db:"align"`
	SizeDesc string             `json:"sizeDesc" db:"sizeDesc"`
	Traits   []psql.Description `json:"description" db:"description"`
}

type abilityMod struct {
	Name string `json:"name" db:"name"`
	Mod  int    `json:"mod" db:"mod"`
}

func (a *abilityMod) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}
	str := string(value.([]byte))
	vals := psql.ParseRow(str)

	a.Name = vals[0]
	a.Mod, _ = strconv.Atoi(vals[1])
	return
}
