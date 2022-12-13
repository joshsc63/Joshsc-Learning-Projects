package main

import (
	"fmt"
	"net/http"

	"github.com/joshsc63/go-website/pkg/handlers"
)

// const cant be changed by application
const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// start web server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
