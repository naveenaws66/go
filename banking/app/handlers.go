package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/naveenaws66/banking/service"
)

type Customer struct {
	Name    string `json:"full name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcity"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCusomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if w.Header().Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
