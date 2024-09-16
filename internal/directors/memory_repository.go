package directors

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
)

type MemoryRepository struct {
	mu        sync.RWMutex
	directors map[string]Director
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		directors: make(map[string]Director),
	}
}

func (m *MemoryRepository) GetAll() []Director {
	m.mu.Lock()
	defer m.mu.Unlock()

	var directors []Director
	for _, director := range m.directors {
		directors = append(directors, director)
	}

	return directors
}

func (m *MemoryRepository) GetByID(id string) (*Director, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	director, exists := m.directors[id]
	if !exists {
		return nil, errors.New("director not found")
	}

	return &director, nil
}
func (m *MemoryRepository) Create(director *Director) (*Director, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	id := strconv.Itoa(rand.Intn(1000000))
	director.ID = id

	m.directors[id] = *director

	return director, nil

}
