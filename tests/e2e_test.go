package tests

import (
	"github.com/gorilla/mux"
	"github.com/paavosoeiro/go-movies/pkg/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *mux.Router {
	return router.New()
}

func TestGetAllMovies_E2E(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/movies", nil)

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
