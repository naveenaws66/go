package main

import (
	"net/http"
	"fmt"
	"log"
)


func main() {

        http.HandleFunc("/", homepage)

	log.Fatal(http.ListenAndServe(":80", nil))

}

func homepage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprintf(w, path)
}



