package handler

import (
	"mastery-project/internal/server"
	"net/http"
	"time"
)

type HealthHandler struct {
	Handler
}

func NewHealthHandler(srv *server.Server) *HealthHandler {
	return &HealthHandler{Handler: NewHandler(srv)}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": "healthy",
		"time":   time.Now().UTC(),
		"env":    h.server.Config.ENV,
	}
	//check db here
	h.JSON(w, http.StatusOK, response)
}
