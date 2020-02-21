package characters

// PCService defines the outward functionality
// available to interact with PC data
type PCService interface {
	Upsert(data PC, uid string) (string, error)
	FindByUser(uid string) ([]PC, error)
	Search(q string) ([]PC, error)
	FindByID(id string) (*PC, error)
	RequestAuth(id, uid string) error
	Invite(id, party string) error
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

// FindByUser will call the repo's function to
// find a PC by ID
func (s *pcService) FindByUser(uid string) (pc []PC, err error) {
	pc, err = s.repo.FindByUser(uid)
	return pc, err
}

// Search will try to match character names to the query (q)
func (s *pcService) Search(q string) (pc []PC, err error) {
	return s.repo.Search(q)
}

// Authorized will call the repo's function to
// check a users permissions for a character
func (s *pcService) Authorized(id, uid string) bool {
	return s.repo.Authorized(id, uid)
}

// RequestAuth will call the repo's function to
// check a users permissions for a character
func (s *pcService) RequestAuth(id, uid string) error {
	return s.repo.RequestAuth(id, uid)
}

// Invite will call the repo's function to
// check a users permissions for a character
func (s *pcService) Invite(id, party string) error {
	return s.repo.Invite(id, party)
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
