package router

import (
	"github.com/gorilla/mux"
	"github.com/paavosoeiro/go-movies/internal/controller/movies"
)

func New(movieHandler *movies.MovieHandler) *mux.Router {
	r := mux.NewRouter()

	movieHandler.InitializeRoutes(r)

	return r
}
