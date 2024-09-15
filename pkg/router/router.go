package router

import (
	"github.com/gorilla/mux"
	"github.com/paavosoeiro/go-movies/internal/controller/movies"
	"github.com/paavosoeiro/go-movies/internal/movies/repository"
	"github.com/paavosoeiro/go-movies/internal/movies/service"
	"github.com/paavosoeiro/go-movies/pkg/middleware"
)

func New() *mux.Router {
	repo := repository.NewMemoryRepository()
	svc := service.NewMovieService(repo)
	movieHandler := movies.NewMovieHandler(svc)

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	movieHandler.InitializeRoutes(r)

	return r
}
