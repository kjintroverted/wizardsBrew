package items

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// ItemRepo defines the necessary actions to
// interact with Item data
type ItemRepo interface {
	FindByID(id string) (*Item, error)
	FindWeapons() ([]Item, error)
	// findArmor() (*[]Item, error)
	// findGear() (*[]Item, error)
}

type itemRepo struct {
	db *sql.DB
}

// NewItemRepo returns a ItemRepo with a db connection
func NewItemRepo(db *sql.DB) ItemRepo {
	return &itemRepo{
		db,
	}
}

func (r *itemRepo) FindByID(id string) (item *Item, err error) {
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

func (r *itemRepo) FindWeapons() (items []Item, err error) {
	sql := `select * from items where type ILIKE '%weapon%'`
	rows, err := r.db.Query(sql)
	for rows.Next() {
		item := new(Item)
		if err := rows.Scan(
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
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
		items = append(items, *item)
	}
	return
}
