package repository

import (
	"github.com/paavosoeiro/go-movies/internal/movies"
	"github.com/stretchr/testify/mock"
)

type MockMoviesRepository struct {
	mock.Mock
}

func (m *MockMoviesRepository) GetAll() ([]movies.Movie, error) {
	args := m.Called()
	return args.Get(0).([]movies.Movie), args.Error(1)
}

func (m *MockMoviesRepository) GetById(id string) (*movies.Movie, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*movies.Movie), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMoviesRepository) Create(movie *movies.Movie) (*movies.Movie, error) {
	panic("implement me")
}
