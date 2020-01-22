package feats

// FeatService defines the outward functionality
// available to interact with Feat data
type FeatService interface {
	FindByID(id string) (*Feat, error)
	List(opt map[string][]string) ([]Feat, error)
}

type featService struct {
	repo FeatRepo
}

// NewFeatService returns a new implementation of FeatService
func NewFeatService(repo FeatRepo) FeatService {
	return &featService{
		repo: repo,
	}
}

// FindById will call the repo's function to
// find a Feat by ID
func (s *featService) FindByID(id string) (feat *Feat, err error) {
	feat, err = s.repo.FindByID(id)
	return feat, err
}

// List will call the repo's function to
// list all Feats that meet the optional criteria
func (s *featService) List(opt map[string][]string) (feats []Feat, err error) {
	return s.repo.List(opt)
}
