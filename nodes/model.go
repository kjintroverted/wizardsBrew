package nodes

// Node defines a story or quiz node
type Node struct {
	ID    int    `json:"id" db:"id"`
	Type  string `json:"type" db:"type"`
	Value string `json:"value" db:"value"`
	Paths []struct {
		Prompt string `json:"prompt" db:"prompt"`
		Value  string `json:"value" db:"value"`
	} `json:"paths" db:"paths"`
}
