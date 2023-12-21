package domain

import (
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/dto"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"
)

type AccountRepository interface {
	Save(Account) (*Account, *exception.AppError)
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{a.AccountId}
}
