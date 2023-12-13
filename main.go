package main

import (
	"fmt"
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
	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}

	db := con.Conn()

	customer := Customer{
		Id:          "2",
		Name:        "Rizqi",
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
