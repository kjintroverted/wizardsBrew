package items

// ItemService defines the outward functionality
// available to interact with Item data
type ItemService interface {
	FindByID(id string) (*Item, error)
	FindWeapons() ([]Item, error)
	FindArmor() ([]Item, error)
}

type itemService struct {
	repo ItemRepo
}

// NewItemService returns a new implementation of ItemService
func NewItemService(repo ItemRepo) ItemService {
	return &itemService{
		repo: repo,
	}
}

// FindById will call the repo's function to
// find a Item by ID
func (s *itemService) FindByID(id string) (item *Item, err error) {
	item, err = s.repo.FindByID(id)
	return item, err
}

// FindWeapons will call the repo's function to
// find all weapon items
func (s *itemService) FindWeapons() ([]Item, error) {
	items, err := s.repo.FindWeapons()
	return items, err
}

// FindArmor will call the repo's function to
// find all armor items
func (s *itemService) FindArmor() ([]Item, error) {
	items, err := s.repo.FindArmor()
	return items, err
}
