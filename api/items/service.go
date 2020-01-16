package items

// ItemService defines the outward functionality
// available to interact with Item data
type ItemService interface {
	FindByID(id string) (*Item, error)
	FindWeapons() ([]Item, error)
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
