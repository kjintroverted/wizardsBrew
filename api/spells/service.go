package spells

// SpellService defines the outward functionality
// available to interact with Spell data
type SpellService interface {
	FindByID(id string) (*Spell, error)
	List(opt map[string][]string) ([]Spell, error)
}

type spellService struct {
	repo SpellRepo
}

// NewSpellService returns a new implementation of SpellService
func NewSpellService(repo SpellRepo) SpellService {
	return &spellService{
		repo: repo,
	}
}

// FindById will call the repo's function to
// find a Spell by ID
func (s *spellService) FindByID(id string) (spell *Spell, err error) {
	spell, err = s.repo.FindByID(id)
	return spell, err
}

// List will call the repo's function to
// list all Spells that meet the optional criteria
func (s *spellService) List(opt map[string][]string) (spells []Spell, err error) {
	return s.repo.List(opt)
}
