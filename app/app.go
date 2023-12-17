package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/domain"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/service"
)

func Start() {
	mux := mux.NewRouter()
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Print("Server is running....")
	http.ListenAndServe("localhost:8000", mux)
}
