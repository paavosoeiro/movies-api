package movies

import (
	"github.com/stretchr/testify/mock"
)

type MockMovieService struct {
	mock.Mock
}

func (m *MockMovieService) GetAllMovies() ([]Movie, error) {
	args := m.Called()
	return args.Get(0).([]Movie), args.Error(1)
}

func (m *MockMovieService) GetMovieById(id string) (*Movie, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*Movie), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockMovieService) CreateMovie(*Movie) (*Movie, error) {
	return nil, nil
}
