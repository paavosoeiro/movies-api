package movies

type MovieService interface {
	GetAllMovies() ([]Movie, error)
	GetMovieById(id string) (*Movie, error)
	CreateMovie(*Movie) (*Movie, error)
}
