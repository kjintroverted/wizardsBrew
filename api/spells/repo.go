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
	List(opt map[string][]string) ([]Spell, error)
}

type spellRepo struct {
	db *sql.DB
}

// NewSpellRepo returns a SpellRepo with a db connection
func NewSpellRepo(db *sql.DB) SpellRepo {
	return &spellRepo{
		db,
	}
}

func (r *spellRepo) FindByID(id string) (spell *Spell, err error) {
	sql := `SELECT * FROM spells WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanSpell(row)
}

func (r *spellRepo) List(opt map[string][]string) (spells []Spell, err error) {
	sql := `SELECT * FROM spells`
	i := 0
	for o, arr := range opt {
		clause := "where"
		if i > 0 {
			clause = "and"
		}
		switch o {
		case "class":
			sql += fmt.Sprintf(" %s '%s' ILIKE any (class)", clause, arr[0])
		case "level":
			sql += fmt.Sprintf(" %s level = %s", clause, arr[0])
		case "school":
			sql += fmt.Sprintf(" %s school ILIKE '%s'", clause, arr[0])
		}
		i++
	}
	rows, err := r.db.Query(sql + ";")
	if err != nil {
		return nil, fmt.Errorf("ERROR running query: `%s` (%s)", sql, err)
	}
	for rows.Next() {
		if spell, err := scanSpell(rows); err == nil {
			spells = append(spells, *spell)
		} else {
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
	}
	return
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
