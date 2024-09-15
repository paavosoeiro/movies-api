package service

import (
	"github.com/paavosoeiro/go-movies/internal/movies"
	"github.com/paavosoeiro/go-movies/internal/movies/repository"
)

type MovieServiceImpl struct {
	repo repository.Repository
}

func NewMovieService(repo repository.Repository) *MovieServiceImpl {
	return &MovieServiceImpl{
		repo: repo,
	}
}

func (m *MovieServiceImpl) GetAllMovies() ([]movies.Movie, error) {
	return m.repo.GetAll()
}

func (m *MovieServiceImpl) GetMovieById(id string) (*movies.Movie, error) {
	return m.repo.GetById(id)
}
