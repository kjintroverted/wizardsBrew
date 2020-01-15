package psql

import (
	"database/sql"
	"fmt"

	"github.com/kjintroverted/wizardsBrew/nodes"

	"github.com/lib/pq"
	_ "github.com/lib/pq" // access to psql driver
)

type nodeRepo struct {
	db *sql.DB
}

// NewNodeRepo creates an instance of the PSQL db connection
func NewNodeRepo(db *sql.DB) *nodeRepo {
	return &nodeRepo{
		db: db,
	}
}

func (r *nodeRepo) FindByID(id string) (node *nodes.Node, err error) {
	node = new(nodes.Node)
	sql := `SELECT * FROM story_nodes WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	if err := row.Scan(&node.ID, &node.Type, &node.Value, pq.Array(&node.Paths)); err != nil {
		return node, fmt.Errorf("Could not find row: %s", err)
	}
	return
}
