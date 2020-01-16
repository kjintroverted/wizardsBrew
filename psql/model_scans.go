package psql

import (
	"database/sql"
	"encoding/json"
	"strings"
)

// Description is a generic type used to store
// item/character descriptions
type Description struct {
	Title string   `json:"title" db:"title"`
	Body  []string `json:"body" db:"body"`
}

// Scan informs the sql driver how to Unmarshal Desscriptions
func (d *Description) Scan(value interface{}) (err error) {
	str := string(value.([]byte))
	vals := strings.Split(strings.Trim(str, "()"), ",\"")

	d.Title = strings.Trim(vals[0], "\"")

	var p []string
	for _, s := range vals[1:] {
		p = append(p, strings.Trim(s, "{}\""))
	}
	d.Body = p

	return
}

// Scannable is an abstraction of a
// row to be scanned to all for more flexible handlers
type Scannable interface {
	Scan(dest ...interface{}) error
}

// NullFloat is an alias fot sql nullable float
type NullFloat struct {
	sql.NullFloat64
}

// MarshalJSON unwraps the valid value when writing to JSON
func (n *NullFloat) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Float64)
}

// NullInt is an alias fot sql nullable float
type NullInt struct {
	sql.NullInt64
}

// MarshalJSON unwraps the valid value when writing to JSON
func (n *NullInt) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Int64)
}

// NullString is an alias fot sql nullable float
type NullString struct {
	sql.NullString
}

// MarshalJSON unwraps the valid value when writing to JSON
func (n *NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.String)
}
