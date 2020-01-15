package items

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// ItemRepo defines the necessary actions to
// interact with Item data
type ItemRepo interface {
	findByID(id string) (*Item, error)
}

type itemRepo struct {
	db *sql.DB
}

// NewNodeRepo returns a ItemRepo with a db connection
func NewNodeRepo(db *sql.DB) *itemRepo {
	return &itemRepo{
		db,
	}
}

func (r *itemRepo) findByID(id string) (item *Item, err error) {
	item = new(Item)
	sql := `SELECT * FROM items WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	if err := row.Scan(
		&item.ID,
		&item.Name,
		&item.Type,
		&item.Cost,
		&item.Weight,
		&item.Attune,
		&item.Rarity,
		&item.Weapon,
		&item.AC,
		pq.Array(&item.Info)); err != nil {
		return item, fmt.Errorf("Could not find row: %s", err)
	}
	return
}
