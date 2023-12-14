package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/customers", getAllCustomers)
	http.ListenAndServe("localhost:8000", nil)
}
func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Kaung Myat", City: "Yangon", Zipcode: "11212"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
