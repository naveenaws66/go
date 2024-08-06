package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/naveenaws66/banking/domain"
	"github.com/naveenaws66/banking/service"
)

func Start() {
	// routes
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//define routes
	router.HandleFunc("/customers", ch.getAllCusomers).Methods(http.MethodGet)
	//router

	log.Fatal(http.ListenAndServe(":8080", router))
}
