package router

import (
	"net/http"
	"time"

	"mastery-project/internal/handler"
	authMiddleware "mastery-project/internal/middleware"

	"github.com/go-chi/httprate"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(
	h *handler.Handlers,
	authMW *authMiddleware.AuthMiddleware,
) chi.Router {

	r := chi.NewRouter()

	//Global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//rate limit
	r.Use(httprate.Limit(
		10,
		time.Minute,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, `{"error": "Rate-limited. Please, slow down."}`, http.StatusTooManyRequests)
		}),
	))

	//Static uploads (public)
	r.Handle(
		"/uploads/*",
		http.StripPrefix(
			"/uploads/",
			http.FileServer(http.Dir("uploads")),
		),
	)

	//API v1
	r.Route("/api/v1", func(r chi.Router) {
		registerSystemRoutes(r, h)
		//Public auth routes
		r.Route("/auth", func(r chi.Router) {
			registerAuthRoutes(r, h)
		})

		//Protected routes
		r.With(authMW.Protected).Group(func(r chi.Router) {
			registerItemRoutes(r, h)
		})
	})

	return r
}
