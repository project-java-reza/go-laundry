package main

import (
	"fmt"
	"laundry/config"
	"laundry/model"
	"laundry/repository"
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

	uomRepo := repository.NewUomRepository(db)
	uomRepo.Save(model.Uom{
		Id:          "1",
		Description: "Kg",
	})

	uom, err := uomRepo.FindById("1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uom)
}
