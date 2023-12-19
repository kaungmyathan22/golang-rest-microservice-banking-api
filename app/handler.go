package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/service"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipCode" xml:"zipCode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	statusQuery := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(statusQuery)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["id"]
	customer, err := ch.service.GetCustomer(customer_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)

}
