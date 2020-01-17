package items

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"

	"github.com/lib/pq"
)

const fields string = "id,name,type,cost,weight,attune,rarity,(weapon).category,(weapon).damage,(weapon).damage_type,armor_class,info"

// ItemRepo defines the necessary actions to
// interact with Item data
type ItemRepo interface {
	FindByID(id string) (*Item, error)
	FindWeapons() ([]Item, error)
	FindArmor() ([]Item, error)
	FindItems() ([]Item, error)
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
	sql := fmt.Sprintf(`SELECT %s FROM items WHERE id=$1`, fields)
	row := r.db.QueryRow(sql, id)
	return scanItem(row)
}

func (r *itemRepo) FindWeapons() (items []Item, err error) {
	sql := fmt.Sprintf(`select %s from items where weapon is not null`, fields)
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
	sql := fmt.Sprintf(`select %s from items where armor_class is not null`, fields)
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

func (r *itemRepo) FindItems() (items []Item, err error) {
	sql := fmt.Sprintf(`select %s from items where armor_class is null and weapon is null`, fields)
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
	item.Weapon = new(weaponInfo)
	if err := row.Scan(
		&item.ID,
		&item.Name,
		&item.Type,
		&item.Cost,
		&item.Weight,
		&item.Attune,
		&item.Rarity,
		&item.Weapon.Category,
		&item.Weapon.Damage,
		&item.Weapon.DamageType,
		&item.AC,
		pq.Array(&item.Info)); err != nil {
		return nil, fmt.Errorf("Could not find row: %s", err)
	}
	return
}
