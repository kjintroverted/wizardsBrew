package parties

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/kjintroverted/wizardsBrew/psql"
)

const update string = `
UPDATE parties
SET
	name = '%v',
	photo_url = '%v',
	admin = '%v'
WHERE id = '%v'
AND admin = '%v'`

const insert string = `
INSERT INTO parties
(
	name,
	photo_url,
	admin,
	id
)
VALUES ('%v', '%v', '%v', '%v')
`

// PartyRepo defines the necessary actions to
// interact with Party data
type PartyRepo interface {
	Upsert(data Party, uid string) (string, error)
	// AddMember(id string) error
	// KickMember(id string) error
	FindByMember(id string) ([]Party, error)
	FindByID(id string) (*Party, error)
	Delete(id string, uid string) error
}

type partyRepo struct {
	db *sql.DB
}

// NewPartyRepo returns a PartyRepo with a db connection
func NewPartyRepo(db *sql.DB) PartyRepo {
	return &partyRepo{
		db,
	}
}

func (r *partyRepo) FindByID(id string) (*Party, error) {
	sql := `SELECT * FROM parties WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	return scanParty(row)
}

func (r *partyRepo) FindByMember(id string) (arr []Party, err error) {
	sql := `SELECT * FROM parties WHERE admin=$1 or $1=any(members)`
	rows, _ := r.db.Query(sql, id)
	for rows.Next() {
		if pc, err := scanParty(rows); err == nil {
			arr = append(arr, *pc)
		}
	}
	return
}

func (r *partyRepo) Delete(id string, uid string) error {
	sql := `DELETE FROM parties WHERE id=$1 and admin = $2`
	_, err := r.db.Exec(sql, id, uid)
	return err
}

func (r *partyRepo) Upsert(data Party, uid string) (string, error) {

	admin := data.Admin
	if admin == "" {
		admin = uid
	}

	var vals = []interface{}{
		data.Name,
		data.PhotoURL,
		admin,
	}

	var statement string
	id := data.ID
	if id != "" {
		statement = update
		vals = append(vals, id, uid)
	} else {
		statement = insert
		id = psql.GetUID()
		vals = append(vals, id)
	}

	sql := fmt.Sprintf(statement, vals...)

	_, err := r.db.Exec(sql)
	if err != nil {
		fmt.Println("ERR running:", sql)
		return "", errors.WithStack(err)
	}

	return id, nil
}

func scanParty(row psql.Scannable) (party *Party, err error) {
	party = new(Party)
	if err := row.Scan(
		&party.ID,
		&party.Name,
		&party.PhotoURL,
		&party.Admin,
		pq.Array(&party.Members)); err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}
	return
}
