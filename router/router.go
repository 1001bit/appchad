package router

import (
	"github.com/McCooll75/appchad/api/auth"
	"github.com/McCooll75/appchad/api/chatchad"
	"github.com/McCooll75/appchad/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// url handler
// func Router(w http.ResponseWriter, r *http.Request) {
// 	splitUrl := append(strings.Split(r.URL.Path, "/")[1:], "")

// 	// if not logged
// 	if !isLogged(w, r) {
// 		if splitUrl[0] == "api" && splitUrl[1] == "auth" {
// 			switch splitUrl[2] {
// 			case "login":
// 				auth.Login(w, r)
// 			case "register":
// 				auth.Register(w, r)
// 			}
// 			return
// 		}

// 		handlers.Auth(w, r)
// 		return
// 	}

// 	// if logged
// 	switch splitUrl[0] {
// 	case "api":
// 		switch splitUrl[1] {
// 		case "chatchad":
// 			switch r.Method {
// 			case http.MethodGet:
// 				chatchad.ChatGet(w, r)
// 			case http.MethodPost:
// 				chatchad.ChatPost(w, r)
// 			}
// 		}
// 	case "logout":
// 		handlers.Logout(w, r)
// 	case "", "home":
// 		handlers.Home(w, r)
// 	case "chatchad":
// 		handlers.Chatchad(w, r)
// 	case "blogchad":
// 		switch splitUrl[1] {
// 		case "":
// 			handlers.Blogchad(w, r)
// 		case "write":
// 			handlers.BlogchadWrite(w, r)
// 		case "article":
// 			if splitUrl[2] == "" {
// 				http.Error(w, "incorrect id", http.StatusBadRequest)
// 				return
// 			}
// 			id, err := strconv.Atoi(splitUrl[2])
// 			if err != nil {
// 				log.Println("error converting to int:", err)
// 				http.Error(w, "incorrect id", http.StatusBadRequest)
// 				return
// 			}
// 			handlers.BlogchadArticle(w, r, id)
// 		}
// 	default:
// 		http.Error(w, "404 not found!!", http.StatusNotFound)
// 	}
// }

func RouterSetup() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Group(func(r chi.Router) {
		r.Use(guestMiddleware)

		// api
		r.Post("/api/auth/login", auth.Login)
		r.Post("/api/auth/register", auth.Register)
		// pages
		r.Get("/auth", handlers.Auth)
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
