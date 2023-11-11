package main

import (
	"InrixBackend/internal/handlers"
	"github.com/go-chi/chi"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/auth", func(mux chi.Router) {
		mux.Get("/", handlers.Repo.Home)
	})

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/shorten", handlers.Repo.ShortenURL)
	mux.Get("/{shortKey:[a-zA-Z0-9]+}", handlers.Repo.Redirect)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
