package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepository())}

	router.HandleFunc("/customers", ch.getAllConsumner).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
