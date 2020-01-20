package races

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"
	"github.com/lib/pq"
)

// RaceRepo defines the necessary actions to
// interact with Race data
type RaceRepo interface {
	List() ([]Race, error)
	FindByID(id string) (*Race, error)
}

type raceRepo struct {
	db *sql.DB
}

// NewRaceRepo returns a RaceRepo with a db connection
func NewRaceRepo(db *sql.DB) RaceRepo {
	return &raceRepo{
		db,
	}
}

func (r *raceRepo) FindByID(id string) (race *Race, err error) {
	sql := `SELECT * FROM races WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanRace(row)
}

func (r *raceRepo) List() (races []Race, err error) {
	sql := `SELECT * FROM races`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("ERROR running query: `%s` (%s)", sql, err)
	}
	for rows.Next() {
		race, _ := scanRace(rows)
		races = append(races, *race)
	}
	return
}

func scanRace(row psql.Scannable) (race *Race, err error) {
	race = new(Race)
	if err := row.Scan(
		&race.ID,
		&race.Name,
		pq.Array(&race.Ability),
		&race.Size,
		&race.Speed,
		&race.Age,
		&race.Align,
		&race.SizeDesc,
		pq.Array(&race.Traits)); err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}
	return
}
