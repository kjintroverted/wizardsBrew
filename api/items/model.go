package items

import (
	"strings"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// Item describes gear, weapons, and armor
type Item struct {
	ID     string             `json:"id" db:"id"`
	Name   string             `json:"name" db:"name"`
	Type   string             `json:"type" db:"type"`
	Cost   psql.NullFloat     `json:"cost" db:"cost"`
	Weight psql.NullFloat     `json:"weight" db:"weight"`
	Attune psql.NullString    `json:"attune" db:"attune"`
	Rarity psql.NullString    `json:"rarity" db:"rarity"`
	Weapon *weaponInfo        `json:"weapon,omitempty" db:"weapon"`
	AC     psql.NullInt       `json:"ac,omitempty" db:"ac"`
	Info   []psql.Description `json:"info,omitempty" db:"info"`
}

type weaponInfo struct {
	Category   string `json:"category,omitempty" db:"category"`
	Damage     string `json:"damage,omitempty" db:"damage"`
	DamageType string `json:"damageType,omitempty" db:"damageType"`
}

func (w *weaponInfo) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}
	str := string(value.([]byte))
	vals := strings.Split(strings.Trim(str, "()"), ",")
	w.Category = vals[0]
	w.Damage = vals[1]
	w.DamageType = vals[2]
	return
}
