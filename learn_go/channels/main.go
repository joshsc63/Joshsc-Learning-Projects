package main

import (
	"fmt"
	"net/http"
)

// Check status of websites with channels
// Use go routine to check each URL in parallel
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// ignore index w/ _
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // 2 return vals... resp & error
	if err != nil {
		fmt.Println(link, "might be down!")
		return // exit function
	}

	fmt.Println(link, "is up!")
}
