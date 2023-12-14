package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greetHandler)
	mux.HandleFunc("/customers", getAllCustomers)
	log.Print("Server is running....")
	http.ListenAndServe("localhost:8000", mux)
}
