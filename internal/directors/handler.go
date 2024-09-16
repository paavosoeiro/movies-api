package directors

import (
	"encoding/json"
	"net/http"
)

type DirectorHandler struct {
	service Service
}

func New(service Service) *DirectorHandler {
	return &DirectorHandler{
		service: service,
	}
}

func NewDirectorHandlerFactory() *DirectorHandler {
	repo := NewMemoryRepository()
	svc := NewDirectorService(repo)
	return New(svc)
}

func (d DirectorHandler) GetAllDirectors(w http.ResponseWriter, _ *http.Request) {
	directors, err := d.service.GetAll()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	marshal, err := json.Marshal(directors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(marshal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
