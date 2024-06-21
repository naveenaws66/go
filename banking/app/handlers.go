package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcity"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCusomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "person1", City: "city1", Zipcode: "1"},
		{Name: "person2", City: "city2", Zipcode: "2"},
	}

	if w.Header().Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
