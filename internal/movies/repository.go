package movies

type Repository interface {
	GetAll() ([]Movie, error)
	GetById(id string) (*Movie, error)
	Create(movie *Movie) (*Movie, error)
}
