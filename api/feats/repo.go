package feats

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"
	"github.com/lib/pq"
)

// FeatRepo defines the necessary actions to
// interact with Feat data
type FeatRepo interface {
	FindByID(id string) (*Feat, error)
	List(opt map[string][]string) ([]Feat, error)
}

type featRepo struct {
	db *sql.DB
}

// NewFeatRepo returns a FeatRepo with a db connection
func NewFeatRepo(db *sql.DB) FeatRepo {
	return &featRepo{
		db,
	}
}

func (r *featRepo) FindByID(id string) (feat *Feat, err error) {
	sql := `SELECT * FROM feats WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanFeat(row)
}

func (r *featRepo) List(opt map[string][]string) (feats []Feat, err error) {
	sql := `SELECT * FROM feats`
	i := 0
	for o, arr := range opt {
		clause := "where"
		if i > 0 {
			clause = "and"
		}
		switch o {
		case "class":
			sql += fmt.Sprintf(" %s '%s' ILIKE class", clause, arr[0])
		case "subclass":
			sql += fmt.Sprintf(" %s ('%s' ILIKE subclass or subclass is null)", clause, arr[0])
		case "background":
			sql += fmt.Sprintf(" %s '%s' ILIKE background", clause, arr[0])
		case "level":
			sql += fmt.Sprintf(" %s level <= %s", clause, arr[0])
		}
		i++
	}
	sql += " order by level"
	fmt.Println(sql)
	rows, err := r.db.Query(sql + ";")
	if err != nil {
		return nil, fmt.Errorf("ERROR running query: `%s` (%s)", sql, err)
	}
	for rows.Next() {
		if feat, err := scanFeat(rows); err == nil {
			feats = append(feats, *feat)
		} else {
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
	}
	return
}

func scanFeat(row psql.Scannable) (feat *Feat, err error) {
	feat = new(Feat)
	if err := row.Scan(
		&feat.ID,
		&feat.Name,
		pq.Array(&feat.Ability),
		pq.Array(&feat.Description),
		&feat.Class,
		&feat.Subclass,
		&feat.Background,
		&feat.Level,
		&feat.Prereq); err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}
	return
}
