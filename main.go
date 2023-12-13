package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"laundry/config"
)

// Struct
type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	Address     string
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbConfig.Host,
		cfg.DbConfig.Port,
		cfg.DbConfig.User,
		cfg.DbConfig.Password,
		cfg.DbConfig.Name,
	)
	db, err := sql.Open(cfg.DbConfig.Driver, dsn)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db)

	customer := Customer{
		Id:          "1",
		Name:        "Reza",
		PhoneNumber: "0857756",
		Address:     "Jl. Raya Serang Km.24",
	}

	// Exec entry query DML
	_, err = db.Exec("INSERT INTO m_customer VALUES ($1, $2, $3, $4)",
		customer.Id,
		customer.Name,
		customer.PhoneNumber,
		customer.Address,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success inserting data")

}
