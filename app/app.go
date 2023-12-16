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
	mux.HandleFunc("/customers", ch.getAllCustomers)
	log.Print("Server is running....")
	http.ListenAndServe("localhost:8000", mux)
}

//fa5cef883317544e66e8
// client secret: 22cb1b4ec26f3122ba43c3466e585b29224d0594
// pat: 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
