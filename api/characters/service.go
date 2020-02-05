package characters

// PCService defines the outward functionality
// available to interact with PC data
type PCService interface {
	Upsert(data PC) (string, error)
	FindByID(id string) (*PC, error)
	Delete(id string) error
	Authorized(id, uid string) bool
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

// Authorized will call the repo's function to
// check a users permissions for a character
func (s *pcService) Authorized(id, uid string) bool {
	return s.repo.Authorized(id, uid)
}

// Delete will call the repo's function to
// find a PC by ID
func (s *pcService) Delete(id string) error {
	return s.repo.Delete(id)
}
