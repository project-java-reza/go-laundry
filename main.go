package main

import (
	"fmt"
	"laundry/config"
	"laundry/model"
	"laundry/repository"
	"laundry/usecase"
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
	uomUC := usecase.NewUomUseCase(uomRepo)
	uomUC.CreateNew(model.Uom{
		Id:          "3",
		Description: "Buah",
	})

}
