package classes

import (
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"
	"github.com/lib/pq"
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
	Progress       table              `json:"progress" json:"progress"`
}

type table []column

func (t *table) Scan(value interface{}) (err error) {
	for _, c := range *t {
		if err := c.Scan(value); err != nil {
			return err
		}
		fmt.Println("COL:", c)
		// *t = append(*t, c)
	}
	return
}

type column []string

func (c *column) Scan(value interface{}) error {
	scanner := pq.Array(&c)
	return scanner.Scan(value)
}
