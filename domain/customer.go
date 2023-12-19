package domain

import "github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"

type Customer struct {
	Id          string `db:"customer_id" json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipCode"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	Status      string `json:"status"`
}
type CustomerRepository interface {
	FindAll(string) ([]Customer, *exception.AppError)
	ById(string) (*Customer, *exception.AppError)
}
