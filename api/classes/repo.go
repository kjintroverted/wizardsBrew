package classes

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"
	"github.com/lib/pq"
)

// ClassRepo defines the necessary actions to
// interact with Class data
type ClassRepo interface {
	FindByID(id string) (*Class, error)
	List() ([]Class, error)
}

type classRepo struct {
	db *sql.DB
}

// NewClassRepo returns a ClassRepo with a db connection
func NewClassRepo(db *sql.DB) ClassRepo {
	return &classRepo{
		db,
	}
}

func (r *classRepo) FindByID(id string) (class *Class, err error) {
	sql := `SELECT * FROM classes WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanClass(row)
}

func (r *classRepo) List() (classes []Class, err error) {
	sql := `SELECT * FROM classes`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("ERROR running query: `%s` (%s)", sql, err)
	}
	for rows.Next() {
		class, err := scanClass(rows)
		if err != nil {
			return nil, fmt.Errorf("ERROR running query: `%s` (%s)", sql, err)
		}
		classes = append(classes, *class)
	}
	return
}

func scanClass(row psql.Scannable) (class *Class, err error) {
	class = new(Class)
	if err := row.Scan(
		&class.ID,
		&class.Name,
		&class.HitDice,
		pq.Array(&class.ProArmor),
		pq.Array(&class.ProWeapon),
		&class.ProTool,
		pq.Array(&class.ProSave),
		&class.Skills,
		pq.Array(&class.StartEquipment),
		pq.Array(&class.Description),
		&class.Progress); err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}
	return
}
