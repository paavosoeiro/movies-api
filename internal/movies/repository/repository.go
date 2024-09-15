package repository

import "github.com/paavosoeiro/go-movies/internal/movies"

type Repository interface {
	GetAll() ([]movies.Movie, error)
	GetById(id string) (*movies.Movie, error)
	Create(movie *movies.Movie) (*movies.Movie, error)
}
