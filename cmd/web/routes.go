package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/xuoxod/weblab/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Compress(5))
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)
	// mux.Use(RecoverPanic)
	mux.Use(WriteToConsole)
	mux.Use(middleware.NoCache)
	mux.Use(NoSurf)

	mux.Route("/", func(mux chi.Router) {
		mux.Use(Unauth)
		mux.Get("/", handlers.Repo.Home)
		mux.Get("/about", handlers.Repo.About)
		mux.Post("/login", handlers.Authenticate)
	})

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
