package movies

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
)

type MemoryRepository struct {
	mu     sync.RWMutex
	movies map[string]Movie
}

func NewMemoryRepository() *MemoryRepository {
	instance := &MemoryRepository{
		movies: make(map[string]Movie),
	}

	instance.movies["1"] = Movie{ID: "1", Isbn: "Isbn", Title: "O Senhor dos Aneis",
		Director: &Director{Firstname: "Peter", Lastname: "Jackson"}}

	return instance
}

func (r *MemoryRepository) GetAll() ([]Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var movies []Movie
	for _, movie := range r.movies {
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *MemoryRepository) GetById(id string) (*Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	movie, exists := r.movies[id]
	if !exists {
		return nil, errors.New("movie not found")
	}
	return &movie, nil
}

func (r *MemoryRepository) Create(movie *Movie) (*Movie, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := strconv.Itoa(rand.Intn(1000000))
	movie.ID = id

	r.movies[id] = *movie

	return movie, nil

}
