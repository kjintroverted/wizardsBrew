package nodes

import (
	"strings"
)

// Node defines a story or quiz node
type Node struct {
	ID    int    `json:"id" db:"id"`
	Type  string `json:"type" db:"type"`
	Value string `json:"value" db:"value"`
	Paths []path `json:"paths" db:"paths"`
}

type path struct {
	Prompt string `json:"prompt" db:"prompt"`
	Value  string `json:"value" db:"value"`
}

func (p *path) Scan(value interface{}) error {
	str := string(value.([]byte))
	vals := strings.Split(strings.Trim(str, "()"), ",")

	p.Prompt = strings.Trim(vals[0], "\"")
	p.Value = strings.Trim(vals[1], "\"")

	return nil
}
