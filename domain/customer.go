package domain

import "github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}
type CustomerRepository interface {
	FindAll(string) ([]Customer, *exception.AppError)
	ById(string) (*Customer, *exception.AppError)
}
