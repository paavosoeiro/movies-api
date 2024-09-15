package movies

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/paavosoeiro/go-movies/internal/movies"
	"github.com/paavosoeiro/go-movies/mock/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovieHandler_GetMovieById(t *testing.T) {
	mockService := new(service.MockMovieService)

	movie := &movies.Movie{ID: "1", Isbn: "Isbn", Title: "O Senhor dos Aneis",
		Director: &movies.Director{Firstname: "Peter", Lastname: "Jackson"}}

	mockService.On("GetMovieById", "1").Return(movie, nil)

	handler := NewMovieHandler(mockService)

	req, _ := http.NewRequest("GET", "/movies/1", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", handler.GetMovieById)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var returnedMovie movies.Movie
	err := json.Unmarshal(rr.Body.Bytes(), &returnedMovie)
	if err != nil {
		return
	}
	assert.Equal(t, movie, &returnedMovie)

	mockService.AssertExpectations(t)
}
