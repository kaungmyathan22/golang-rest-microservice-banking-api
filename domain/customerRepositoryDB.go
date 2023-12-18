package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kaungmyathan22/golang-rest-microservice-banking-api/exception"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "kaung:kaung@/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
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
	rows, err := d.client.Query(sqlStatement)
	if err != nil {
		log.Println("Error wile querying customer table.", err.Error())
		return nil, exception.HttpInternalServerError("something went wrong.")
	}

	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {
		log.Println("Error wile querying customer table.", err.Error())
		return nil, exception.HttpInternalServerError("something went wrong.")
	}
	return customers, nil
}
func (d CustomerRepositoryDb) ById(id string) (*Customer, *exception.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.HttpNotFoundException("customer not found.")
		}
		log.Println("Error wile querying customer table.", err.Error())
		return nil, exception.HttpInternalServerError("something went wrong.")
	}
	return &c, nil
}
