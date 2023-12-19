package domain

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{
		client: client,
	}
}
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *exception.AppError) {
	filterQuery := ""
	if status != "" {
		if status == "active" {
			filterQuery = "where status = 1"
		} else {
			filterQuery = "where status = 0"
		}
	}
	sqlStatement := "SELECT customer_id,name,city, zipcode, date_of_birth,status from customers " + filterQuery
	customers := make([]Customer, 0)
	err := d.client.Select(&customers, sqlStatement)
	if err != nil {
		log.Println("Error wile querying customer table.", err.Error())
		return nil, exception.HttpInternalServerError("something went wrong.")
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *exception.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.HttpNotFoundException("customer not found.")
		}
		log.Println("Error wile querying customer table.", err.Error())
		return nil, exception.HttpInternalServerError("something went wrong.")
	}
	return &c, nil
}
