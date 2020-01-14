package nodes

// NodeRepo defines the necessary actions to
// interact with Node data
type NodeRepo interface {
	FindByID(id int) (*Node, error)
}
