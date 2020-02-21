package parties

// PartyService defines the outward functionality
// available to interact with Party data
type PartyService interface {
	Upsert(data Party, uid string) (string, error)
	FindByMember(uid string) ([]Party, error)
	FindByID(id string) (*Party, error)
	Delete(id, uid string) error
}

type partyService struct {
	repo PartyRepo
}

// NewPartyService returns a new implementation of PartyService
func NewPartyService(repo PartyRepo) PartyService {
	return &partyService{
		repo: repo,
	}
}

// Upsert will call the repo's function to
// insert or update a character
func (s *partyService) Upsert(data Party, uid string) (id string, err error) {
	return s.repo.Upsert(data, uid)
}

// FindByID will call the repo's function to
// find a Party by ID
func (s *partyService) FindByID(id string) (party *Party, err error) {
	party, err = s.repo.FindByID(id)
	return party, err
}

// FindByMember will call the repo's function to
// find all Parties where the specified user id is owner or member
func (s *partyService) FindByMember(uid string) (party []Party, err error) {
	party, err = s.repo.FindByMember(uid)
	return party, err
}

// Delete will call the repo's function to
// find a Party by ID
func (s *partyService) Delete(id, uid string) error {
	return s.repo.Delete(id, uid)
}