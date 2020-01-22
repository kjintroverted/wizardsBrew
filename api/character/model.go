package character

// PC describes the data for a Playable Character
type PC struct {
	ID           int        `json:"id" db:"id"`
	Owner        string     `json:"owner" db:"owner"`
	AuthUsers    []string   `json:"authUsers" db:"auth_users"`
	ReadUsers    []string   `json:"readUsers" db:"read_users"`
	Name         string     `json:"name" db:"name"`
	RaceID       int        `json:"raceID" db:"race_id"`
	ClassID      int        `json:"classID" db:"class_id"`
	Subclass     string     `json:"Subclass" db:"Subclass"`
	Stats        stats      `json:"stats" db:"stats"`
	XP           int        `json:"xp" db:"xp"`
	HP           int        `json:"hp" db:"hp"`
	MaxHP        int        `json:"maxHP" db:"maxHP"`
	Init         int        `json:"initiative" db:"initiative"`
	ProSkills    []proSkill `json:"proSkills" db:"proSkills"`
	ProTools     []string   `json:"proTools" db:"proTools"`
	Lang         []string   `json:"languages" db:"languages"`
	EquipmentIDs []int      `json:"equipmentIDs" db:"equipment_ids"`
	WeaponIDs    []int      `json:"weaponIDs" db:"weapon_ids"`
	InventoryIDs []int      `json:"inventoryIDs" db:"inventory_ids"`
	Gold         float64    `json:"gold" db:"gold"`
	SpellIDs     []int      `json:"spellIDs" db:"spell_ids"`
	SpecFeatIDs  []int      `json:"specFeatIDs" db:"specFeat_ids"`
}

type stats struct {
	STR int `json:"str" db:"str"`
	DEX int `json:"dex" db:"dex"`
	CON int `json:"con" db:"con"`
	INT int `json:"int" db:"int"`
	WIS int `json:"wis" db:"wis"`
	CHA int `json:"cha" db:"cha"`
}

type proSkill struct {
	Name string `json:"name" db:"name"`
	Mult int    `json:"multiplier" db:"multiplier"`
}
