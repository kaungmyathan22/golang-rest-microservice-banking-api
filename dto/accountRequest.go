package dto

import (
	"strings"

	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *exception.AppError {
	if r.Amount < 5000 {
		return exception.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return exception.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
