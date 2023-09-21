package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/checkSiteStatus", checkSiteStatusHandler)
	http.HandleFunc("/ping", pingHandler)

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func checkSiteStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "url" query parameter from the request
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	// Send an HTTP HEAD request to the site to check its status
	resp, err := http.Head(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the status code and status text as the response
	fmt.Fprintf(w, "Status code: %d, Status text: %s", resp.StatusCode, resp.Status)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Return a simple "pong" response
	fmt.Fprintf(w, "Pong! Timestamp: %s", time.Now().Format("2006-01-02 15:04:05"))
}
