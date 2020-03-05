package items

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/data/tasks"
	"github.com/pkg/errors"

	"github.com/kjintroverted/wizardsBrew/psql"

	"github.com/lib/pq"
)

const fields string = "id,name,type,cost,weight,attune,rarity,(weapon).category,(weapon).damage,(weapon).damage_type,armor_class,info,homebrew"

// ItemRepo defines the necessary actions to
// interact with Item data
type ItemRepo interface {
	FindByID(string) (*Item, error)
	FindByIDs([]psql.NullInt) ([]Item, error)
	FindWeapons(string) ([]Item, error)
	FindArmor(string) ([]Item, error)
	FindItems(string) ([]Item, error)
	InsertItem(Item) (int, error)
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

func (r *itemRepo) FindWeapons(q string) (items []Item, err error) {
	sql := `select ` + fields + ` from items where weapon is not null and name ilike '%` + q + `%'`
	rows, err := r.db.Query(sql)
	if err != nil {
		fmt.Printf("%v\n%+v\n", sql, errors.WithStack(err))
		return nil, err
	}
	for rows.Next() {
		if item, err := scanItem(rows); err == nil {
			items = append(items, *item)
		} else {
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
	}
	return
}

func (r *itemRepo) FindArmor(q string) (items []Item, err error) {
	sql := `select ` + fields + ` from items where armor_class is not null and name ilike '%` + q + `%'`
	rows, err := r.db.Query(sql)
	if err != nil {
		fmt.Printf("%v\n%+v\n", sql, errors.WithStack(err))
		return nil, err
	}
	for rows.Next() {
		if item, err := scanItem(rows); err == nil {
			items = append(items, *item)
		} else {
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
	}
	return
}

func (r *itemRepo) FindItems(q string) (items []Item, err error) {
	sql := `select ` + fields + ` from items where armor_class is null and weapon is null and name ilike '%` + q + `%'`
	rows, err := r.db.Query(sql)
	if err != nil {
		fmt.Printf("%v\n%+v\n", sql, errors.WithStack(err))
		return nil, err
	}
	for rows.Next() {
		if item, err := scanItem(rows); err == nil {
			items = append(items, *item)
		} else {
			return nil, fmt.Errorf("Could not find row: %s", err)
		}
	}
	return
}

func (r *itemRepo) Search(q string) (arr []Item, err error) {
	sql := `SELECT ` + fields + ` FROM characters WHERE name ilike '%` + q + `%'`

	rows, _ := r.db.Query(sql)
	for rows.Next() {
		if item, err := scanItem(rows); err == nil {
			arr = append(arr, *item)
		}
	}
	return
}

func (r *itemRepo) InsertItem(item Item) (id int, err error) {

	var infoArr []interface{}
	for _, section := range item.Info {
		infoArr = append(infoArr, fmt.Sprintf("row('%v',%v)::section", section.Title, tasks.SimplerStrArray(section.Body)))
	}

	sql := fmt.Sprintf(`
		INSERT INTO items 
			(name, type, cost, weight, attune, rarity, weapon, armor_class, info, homebrew) 
		VALUES ($1, $2, $3, $4, $5, $6, %v, $7, %v, true)
		returning id
		`,
		tasks.StructRow("weapon_info", item.Weapon),
		tasks.SimpleArray(infoArr))

	row := r.db.QueryRow(sql,
		item.Name,
		item.Type,
		item.Cost,
		item.Weight,
		item.Attune.Value(),
		item.Rarity.Value(),
		item.AC,
	)
	err = row.Scan(&id)

	return id, errors.WithStack(err)
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
		pq.Array(&item.Info),
		&item.IsHomebrew); err != nil {
		return nil, fmt.Errorf("Could not find row: %s", err)
	}

	if weapon.Category != nil && weapon.Category.Valid {
		item.Weapon = weapon
	}

	return
}
