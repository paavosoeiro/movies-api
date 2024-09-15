package err

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func New(status int, message string, details string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Details: details,
	}
}

func SendErrorResponse(w http.ResponseWriter, response *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
