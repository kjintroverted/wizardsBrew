package items

import (
	"github.com/kjintroverted/wizardsBrew/psql"
)

// Item describes gear, weapons, and armor
type Item struct {
	ID         string             `json:"id" db:"id"`
	Name       string             `json:"name" db:"name"`
	Type       string             `json:"type" db:"type"`
	Cost       *psql.NullFloat    `json:"cost" db:"cost"`
	Weight     *psql.NullFloat    `json:"weight" db:"weight"`
	Attune     *psql.NullString   `json:"attune" db:"attune"`
	Rarity     *psql.NullString   `json:"rarity" db:"rarity"`
	Weapon     *weaponInfo        `json:"weapon,omitempty" db:"weapon"`
	AC         *psql.NullInt      `json:"ac,omitempty" db:"ac"`
	Info       []psql.Description `json:"info,omitempty" db:"info"`
	IsHomebrew bool               `json:"isHomebrew" db:"homebrew"`
}

type weaponInfo struct {
	Category   *psql.NullString `json:"category,omitempty" db:"category"`
	Damage     *psql.NullString `json:"damage,omitempty" db:"damage"`
	DamageType *psql.NullString `json:"damageType,omitempty" db:"damageType"`
}
