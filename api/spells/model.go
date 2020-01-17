package spells

import (
	"strconv"
	"strings"

	"github.com/kjintroverted/wizardsBrew/psql"
)

// Spell defines a spell, prereqs, and all components
type Spell struct {
	ID          string             `json:"id" db:"id"`
	Name        string             `json:"name" db:"name"`
	School      string             `json:"school" db:"school"`
	Time        spellTime          `json:"time" db:"time"`
	Duration    string             `json:"duration" db:"duration"`
	Components  []component        `json:"comp" db:"comp"`
	Concentrate bool               `json:"concentrate" db:"concentrate"`
	Range       string             `json:"range" db:"range"`
	Level       int                `json:"level" db:"level"`
	Class       []string           `json:"class" db:"class"`
	Description []psql.Description `json:"description" db:"description"`
}

type spellTime struct {
	Time      string `json:"time" db:"time"`
	Condition string `json:"condition,omitempty" db:"condition,omitempty"`
}

func (t *spellTime) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}
	str := string(value.([]byte))
	vals := strings.Split(strings.Trim(str, "()"), "\",")

	t.Time = strings.Trim(vals[0], "\"")
	t.Condition = strings.Trim(vals[1], "\"")
	return
}

type component struct {
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Cost        float64 `json:"cost,omitempty" db:"cost,omitempty"`
	Consume     bool    `json:"consume,omitempty" db:"consume,omitempty"`
}

func (c *component) Scan(value interface{}) (err error) {
	if value == nil {
		return
	}
	str := string(value.([]byte))
	vals := psql.ParseRow(str)

	c.Name = vals[0]
	c.Description = vals[1]
	if vals[2] != "" {
		c.Cost, _ = strconv.ParseFloat(vals[2], 64)
	}
	if vals[3] != "" {
		c.Consume, _ = strconv.ParseBool(vals[3])
	}
	return
}
