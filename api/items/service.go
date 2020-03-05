package items

import "github.com/kjintroverted/wizardsBrew/psql"

// ItemService defines the outward functionality
// available to interact with Item data
type ItemService interface {
	FindByID(id string) (*Item, error)
	FindByIDs(ids []psql.NullInt) ([]Item, error)
	FindWeapons(string) ([]Item, error)
	FindArmor(string) ([]Item, error)
	FindItems(string) ([]Item, error)
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
func (s *itemService) FindByID(id string) (*Item, error) {
	return s.repo.FindByID(id)
}

// FindByIDs will call the repo's function to
// find all weapon items
func (s *itemService) FindByIDs(arr []psql.NullInt) ([]Item, error) {
	return s.repo.FindByIDs(arr)
}

// FindWeapons will call the repo's function to
// find all weapon items
func (s *itemService) FindWeapons(q string) ([]Item, error) {
	return s.repo.FindWeapons(q)
}

// FindArmor will call the repo's function to
// find all armor items
func (s *itemService) FindArmor(q string) ([]Item, error) {
	return s.repo.FindArmor(q)
}

// FindItems will call the repo's function to
// find all armor items
func (s *itemService) FindItems(q string) ([]Item, error) {
	return s.repo.FindItems(q)
}

func (s *itemService) InsertItem(i Item) (int, error) {
	return s.repo.InsertItem(i)
}
