package router

import (
	"mastery-project/internal/handler"

	"github.com/go-chi/chi/v5"
)

func registerSystemRoutes(r chi.Router, h *handler.Handlers) {
	r.Get("/health", h.Health.CheckHealth)
}

func registerAuthRoutes(r chi.Router, h *handler.Handlers) {
	r.Post("/login", h.Auth.Login)
	r.Post("/register", h.Auth.Signup)
	r.Get("/logout", h.Auth.Logout)
}

func registerItemRoutes(r chi.Router, h *handler.Handlers) {
	r.Route("/items", func(r chi.Router) {
		r.Get("/", h.Item.GetAll)
		r.Post("/", h.Item.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.Item.GetOne)
			r.Patch("/", h.Item.Update)
			r.Delete("/", h.Item.Delete)
		})
	})
}
