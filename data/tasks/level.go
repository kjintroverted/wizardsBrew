package tasks

// GetLevelInfo returns relevant level data based on XP
func GetLevelInfo(xp int) (levelInfo level) {
	for i, level := range advancement {

		if i >= len(advancement)-1 {
			levelInfo = level
			break
		}

		nextLevel := advancement[i+1]
		if nextLevel.MinXP > xp {
			levelInfo = level
			levelInfo.NextXP = nextLevel.MinXP
			break
		}
	}
	return
}

type level struct {
	MinXP  int `json:"-"`
	Level  int `json:"level"`
	Bonus  int `json:"proBonus"`
	NextXP int `json:"next,omitempty"`
}

var advancement = []level{
	level{MinXP: 0, Level: 1, Bonus: 2},
	level{MinXP: 300, Level: 2, Bonus: 2},
	level{MinXP: 900, Level: 3, Bonus: 2},
	level{MinXP: 2700, Level: 4, Bonus: 2},
	level{MinXP: 6500, Level: 5, Bonus: 3},
	level{MinXP: 14000, Level: 6, Bonus: 3},
	level{MinXP: 23000, Level: 7, Bonus: 3},
	level{MinXP: 34000, Level: 8, Bonus: 3},
	level{MinXP: 48000, Level: 9, Bonus: 4},
	level{MinXP: 64000, Level: 10, Bonus: 4},
	level{MinXP: 85000, Level: 11, Bonus: 4},
	level{MinXP: 100000, Level: 12, Bonus: 4},
	level{MinXP: 120000, Level: 13, Bonus: 5},
	level{MinXP: 140000, Level: 14, Bonus: 5},
	level{MinXP: 165000, Level: 15, Bonus: 5},
	level{MinXP: 195000, Level: 16, Bonus: 5},
	level{MinXP: 225000, Level: 17, Bonus: 6},
	level{MinXP: 265000, Level: 18, Bonus: 6},
	level{MinXP: 305000, Level: 19, Bonus: 6},
	level{MinXP: 355000, Level: 20, Bonus: 6},
}
