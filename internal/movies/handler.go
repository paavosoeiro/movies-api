package movies

import (
	"encoding/json"
	"github.com/gorilla/mux"
	e "github.com/paavosoeiro/go-movies/internal/common/err"
	"net/http"
)

type MovieHandler struct {
	service MovieService
}

func NewMovieHandler(service MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func NewMovieHandlerFactory() *MovieHandler {
	repo := NewMemoryRepository()
	svc := NewMovieService(repo)
	return NewMovieHandler(svc)
}

func (m *MovieHandler) List(w http.ResponseWriter, _ *http.Request) {
	movies, err := m.service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := m.service.GetMovieById(params["id"])

	if err != nil {
		errorResponse := e.New(http.StatusNotFound, "Movie not Found", "")
		e.SendErrorResponse(w, errorResponse)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	newMovie, err := m.service.CreateMovie(&movie)
	if err != nil {
		errorResponse := e.New(http.StatusInternalServerError, "An error has occurred", "")
		e.SendErrorResponse(w, errorResponse)
		return
	}

	err = json.NewEncoder(w).Encode(newMovie)
	if err != nil {
		return
	}
}
