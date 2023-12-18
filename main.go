package main

import (
	"fmt"
	"laundry/config"
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
	productRepo := repository.NewProductRepository(db)

	uomUC := usecase.NewUomUseCase(uomRepo)
	productUC := usecase.NewProductUseCase(productRepo, uomUC)

	// Create Uom
	//uomUC.CreateNew(model.Uom{
	//	Id:          "4",
	//	Description: "Buah",
	//})

	// Get By Id UOM
	//uom, err := uomUC.FindById("3")
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//} else {
	//	fmt.Printf("id: %s, Description: %s", uom.Id, uom.Description)
	//}

	// Get All UOM
	//uoms, err := uomUC.FindAll()
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//} else {
	//	for _, uom := range uoms {
	//		fmt.Printf("id: %s, description: %s\n",
	//			uom.Id,
	//			uom.Description,
	//		)
	//	}
	//}

	//// Update UOM
	//err = uomUC.Update(model.Uom{
	//	Id:          "1",
	//	Description: "Kg",
	//})
	//
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//}

	//Create Product
	//err = productUC.CreateNew(model.Product{
	//	Id:    "1",
	//	Name:  "Cuci Baju",
	//	Price: 20000,
	//	Uom: model.Uom{
	//		Id: "4",
	//	},
	//})
	//
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	return
	//}

	// FindById Product
	//product, err := productUC.FindById("5")
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//} else {
	//	fmt.Println(product)
	//}

	// FindAll Product
	products, err := productUC.FindAll()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		for _, product := range products {
			fmt.Printf("id: %s, name: %s, price: %s, uom_id: %s\n",
				product.Id,
				product.Name,
				product.Price,
				product.Uom,
			)
		}
	}

	// GetByName Product
	//products, err := productUC.GetByName("Cuci Boneka")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	//fmt.Println("Products found:", products)

	// Update Product
	//updatePayload := model.Product{
	//	Id:    "10",
	//	Name:  "Cuci Boneka new",
	//	Price: 15000,
	//	Uom: model.Uom{
	//		Id: "3",
	//	},
	//}

	//err = productUC.Update(updatePayload)
	//if err != nil {
	//	fmt.Println("Error : ", err)
	//}

	// Delete Product
	//err = productUC.Delete("3")
	//if err != nil {
	//	fmt.Printf("Error deleting product: %v\n", err)
	//	return
	//}
}
