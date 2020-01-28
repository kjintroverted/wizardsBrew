package characters

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/kjintroverted/wizardsBrew/data/tasks"
)

const update string = `
UPDATE characters
SET
	name = '%s',
	owner = '%s',
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
	owner,
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
	feat_ids
)
VALUES ('%s', '%s', %s, %s, %d, %d, '%s', %d, %s, %d, %d, %d, %d, %s, %s, %s, %s, %s, %s, %f, %s, %s)
RETURNING id`

// PCRepo defines the necessary actions to
// interact with PC data
type PCRepo interface {
	Upsert(data PC) (string, error)
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

func (r *characterRepo) Upsert(data PC) (string, error) {

	var skillArr []interface{}
	for _, skill := range data.ProSkills {
		row := tasks.RowString("skill", "'"+skill.Name+"'", skill.Mult)
		skillArr = append(skillArr, row)
	}

	var vals = []interface{}{
		data.Name,
		data.Owner,
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

	statement := insert
	if data.ID != 0 {
		vals = append(vals, data.ID)
		statement = update
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
