package app

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipCode" xml:"zipCode"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Kaung Myat", City: "Yangon", Zipcode: "11212"},
	}
	w.Header().Add("Content-Type", "application/xml")
	//w.Header().Add("Content-Type", "application/json")
	xml.NewEncoder(w).Encode(customers)
}
