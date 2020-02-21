package characters

import (
	"strconv"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// PC describes the data for a Playable Character
type PC struct {
	ID           string         `json:"id" db:"id"`
	Name         string         `json:"name" db:"name"`
	Owner        string         `json:"owner" db:"owner"`
	AuthUsers    []string       `json:"authUsers" db:"auth_users"`
	ReadUsers    []string       `json:"readUsers" db:"read_users"`
	RaceID       int            `json:"raceID" db:"race_id"`
	ClassID      int            `json:"classID" db:"class_id"`
	Subclass     string         `json:"subclass" db:"subclass"`
	BackgroundID int            `json:"backgroundID" db:"backgroundID"`
	Stats        stats          `json:"stats" db:"stats"`
	XP           int            `json:"xp" db:"xp"`
	HP           int            `json:"hp" db:"hp"`
	MaxHP        int            `json:"maxHP" db:"maxHP"`
	Init         int            `json:"initiative" db:"initiative"`
	ProSkills    []proSkill     `json:"proSkills" db:"proSkills"`
	ProTools     []string       `json:"proTools" db:"proTools"`
	ProWeapons   []string       `json:"proWeapons" db:"proWeapons"`
	Lang         []string       `json:"languages" db:"languages"`
	EquipmentIDs []psql.NullInt `json:"equipmentIDs" db:"equipment_ids"`
	WeaponIDs    []psql.NullInt `json:"weaponIDs" db:"weapon_ids"`
	InventoryIDs []psql.NullInt `json:"inventoryIDs" db:"inventory_ids"`
	Gold         float64        `json:"gold" db:"gold"`
	SpellIDs     []psql.NullInt `json:"spellIDs" db:"spell_ids"`
	SpecFeatIDs  []psql.NullInt `json:"specFeatIDs" db:"specFeat_ids"`
}

// func (pc PC) String() string {
// 	if b, err := json.Marshal(pc); err == nil {
// 		return string(b)
// 	}
// 	return "ERROR printing character..."
// }

type stats struct {
	STR int `json:"str" db:"str"`
	DEX int `json:"dex" db:"dex"`
	CON int `json:"con" db:"con"`
	INT int `json:"int" db:"int"`
	WIS int `json:"wis" db:"wis"`
	CHA int `json:"cha" db:"cha"`
}

// Scan is used to scan a record from the DB into a struct
func (s *stats) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}
	str := string(value.([]byte))
	vals := psql.ParseRow(str)

	s.STR, _ = strconv.Atoi(vals[0])
	s.DEX, _ = strconv.Atoi(vals[1])
	s.CON, _ = strconv.Atoi(vals[2])
	s.INT, _ = strconv.Atoi(vals[3])
	s.WIS, _ = strconv.Atoi(vals[4])
	s.CHA, _ = strconv.Atoi(vals[5])

	return
}

type proSkill struct {
	Name string `json:"name" db:"name"`
	Mult int    `json:"multiplier" db:"multiplier"`
}

// Scan is used to scan a record from the DB into a struct
func (s *proSkill) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}
	str := string(value.([]byte))
	vals := psql.ParseRow(str)

	s.Name = vals[0]
	s.Mult, _ = strconv.Atoi(vals[1])

	return
}
