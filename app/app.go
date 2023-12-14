package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()
	mux.HandleFunc("/greet", greetHandler)
	mux.HandleFunc("/customers", getAllCustomers)
	mux.HandleFunc("/customers/{customer_id}", getCustomer)
	log.Print("Server is running....")
	http.ListenAndServe("localhost:8000", mux)
}
