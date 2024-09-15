package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/paavosoeiro/go-movies/internal/controller/movies"
	"github.com/paavosoeiro/go-movies/internal/movies/repository"
	"github.com/paavosoeiro/go-movies/internal/movies/service"
	"github.com/paavosoeiro/go-movies/pkg/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	repo := repository.NewMemoryRepository()

	svc := service.NewMovieService(repo)

	movieHandler := movies.NewMovieHandler(svc)
	r := router.New(movieHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	done := make(chan bool)

	go func() {
		fmt.Printf("Starting server at port :8080\n")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Server failed to start:", err)
		}
	}()

	go func() {
		<-stop

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Println("Shutting down gracefully...")
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}

		close(done)
	}()

	<-done
	log.Println("Server stopped")

}
