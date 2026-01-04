package handler

import (
	"mastery-project/internal/server"
	"mastery-project/internal/service"
)

type Handlers struct {
	Health *HealthHandler
	Auth   *AuthHandler
	Item   *ItemHandler
}

func NewHandlers(srv *server.Server, service *service.Services) *Handlers {
	return &Handlers{
		Health: NewHealthHandler(srv),
		Auth:   NewAuthHandler(srv, service.Auth),
		Item:   NewItemHandler(srv, service.Item),
	}
}
