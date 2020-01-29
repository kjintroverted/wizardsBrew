package characters

// PCService defines the outward functionality
// available to interact with PC data
type PCService interface {
	Upsert(data PC) (string, error)
	FindByID(id string) (*PC, error)
}

type pcService struct {
	repo PCRepo
}

// NewPCService returns a new implementation of PCService
func NewPCService(repo PCRepo) PCService {
	return &pcService{
		repo: repo,
	}
}

// Upsert will call the repo's function to
// insert or update a character
func (s *pcService) Upsert(data PC) (id string, err error) {
	id, err = s.repo.Upsert(data)
	return id, err
}

// FindByID will call the repo's function to
// find a PC by ID
func (s *pcService) FindByID(id string) (pc *PC, err error) {
	pc, err = s.repo.FindByID(id)
	return pc, err
}
