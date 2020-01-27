package character

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/kjintroverted/wizardsBrew/data/tasks"
)

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

func (r *characterRepo) Upsert(data PC) (id string, err error) {

	var skillArr []string
	for _, skill := range data.ProSkills {
		row := tasks.RowString("skill", "'"+skill.Name+"'", skill.Mult)
	}

	sql := fmt.Sprintf(`INSERT INTO characters
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
	VALUES ('%s', '%s', %s, %s, %d, %d, '%s', '%s', %d, %d, %d, %d, %s, %s, %s, %s, %s, %s, %s, %d, %s, %s)`,
		data.Name,
		data.Owner,
		tasks.SimplerStrArray(data.AuthUsers),
		tasks.SimplerStrArray(data.ReadUsers),
		data.RaceID,
		data.ClassID,
		data.Subclass,
		data.Background,
		tasks.RowString("raw_stat", data.Stats.STR, data.Stats.DEX, data.Stats.CON, data.Stats.INT, data.Stats.WIS, data.Stats.CHA),
		data.XP,
		data.MaxHP,
		data.HP,
		data.Init,
		tasks.SimpleArray(skillArr),
		tasks.SimplerStrArray(data.ProTools),
		tasks.SimplerStrArray(data.Lang),
		tasks.SimpleArray(data.EquipmentIDs),
		tasks.SimpleArray(data.WeaponIDs),
		tasks.SimpleArray(data.InventoryIDs),
		data.Gold,
		tasks.SimpleArray(data.SpellIDs),
		tasks.SimpleArray(data.SpecFeatIDs),
	)

	fmt.Printf("RUNNING INSERT=========\n %s\n\n", sql)

	res, err := r.db.Exec(sql)
	if err != nil {
		return "", err
	}

	rawID, err := res.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(rawID)), nil
}
