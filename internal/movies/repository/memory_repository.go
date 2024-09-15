package repository

import (
	"errors"
	movies2 "github.com/paavosoeiro/go-movies/internal/movies"
	"math/rand"
	"strconv"
	"sync"
)

type MemoryRepository struct {
	mu     sync.RWMutex
	movies map[string]movies2.Movie
}

func NewMemoryRepository() *MemoryRepository {
	instance := &MemoryRepository{
		movies: make(map[string]movies2.Movie),
	}

	instance.movies["1"] = movies2.Movie{ID: "1", Isbn: "Isbn", Title: "O Senhor dos Aneis",
		Director: &movies2.Director{Firstname: "Peter", Lastname: "Jackson"}}

	return instance
}

func (r *MemoryRepository) GetAll() ([]movies2.Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var movies []movies2.Movie
	for _, movie := range r.movies {
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *MemoryRepository) GetById(id string) (*movies2.Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	movie, exists := r.movies[id]
	if !exists {
		return nil, errors.New("movie not found")
	}
	return &movie, nil
}

func (r *MemoryRepository) Create(movie *movies2.Movie) (*movies2.Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := strconv.Itoa(rand.Intn(1000000))
	movie.ID = id

	r.movies[id] = *movie

	return movie, nil

}
