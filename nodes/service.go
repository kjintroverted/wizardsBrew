package nodes

// NodeService defines the outward functionality
// available to interact with Node data
type NodeService interface {
	FindByID(id int) (*Node, error)
}

type nodeService struct {
	repo NodeRepo
}

// NewNodeService returns a new implementation of NodeService
func NewNodeService(repo NodeRepo) NodeService {
	return &nodeService{
		repo: repo,
	}
}

// FindById will call the repo's function to
// find a Node by ID
func (s *nodeService) FindByID(id int) (*Node, error) {
	return s.repo.FindByID(id)
}
