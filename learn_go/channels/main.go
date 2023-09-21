package main

import (
	"fmt"
	"net/http"
	"time"
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

	// receive data from channels
	//for i := 0; i < len(links); i++ {
	//fmt.Println(<-c)
	//}

	// infinite loop to keep checking link
	//for {
	//	go checkLink(<-c, c)
	//}

	// l for link

	for l := range c {
		//time.Sleep(5 * time.Second)
		//go checkLink(l, c)

		// function literal - unnamed function
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)

		}(l) // () to call/invoke it

	}

}

// Send data to channel:   channel <- 5    (send the value 5 into channel
// myNumber <- channel : wait for value to be sent into channel.. when returned, assign value to myNumber
// fmt.Println(<- channel) : wait for value to be sent in channel. When returned, log it out

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // 2 return vals... resp & error
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return // exit function
	}

	fmt.Println(link, "is up!")
	c <- link
}
