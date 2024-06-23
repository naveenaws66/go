package main

import (
	"fmt"
	"net/http"
)

// main is the entry point of the program.
//
// It sets up a HTTP server to handle requests to the root path ("/"). When a request is received,
// it writes the string "Hello, world" to the response writer and logs a message indicating that the request
// has been served. The server listens on port 8080.
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Hello, world")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Request served")

	})
	_ = http.ListenAndServe(":8080", nil)
}
