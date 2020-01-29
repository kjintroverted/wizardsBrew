package feats

import (
	"github.com/kjintroverted/wizardsBrew/psql"
)

// FeatService defines the outward functionality
// available to interact with Feat data
type FeatService interface {
	FindByID(id string) (*Feat, error)
	FindByIDs(ids []psql.NullInt) ([]Feat, error)
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

// FindByIds will call the repo's function to
// find a Feat by ID
func (s *featService) FindByIDs(ids []psql.NullInt) (feat []Feat, err error) {
	feat, err = s.repo.FindByIDs(ids)
	return feat, err
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
