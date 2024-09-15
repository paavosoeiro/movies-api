package service

import "github.com/paavosoeiro/go-movies/internal/movies"

type MovieService interface {
	GetAllMovies() ([]movies.Movie, error)
	GetMovieById(id string) (*movies.Movie, error)
	CreateMovie(*movies.Movie) (*movies.Movie, error)
}
