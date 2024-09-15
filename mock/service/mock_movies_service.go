package service

import (
	"github.com/paavosoeiro/go-movies/internal/movies"
	"github.com/stretchr/testify/mock"
)

type MockMovieService struct {
	mock.Mock
}

func (m *MockMovieService) GetAllMovies() ([]movies.Movie, error) {
	args := m.Called()
	return args.Get(0).([]movies.Movie), args.Error(1)
}

func (m *MockMovieService) GetMovieById(id string) (*movies.Movie, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*movies.Movie), args.Error(1)
	}

	return nil, args.Error(1)
}
