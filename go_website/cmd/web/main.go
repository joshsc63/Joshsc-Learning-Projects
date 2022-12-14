package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joshsc63/go-website/pkg/config"
	"github.com/joshsc63/go-website/pkg/handlers"
	"github.com/joshsc63/go-website/pkg/render"
)

// const cant be changed by application
const portNumber = ":8080"

func main() {
	// get the template cache from app config
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// call new template cache
	render.NewTemplates(&app) // takes in pointer, use & ref

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	// start web server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
