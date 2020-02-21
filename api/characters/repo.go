package characters

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/kjintroverted/wizardsBrew/data/tasks"
	"github.com/kjintroverted/wizardsBrew/psql"
)

const update string = `
UPDATE characters
SET
	name = '%v',
	photo_url = '%v',
	auth_users = %v,
	auth_req = %v,
	party_inv = %v,
	race_id = %v,
	class_id = %v,
	subclass = '%v',
	background = %v,
	stats = %v,
	xp = %v,
	max_hp = %v,
	hp = %v,
	init = %v,
	pro_skill = %v,
	pro_tool = %v,
	lang = %v,
	equip_ids = %v,
	weapon_ids = %v,
	inventory_ids = %v,
	gold = %v,
	spell_ids = %v,
	feat_ids = %v
WHERE id = '%v'
RETURNING id`

const insert string = `
INSERT INTO characters
(
	name,
	photo_url,
	auth_users,
	auth_req,
	party_inv,
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
	owner,
	id
)
VALUES ('%v', '%v', %v, %v, %v, %v, %v, '%v', %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, '%v', '%v')
`

// PCRepo defines the necessary actions to
// interact with PC data
type PCRepo interface {
	Upsert(data PC, uid string) (string, error)
	FindByUser(uid string) ([]PC, error)
	FindByID(id string) (*PC, error)
	RequestAuth(id, uid string) error
	Invite(id, party string) error
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
	var s string
	err := row.Scan(&s)
	return err == nil
}

func (r *characterRepo) Delete(id string) error {
	sql := `DELETE FROM characters WHERE id=$1`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *characterRepo) RequestAuth(id, uid string) error {
	sql := `UPDATE characters SET auth_req = array_append(auth_req, $1) WHERE id = $2`
	_, err := r.db.Exec(sql, uid, id)
	return err
}

func (r *characterRepo) Invite(id, party string) error {
	sql := `UPDATE characters SET party_inv = array_append(party_inv, $1) WHERE id = $2`
	_, err := r.db.Exec(sql, party, id)
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
		data.PhotoURL,
		tasks.SimplerStrArray(data.AuthUsers),
		tasks.SimplerStrArray(data.AuthReq),
		tasks.SimplerStrArray(data.PartyInv),
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

	id := data.ID
	statement := update
	if id == "" {
		id = psql.GetUID()
		vals = append(vals, uid)
		statement = insert
	}
	vals = append(vals, id)

	sql := fmt.Sprintf(statement, vals...)

	_, err := r.db.Exec(sql)
	if err != nil {
		fmt.Println("ERR running:", sql)
		return "", errors.WithStack(err)
	}

	return id, nil
}

func scanPC(row psql.Scannable) (character *PC, err error) {
	character = new(PC)
	if err := row.Scan(
		&character.ID,
		&character.Name,
		&character.PhotoURL,
		&character.Owner,
		pq.Array(&character.AuthUsers),
		pq.Array(&character.AuthReq),
		pq.Array(&character.PartyInv),
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
