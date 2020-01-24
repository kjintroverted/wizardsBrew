package backgrounds

import (
	"github.com/kjintroverted/wizardsBrew/psql"
)

// Background defines background ideas for a character
type Background struct {
	ID            string             `json:"id" db:"id"`
	Name          string             `json:"name" db:"name"`
	ProSkill      []string           `json:"proSkill" db:"proSkill"`
	ProTool       []string           `json:"proTool" db:"proTool"`
	Language      []string           `json:"language" db:"language"`
	Equipment     []string           `json:"equipment" db:"equipment"`
	SpecialOpts   []string           `json:"specialOpts" db:"specialOpts"`
	CharacterOpts []psql.Description `json:"characterOpts" db:"characterOpts"`
}
