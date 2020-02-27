package items

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/data/tasks"

	"github.com/kjintroverted/wizardsBrew/psql"

	"github.com/lib/pq"
)

const fields string = "id,name,type,cost,weight,attune,rarity,(weapon).category,(weapon).damage,(weapon).damage_type,armor_class,info"

// ItemRepo defines the necessary actions to
// interact with Item data
type ItemRepo interface {
	FindByID(id string) (*Item, error)
	FindByIDs(ids []psql.NullInt) ([]Item, error)
	FindWeapons() ([]Item, error)
	FindArmor() ([]Item, error)
	FindItems() ([]Item, error)
	InsertItem(item Item) error
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

func (r *itemRepo) FindByIDs(ids []psql.NullInt) (items []Item, err error) {
	sql := fmt.Sprintf(`select %s from items where id = any (array[%s])`, fields, tasks.JoinInt(ids, ","))
	if rows, err := r.db.Query(sql); err == nil {
		for rows.Next() {
			if item, err := scanItem(rows); err == nil {
				items = append(items, *item)
			} else {
				return nil, fmt.Errorf("Could not find row: %s", err)
			}
		}
	}
	return
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

func (r *itemRepo) InsertItem(item Item) (id int, err error) {
	sql := `INSERT INTO items 
						(name, type, cost, weight, attune, rarity, weapon, armor_class, info, homebrew) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, true)
					returning id
					`
	// Needs item fields
	row := r.db.QueryRow(sql)
	row.Scan(&id)

	return id, nil
}

func scanItem(row psql.Scannable) (item *Item, err error) {
	item = new(Item)
	weapon := new(weaponInfo)
	if err := row.Scan(
		&item.ID,
		&item.Name,
		&item.Type,
		&item.Cost,
		&item.Weight,
		&item.Attune,
		&item.Rarity,
		&weapon.Category,
		&weapon.Damage,
		&weapon.DamageType,
		&item.AC,
		pq.Array(&item.Info)); err != nil {
		return nil, fmt.Errorf("Could not find row: %s", err)
	}

	if weapon.Category != nil && weapon.Category.Valid {
		item.Weapon = weapon
	}

	return
}
