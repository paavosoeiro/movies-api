package movies

import (
	"encoding/json"
	"github.com/gorilla/mux"
	e "github.com/paavosoeiro/go-movies/internal/common/err"
	movies2 "github.com/paavosoeiro/go-movies/internal/movies"
	"github.com/paavosoeiro/go-movies/internal/movies/service"
	"net/http"
)

type MovieHandler struct {
	service service.MovieService
}

func NewMovieHandler(service service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (m *MovieHandler) InitializeRoutes(r *mux.Router) {
	r.HandleFunc("/movies", m.List).Methods("GET")
	r.HandleFunc("/movies/{id}", m.GetMovieById).Methods("GET")
	r.HandleFunc("/movies", m.CreateMovie).Methods("POST")
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
	var movie movies2.Movie
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
