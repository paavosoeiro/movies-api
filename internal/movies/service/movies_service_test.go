package service

import (
	"errors"
	"github.com/paavosoeiro/go-movies/internal/movies"
	"github.com/paavosoeiro/go-movies/internal/movies/repository"
	repository2 "github.com/paavosoeiro/go-movies/mock/repository"
	"reflect"
	"testing"
)

func TestMovieServiceImpl_GetAllMovies(t *testing.T) {
	mockRepo := new(repository2.MockMoviesRepository)

	mockRepo.On("GetAll").Return([]movies.Movie{{ID: "1", Isbn: "Isbn", Title: "O Senhor dos Aneis",
		Director: &movies.Director{Firstname: "Peter", Lastname: "Jackson"}}}, nil)

	type fields struct {
		repo repository.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []movies.Movie
		wantErr bool
	}{
		{
			name:   "return all movies",
			fields: fields{mockRepo},
			want: []movies.Movie{{ID: "1", Isbn: "Isbn", Title: "O Senhor dos Aneis",
				Director: &movies.Director{Firstname: "Peter", Lastname: "Jackson"}}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieServiceImpl{
				repo: tt.fields.repo,
			}
			got, err := m.GetAllMovies()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieServiceImpl_GetMovieById(t *testing.T) {
	mockRepo := new(repository2.MockMoviesRepository)

	movie := &movies.Movie{ID: "1", Isbn: "Isbn", Title: "O Senhor dos Aneis",
		Director: &movies.Director{Firstname: "Peter", Lastname: "Jackson"}}

	mockRepo.On("GetById", "1").Return(&movies.Movie{ID: "1", Isbn: "Isbn", Title: "O Senhor dos Aneis",
		Director: &movies.Director{Firstname: "Peter", Lastname: "Jackson"}}, nil)

	mockRepo.On("GetById", "2").Return(nil, errors.New("movie not found"))

	type fields struct {
		repo repository.Repository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *movies.Movie
		wantErr bool
	}{
		{
			name:    "movie with id 1 will return",
			fields:  fields{mockRepo},
			args:    args{"1"},
			want:    movie,
			wantErr: false,
		},
		{
			name:    "movie with id 2 return error",
			fields:  fields{mockRepo},
			args:    args{"2"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MovieServiceImpl{
				repo: tt.fields.repo,
			}
			got, err := m.GetMovieById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovieById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMovieById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
