package characters

// PCService defines the outward functionality
// available to interact with PC data
type PCService interface {
	Upsert(data PC, uid string) (string, error)
	FindByID(id string) (*PC, error)
	Delete(id string) error
	Authorized(id, uid string) bool
	AuthorizedLocal(pc PC, uid string) bool
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
func (s *pcService) Upsert(data PC, uid string) (id string, err error) {
	return s.repo.Upsert(data, uid)
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

// Authorized will call the repo's function to
// check a users permissions for a character
func (s *pcService) AuthorizedLocal(pc PC, uid string) bool {
	if pc.Owner == uid {
		return true
	}
	for _, u := range pc.AuthUsers {
		if u == uid {
			return true
		}
	}
	return false
}

// Delete will call the repo's function to
// find a PC by ID
func (s *pcService) Delete(id string) error {
	return s.repo.Delete(id)
}
