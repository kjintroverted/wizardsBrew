package classes

import (
	"strings"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// Class defines a character class
type Class struct {
	ID             string             `json:"id" json:"id"`
	Name           string             `json:"name" json:"name"`
	HitDice        string             `json:"hitDice" json:"hitDice"`
	ProArmor       []string           `json:"proArmor" json:"proArmor"`
	ProWeapon      []string           `json:"proWeapon" json:"proWeapon"`
	ProTool        psql.NullString    `json:"proTool" json:"proTool"`
	ProSave        []string           `json:"proSave" json:"proSave"`
	Skills         string             `json:"skills" json:"skills"`
	StartEquipment []string           `json:"startEquip" json:"startEquip"`
	Description    []psql.Description `json:"description" json:"description"`
	ProgressString []string           `json:"-" json:"progress"`
	Progress       [][]string         `json:"progress"`
}

func (c *Class) expandTable() {
	for _, s := range c.ProgressString {
		c.Progress = append(c.Progress, strings.Split(s, "|"))
	}
}
