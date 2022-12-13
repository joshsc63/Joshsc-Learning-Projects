package main

import (
	"fmt"
	"net/http"
)

// const cant be changed by application
const portNumber = ":8080"

// Handler Functions
// uppercase becomes public/visible outside of package
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")

}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))

}

// starting lowercase makes it private. Can only be called in this package
// adds two ints & returns sum
func addValues(x, y int) int {
	//var sum int
	//sum = x + y
	//return sum
	return x + y
}

func main() {

	// HELLO WORLD
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "hello world!")
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	//})

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// start web server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
