package backgrounds

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/psql"
	"github.com/lib/pq"
)

const fields string = "id,name,pro_skill,pro_tool,language,equipment,special_opts"

// BackgroundRepo defines the necessary actions to
// interact with Background data
type BackgroundRepo interface {
	List() ([]Background, error)
	FindByID(id string) (*Background, error)
}

type backgroundRepo struct {
	db *sql.DB
}

// NewBackgroundRepo returns a BackgroundRepo with a db connection
func NewBackgroundRepo(db *sql.DB) BackgroundRepo {
	return &backgroundRepo{
		db,
	}
}

func (r *backgroundRepo) FindByID(id string) (background *Background, err error) {
	sql := fmt.Sprintf(`SELECT %s FROM backgrounds WHERE id=$1`, fields)
	row := r.db.QueryRow(sql, id)
	return scanBackground(row)
}

func (r *backgroundRepo) List() (backgrounds []Background, err error) {
	sql := fmt.Sprintf(`SELECT %s FROM backgrounds`, fields)
	sql += " where name !~~* 'Variant%'" // IGNORE VARIANTS
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("ERROR running query: `%s` (%s)", sql, err)
	}
	for rows.Next() {
		background, _ := scanBackground(rows)
		backgrounds = append(backgrounds, *background)
	}
	return
}

func scanBackground(row psql.Scannable) (background *Background, err error) {
	background = new(Background)
	if err := row.Scan(
		&background.ID,
		&background.Name,
		pq.Array(&background.ProSkill),
		pq.Array(&background.ProTool),
		pq.Array(&background.Language),
		pq.Array(&background.Equipment),
		pq.Array(&background.SpecialOpts)); err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}
	return
}
