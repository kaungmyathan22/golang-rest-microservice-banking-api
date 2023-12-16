package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "kaung:kaung@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{
		client: client,
	}
}
func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	sqlStatement := "SELECT customer_id,name,city, zipcode, date_of_birth,status from customers"
	rows, err := d.client.Query(sqlStatement)
	if err != nil {
		log.Println("Error wile querying customer table.", err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error wile querying customer table.", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
