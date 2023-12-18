package service

import (
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/domain"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *exception.AppError)
	GetCustomer(string) (*domain.Customer, *exception.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *exception.AppError) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *exception.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repository,
	}
}
