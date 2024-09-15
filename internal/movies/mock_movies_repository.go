package movies

import (
	"github.com/stretchr/testify/mock"
)

type MockMoviesRepository struct {
	mock.Mock
}

func (m *MockMoviesRepository) GetAll() ([]Movie, error) {
	args := m.Called()
	return args.Get(0).([]Movie), args.Error(1)
}

func (m *MockMoviesRepository) GetById(id string) (*Movie, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*Movie), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMoviesRepository) Create(movie *Movie) (*Movie, error) {
	return nil, nil
}
