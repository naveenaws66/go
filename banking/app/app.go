package app

import (
	"log"
	"net/http"
)

func Start() {
	// routes
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/users", getAllCusomers)

	//router

	log.Fatal(http.ListenAndServe(":8080", mux))
}
