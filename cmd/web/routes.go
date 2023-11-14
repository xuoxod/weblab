package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/xuoxod/weblab/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/", func(mux chi.Router) {
		mux.Get("/", handlers.Home)
		mux.Get("/about", handlers.About)
		mux.Post("/login", handlers.Authenticate)
	})

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
