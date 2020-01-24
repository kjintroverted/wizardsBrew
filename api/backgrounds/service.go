package backgrounds

// BackgroundService defines the outward functionality
// available to interact with Background data
type BackgroundService interface {
	List() ([]Background, error)
	FindByID(id string) (*Background, error)
}

type backgroundService struct {
	repo BackgroundRepo
}

// NewBackgroundService returns a new implementation of BackgroundService
func NewBackgroundService(repo BackgroundRepo) BackgroundService {
	return &backgroundService{
		repo: repo,
	}
}

// FindById will call the repo's function to
// find a Background by ID
func (s *backgroundService) FindByID(id string) (background *Background, err error) {
	background, err = s.repo.FindByID(id)
	return background, err
}

// FindById will call the repo's function to
// find a Background by ID
func (s *backgroundService) List() (backgrounds []Background, err error) {
	return s.repo.List()
}
