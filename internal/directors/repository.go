package directors

type Repository interface {
	GetAll() []Director
	GetByID(id string) (*Director, error)
	Create(director *Director) (*Director, error)
}
