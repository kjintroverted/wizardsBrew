package feats

import (
	"github.com/kjintroverted/wizardsBrew/api/races"
	"github.com/kjintroverted/wizardsBrew/psql"
)

// Feat
type Feat struct {
	ID          string             `json:"id" db:"id"`
	Name        string             `json:"name" db:"name"`
	Ability     []races.AbilityMod `json:"ability" db:"ability"`
	Description []psql.Description `json:"description" db:"description"`
	Class       psql.NullString    `json:"class" db:"class"`
	Subclass    psql.NullString    `json:"subclass" db:"subclass"`
	Background  psql.NullString    `json:"background" db:"background_req"`
	Level       psql.NullInt       `json:"level" db:"level"`
	Prereq      psql.NullString    `json:"prereq" db:"prereq"`
}
