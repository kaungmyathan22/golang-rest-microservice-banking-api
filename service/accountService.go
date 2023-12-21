package service

import (
	"time"

	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/domain"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/dto"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *exception.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *exception.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	account := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}
	return newAccount.ToNewAccountResponseDto(), nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
