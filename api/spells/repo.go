package spells

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// SpellRepo defines the necessary actions to
// interact with Spell data
type SpellRepo interface {
	FindByID(id string) (*Spell, error)
}

type itemRepo struct {
	db *sql.DB
}

// NewSpellRepo returns a SpellRepo with a db connection
func NewSpellRepo(db *sql.DB) SpellRepo {
	return &itemRepo{
		db,
	}
}

func (r *itemRepo) FindByID(id string) (spell *Spell, err error) {
	sql := `SELECT * FROM spells WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanSpell(row)
}

func scanSpell(row psql.Scannable) (spell *Spell, err error) {
	spell = new(Spell)
	if err := row.Scan(
		&spell.ID,
		&spell.Name,
		&spell.School,
		&spell.Time,
		&spell.Duration,
		pq.Array(&spell.Components),
		&spell.Concentrate,
		&spell.Range,
		&spell.Level,
		pq.Array(&spell.Class),
		pq.Array(&spell.Description)); err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}
	return
}
