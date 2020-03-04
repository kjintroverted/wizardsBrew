package items

import "github.com/kjintroverted/wizardsBrew/psql"

// ItemService defines the outward functionality
// available to interact with Item data
type ItemService interface {
	FindByID(id string) (*Item, error)
	FindByIDs(ids []psql.NullInt) ([]Item, error)
	FindWeapons() ([]Item, error)
	FindArmor() ([]Item, error)
	FindItems() ([]Item, error)
	InsertItem(item Item) (int, error)
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

// FindByIDs will call the repo's function to
// find all weapon items
func (s *itemService) FindByIDs(arr []psql.NullInt) ([]Item, error) {
	items, err := s.repo.FindByIDs(arr)
	return items, err
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

// FindItems will call the repo's function to
// find all armor items
func (s *itemService) FindItems() ([]Item, error) {
	items, err := s.repo.FindItems()
	return items, err
}

func (s *itemService) InsertItem(i Item) (int, error) {
	return s.repo.InsertItem(i)
}
