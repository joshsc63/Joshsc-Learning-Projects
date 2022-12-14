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

	// main channel
	c := make(chan string) //make channel of type string

	// ignore index w/ _
	for _, link := range links {
		// pass channel to func
		go checkLink(link, c) // go routine... use go infront of function calls
	}

	// receive data from channel
	fmt.Println(<-c)
}

// Send data to channel:   channel <- 5    (send the value 5 into channel
// myNumber <- channel : wait for value to be sent into channel.. when returned, assign value to myNumber
// fmt.Println(<- channel) : wait for value to be sent in channel. When returned, log it out

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // 2 return vals... resp & error
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down"
		return // exit function
	}

	fmt.Println(link, "is up!")
	c <- "site is up"
}
