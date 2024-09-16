package directors

type Service interface {
	GetAll() ([]Director, error)
}
