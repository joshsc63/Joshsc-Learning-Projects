package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter() //mux common name of routing in go

	//specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept","Authorization","Content-Type","X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
	
	mux.Use(middleware.Heartbeat("/ping")) //Easy way to check its available
	
	
	return mux
}