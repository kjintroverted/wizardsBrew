package nodes

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// NodeRepo defines the necessary actions to
// interact with Node data
type NodeRepo interface {
	findByID(id string) (*Node, error)
}

type nodeRepo struct {
	db *sql.DB
}

// NewNodeRepo returns a NodeRepo with a db connection
func NewNodeRepo(db *sql.DB) *nodeRepo {
	return &nodeRepo{
		db,
	}
}

func (r *nodeRepo) findByID(id string) (node *Node, err error) {
	node = new(Node)
	sql := `SELECT * FROM story_nodes WHERE id=$1`
	row := r.db.QueryRow(sql, id)
	if err := row.Scan(&node.ID, &node.Type, &node.Value, pq.Array(&node.Paths)); err != nil {
		return node, fmt.Errorf("Could not find row: %s", err)
	}
	return
}
