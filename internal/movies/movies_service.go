package movies

type MovieServiceImpl struct {
	repo Repository
}

func NewMovieService(repo Repository) *MovieServiceImpl {
	return &MovieServiceImpl{
		repo: repo,
	}
}

func (m *MovieServiceImpl) GetAllMovies() ([]Movie, error) {
	return m.repo.GetAll()
}

func (m *MovieServiceImpl) GetMovieById(id string) (*Movie, error) {
	return m.repo.GetById(id)
}

func (m *MovieServiceImpl) CreateMovie(movie *Movie) (*Movie, error) {
	return m.repo.Create(movie)
}
