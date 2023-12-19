package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (db AccountRepositoryDb) Save(account Account) (*Account, *exception.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := db.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, exception.HttpInternalServerError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, exception.HttpInternalServerError("Unexpected error from database")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil

}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
