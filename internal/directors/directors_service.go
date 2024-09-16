package directors

type DirectorService struct {
	repo Repository
}

func NewDirectorService(repo Repository) *DirectorService {
	return &DirectorService{
		repo: repo,
	}
}

func (d *DirectorService) GetAll() ([]Director, error) {
	return d.repo.GetAll(), nil
}
