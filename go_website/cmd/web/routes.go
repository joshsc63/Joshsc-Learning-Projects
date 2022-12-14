package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joshsc63/go-website/pkg/config"
	"github.com/joshsc63/go-website/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/", handlers.Repo.About)

	return mux
}
