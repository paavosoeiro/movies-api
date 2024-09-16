package router

import (
	"github.com/gorilla/mux"
	"github.com/paavosoeiro/go-movies/internal/directors"
	"github.com/paavosoeiro/go-movies/internal/movies"
	"github.com/paavosoeiro/go-movies/pkg/middleware"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type RouteGroup struct {
	Prefix string
	Routes []Route
}

func New() *mux.Router {
	movieHandler := movies.NewMovieHandlerFactory()
	directorHandler := directors.NewDirectorHandlerFactory()

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	routeGroups := []RouteGroup{
		{
			Prefix: "/movies",
			Routes: []Route{
				{Path: "", Method: "GET", Handler: movieHandler.List},
				{Path: "/{id}", Method: "GET", Handler: movieHandler.GetMovieById},
				{Path: "/{id}", Method: "POST", Handler: movieHandler.CreateMovie},
			},
		},
		{
			Prefix: "/directors",
			Routes: []Route{
				{Path: "", Method: "GET", Handler: directorHandler.GetAllDirectors},
			},
		},
	}

	for _, group := range routeGroups {
		registerRoutes(r, group)
	}

	//movieHandler.InitializeRoutes(r)

	return r
}

func registerRoutes(router *mux.Router, group RouteGroup) {
	for _, route := range group.Routes {
		router.HandleFunc(group.Prefix+route.Path, route.Handler).Methods(route.Method)
	}
}
