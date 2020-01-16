package items

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"

	"github.com/lib/pq"
)

// ItemRepo defines the necessary actions to
// interact with Item data
type ItemRepo interface {
	FindByID(id string) (*Item, error)
	FindWeapons() ([]Item, error)
	FindArmor() ([]Item, error)
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
	sql := `SELECT * FROM items WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanItem(row)
}

func (r *itemRepo) FindWeapons() (items []Item, err error) {
	sql := `select * from items where weapon is not null`
	rows, err := r.db.Query(sql)
	for rows.Next() {
		if item, err := scanItem(rows); err == nil {
			items = append(items, *item)
		} else {
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
	}
	return
}

func (r *itemRepo) FindArmor() (items []Item, err error) {
	sql := `select * from items where armor_class is not null`
	rows, err := r.db.Query(sql)
	for rows.Next() {
		if item, err := scanItem(rows); err == nil {
			items = append(items, *item)
		} else {
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
	}
	return
}

func scanItem(row psql.Scannable) (item *Item, err error) {
	item = new(Item)
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
		return nil, fmt.Errorf("Could not find row: %s", err)
	}
	return
}
