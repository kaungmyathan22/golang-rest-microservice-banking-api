package domain

import "github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *exception.AppError)
	ById(string) (*Customer, *exception.AppError)
}
