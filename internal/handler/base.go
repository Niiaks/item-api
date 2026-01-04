package handler

import (
	"encoding/json"
	"mastery-project/internal/server"
	"net/http"
)

type Handler struct {
	server *server.Server
}

// NewHandler creates a new base handler
func NewHandler(s *server.Server) Handler {
	return Handler{server: s}
}

func (h Handler) JSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return
	}
}
