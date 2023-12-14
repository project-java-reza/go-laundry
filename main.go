package main

import (
	"fmt"
	"laundry/config"
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

	// uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Create
	//uomRepo.Save(model.Uom{
	//	Id:          "2",
	//	Description: "gr",
	//})

	// Get By Id
	//uom, err := uomRepo.FindById("1")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(uom)

	// UPDATE CARA 1
	//uomUpdate, err := uomRepo.FindById("1")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//uomUpdate.Description = "Gr"
	//
	//if err = uomRepo.Update(uomUpdate); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	// UPDATE CARA 2
	//err = uomRepo.Update(model.Uom{
	//	Id:          "2",
	//	Description: "new_gr",
	//})
	//
	//if err != nil {
	//	fmt.Println("Error updating UOM: ", err)
	//	return
	//}
	//fmt.Println("UOM updated successfully.")

	// Get All
	//uoms, err := uomRepo.FindAll()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	for _, uom := range uoms {
	//		fmt.Printf("ID: %s, Name: %s\n", uom.Id, uom.Description)
	//	}
	//}

	// Delete By Id
	//err = uomRepo.DeleteById("1")
	//if err != nil {
	//	fmt.Println("Error Deleting:", err)
	//} else {
	//	fmt.Println("Delete Successfully!")
	//}

	// Create Product
	//productRepo.Save(model.Product{
	//	Id:    "1",
	//	Name:  "Cuci",
	//	Price: 10000,
	//	Uom: model.Uom{
	//		Id: "1",
	//	},
	//})

	// Get By Id
	//product, err := productRepo.FindById("1")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(product)

	// Get Product By Name
	//products, err := productRepo.FindByName("Cuci")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(products)

	// GetAll Product
	//products, err := productRepo.FindAll()
	//for _, product := range products {
	//	fmt.Printf("id: %s, Name: %s, Price: %s, Satuan: %s",
	//		product.Id,
	//		product.Name,
	//		product.Price,
	//		product.Uom.Description,
	//	)
	//}

	//if err != nil {
	//	fmt.Println(err)
	//}

	// Update Product
	//err = productRepo.Update(model.Product{
	//	Id:    "1",
	//	Name:  "New Cuci",
	//	Price: 15000,
	//	Uom: model.Uom{
	//		Id: "2",
	//	},
	//})
	//
	//if err != nil {
	//	fmt.Println("Error updating Product: ", err)
	//	return
	//}
	//fmt.Println("Product updated successfully.")

	// Delete Product By Id
	err = productRepo.DeleteById("1")
	if err != nil {
		fmt.Println("Error Deleting:", err)
	} else {
		fmt.Println("Delete Successfully!")
	}

}
