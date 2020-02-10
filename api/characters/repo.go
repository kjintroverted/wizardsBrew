package characters

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/lib/pq"

	"github.com/kjintroverted/wizardsBrew/data/tasks"
	"github.com/kjintroverted/wizardsBrew/psql"
)

const update string = `
UPDATE characters
SET
	name = '%s',
	auth_users = %s,
	read_users = %s,
	race_id = %d,
	class_id = %d,
	subclass = '%s',
	background = %d,
	stats = %s,
	xp = %d,
	max_hp = %d,
	hp = %d,
	init = %d,
	pro_skill = %s,
	pro_tool = %s,
	lang = %s,
	equip_ids = %s,
	weapon_ids = %s,
	inventory_ids = %s,
	gold = %f,
	spell_ids = %s,
	feat_ids = %s
WHERE id = %d
RETURNING id`

const insert string = `
INSERT INTO characters
(
	name,
	auth_users,
	read_users,
	race_id,
	class_id,
	subclass,
	background,
	stats,
	xp,
	max_hp,
	hp,
	init,
	pro_skill,
	pro_tool,
	lang,
	equip_ids,
	weapon_ids,
	inventory_ids,
	gold,
	spell_ids,
	feat_ids,
	owner
)
VALUES ('%v', %v, %v, %v, %v, '%v', %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, '%v')
RETURNING id`

// PCRepo defines the necessary actions to
// interact with PC data
type PCRepo interface {
	Upsert(data PC, uid string) (string, error)
	FindByUser(uid string) ([]PC, error)
	FindByID(id string) (*PC, error)
	Delete(id string) error
	Authorized(id, uid string) bool
}

type characterRepo struct {
	db *sql.DB
}

// NewPCRepo returns a PCRepo with a db connection
func NewPCRepo(db *sql.DB) PCRepo {
	return &characterRepo{
		db,
	}
}

func (r *characterRepo) FindByID(id string) (*PC, error) {
	sql := `SELECT * FROM characters WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanPC(row)
}

func (r *characterRepo) FindByUser(uid string) (arr []PC, err error) {
	sql := `SELECT * FROM characters WHERE owner=$1 or $1=any(auth_users)`
	rows, _ := r.db.Query(sql, uid)
	for rows.Next() {
		if pc, err := scanPC(rows); err == nil {
			arr = append(arr, *pc)
		}
	}
	return
}

func (r *characterRepo) Authorized(id, uid string) (auth bool) {
	sql := `SELECT id FROM characters WHERE id=$1 and (owner=$2 or $2=any(auth_users))`
	row := r.db.QueryRow(sql, id, uid)
	var x int
	err := row.Scan(&x)
	return err == nil
}

func (r *characterRepo) Delete(id string) error {
	sql := `DELETE FROM characters WHERE id=$1`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *characterRepo) Upsert(data PC, uid string) (string, error) {

	var skillArr []interface{}
	for _, skill := range data.ProSkills {
		row := tasks.RowString("skill", "'"+skill.Name+"'", skill.Mult)
		skillArr = append(skillArr, row)
	}

	var vals = []interface{}{
		data.Name,
		tasks.SimplerStrArray(data.AuthUsers),
		tasks.SimplerStrArray(data.ReadUsers),
		data.RaceID,
		data.ClassID,
		data.Subclass,
		data.BackgroundID,
		tasks.RowString("raw_stat", data.Stats.STR, data.Stats.DEX, data.Stats.CON, data.Stats.INT, data.Stats.WIS, data.Stats.CHA),
		data.XP,
		data.MaxHP,
		data.HP,
		data.Init,
		tasks.SimpleArray(skillArr),
		tasks.SimplerStrArray(data.ProTools),
		tasks.SimplerStrArray(data.Lang),
		tasks.SimpleArray(tasks.IntToIArray(data.EquipmentIDs)),
		tasks.SimpleArray(tasks.IntToIArray(data.WeaponIDs)),
		tasks.SimpleArray(tasks.IntToIArray(data.InventoryIDs)),
		data.Gold,
		tasks.SimpleArray(tasks.IntToIArray(data.SpellIDs)),
		tasks.SimpleArray(tasks.IntToIArray(data.SpecFeatIDs)),
	}

	var statement string
	if data.ID != 0 {
		vals = append(vals, data.ID)
		statement = update
	} else {
		vals = append(vals, uid)
		statement = insert
	}

	sql := fmt.Sprintf(statement, vals...)

	row := r.db.QueryRow(sql)

	var id int
	if err := row.Scan(&id); err != nil {
		fmt.Println("ERR running:", sql)
		return "", err
	}

	return strconv.Itoa(int(id)), nil
}

func scanPC(row psql.Scannable) (character *PC, err error) {
	character = new(PC)
	if err := row.Scan(
		&character.ID,
		&character.Name,
		&character.Owner,
		pq.Array(&character.AuthUsers),
		pq.Array(&character.ReadUsers),
		&character.RaceID,
		&character.ClassID,
		&character.Subclass,
		&character.BackgroundID,
		&character.Stats,
		&character.XP,
		&character.HP,
		&character.MaxHP,
		&character.Init,
		pq.Array(&character.ProSkills),
		pq.Array(&character.ProTools),
		pq.Array(&character.ProWeapons),
		pq.Array(&character.Lang),
		pq.Array(&character.EquipmentIDs),
		pq.Array(&character.WeaponIDs),
		pq.Array(&character.InventoryIDs),
		&character.Gold,
		pq.Array(&character.SpellIDs),
		pq.Array(&character.SpecFeatIDs)); err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}
	return
}
