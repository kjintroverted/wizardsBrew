package items

import (
	"fmt"
)

// Item describes gear, weapons, and armor
type Item struct {
	ID     string      `json:"id" db:"id"`
	Name   string      `json:"name" db:"name"`
	Type   string      `json:"type" db:"type"`
	Cost   float64     `json:"cost" db:"cost"`
	Weight float64     `json:"weight" db:"weight"`
	Attune string      `json:"attune" db:"attune"`
	Rarity string      `json:"rarity" db:"rarity"`
	Weapon weaponInfo  `json:"weapon" db:"weapon"`
	AC     int         `json:"ac" db:"ac"`
	Info   description `json:"info" db:"info"`
}

type weaponInfo struct {
	Category   string `json:"category" db:"category"`
	Damage     string `json:"damage" db:"damage"`
	DamageType string `json:"damageType" db:"damageType"`
}

func (w *weaponInfo) Scan(value interface{}) (err error) {
	str := string(value.([]byte))
	fmt.Println("WPN:", str)
	return
}

type description struct {
	Title string   `json:"title" db:"title"`
	Body  []string `json:"body" db:"body"`
}

func (description *description) Scan(value interface{}) (err error) {
	str := string(value.([]byte))
	fmt.Println("DSC:", str)
	return
}
