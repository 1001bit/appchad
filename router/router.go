package router

import (
	"github.com/McCooll75/appchad/api/auth"
	"github.com/McCooll75/appchad/api/chatchad"
	"github.com/McCooll75/appchad/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Set up and return router
func RouterSetup() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Group(func(r chi.Router) {
		r.Use(guestMiddleware)

		// api
		r.Post("/api/auth/login", auth.Login)
		r.Post("/api/auth/register", auth.Register)
		// pages
		r.Get("/{any}", handlers.Auth)
	})

	router.Group(func(r chi.Router) {
		r.Use(wideMiddleware)

		// api
		r.Post("/api/chatchad", chatchad.ChatPost)
		r.Get("/api/chatchad", chatchad.ChatGet)
		// pages
		r.Get("/logout", handlers.Logout)
		r.Get("/", handlers.Home)
		r.Get("/home", handlers.Home)
		r.Get("/chatchad", handlers.Chatchad)
		r.Get("/blogchad", handlers.Blogchad)
		r.Get("/blogchad/write", handlers.BlogchadWrite)
		r.Get("/blogchad/article/{id}", handlers.BlogchadArticle)
	})

	return router
}
