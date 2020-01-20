package classes

// ClassService defines the outward functionality
// available to interact with Class data
type ClassService interface {
	List() ([]Class, error)
	FindByID(id string) (*Class, error)
}

type classService struct {
	repo ClassRepo
}

// NewClassService returns a new implementation of ClassService
func NewClassService(repo ClassRepo) ClassService {
	return &classService{
		repo: repo,
	}
}

// FindById will call the repo's function to
// find a Class by ID
func (s *classService) FindByID(id string) (class *Class, err error) {
	class, err = s.repo.FindByID(id)
	return class, err
}

// FindById will call the repo's function to
// find a Class by ID
func (s *classService) List() ([]Class, error) {
	return s.repo.List()
}
