package main

import (
	"InrixBackend/internal/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/auth", func(mux chi.Router) {
		mux.Post("/login", handlers.Repo.LoginUser)
		mux.Post("/register", handlers.Repo.RegisterUser)
		mux.Get("/accessTokenLogin/{access_token}", handlers.Repo.AccessTokenLogin)
	})

	//fileServer := http.FileServer(http.Dir("./static/"))
	//mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
