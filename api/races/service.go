package races

// RaceService defines the outward functionality
// available to interact with Race data
type RaceService interface {
	List() ([]Race, error)
	FindByID(id string) (*Race, error)
}

type raceService struct {
	repo RaceRepo
}

// NewRaceService returns a new implementation of RaceService
func NewRaceService(repo RaceRepo) RaceService {
	return &raceService{
		repo: repo,
	}
}

// FindById will call the repo's function to
// find a Race by ID
func (s *raceService) FindByID(id string) (race *Race, err error) {
	race, err = s.repo.FindByID(id)
	return race, err
}

// FindById will call the repo's function to
// find a Race by ID
func (s *raceService) List() (races []Race, err error) {
	return s.repo.List()
}
