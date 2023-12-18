package usecase

import (
	"fmt"
	"laundry/model"
	"laundry/repository"
)

type ProductUseCase interface {
	CreateNew(payload model.Product) error
	FindById(id string) (model.Product, error)
	GetByName(name string) ([]model.Product, error)
	FindAll() ([]model.Product, error)
	Update(payload model.Product) error
	Delete(id string) error
}

type productUserCase struct {
	repo  repository.ProductRepository
	uomUC UomUseCase
}

func (p *productUserCase) CreateNew(payload model.Product) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price <= 0 {
		return fmt.Errorf("price must be greater than zero")
	}

	_, err := p.uomUC.FindById(payload.Uom.Id)
	if err != nil {
		return err // karena sudah di validasi di FindById uomUC
	}

	err = p.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new product: %v", err)
	}

	return nil
}

func (p *productUserCase) FindById(id string) (model.Product, error) {
	product, err := p.repo.FindById(id)
	if err != nil {
		return model.Product{}, fmt.Errorf("product not found")
	}
	return product, nil
}

func (p *productUserCase) GetByName(name string) ([]model.Product, error) {
	//byName, err := p.repo.FindByName(name)
	//if err != nil {
	//	return nil, fmt.Errorf("error while fetching products by name: %v", err)
	//}
	//
	//if len(byName) == 0 {
	//	return nil, fmt.Errorf("no products found with the name: %s", name)
	//}
	//
	//return byName, nil
	return p.repo.FindByName(name)
}

func (p *productUserCase) FindAll() ([]model.Product, error) {
	//products, err := p.repo.FindAll()
	//if err != nil {
	//	return nil, fmt.Errorf("data product is empty: %v", err)
	//}
	//
	//return products, nil
	return p.repo.FindAll()
}

func (p *productUserCase) Update(payload model.Product) error {
	_, err := p.repo.FindById(payload.Id)
	if err != nil {
		return fmt.Errorf("id not found: %v", err)
	}

	err = p.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}
	return nil
}

func (p *productUserCase) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}

	err = p.repo.DeleteById(product.Id)
	if err != nil {
		return fmt.Errorf("Failed to delete product: %v", err)
	}
	return nil
}

// Constructor
func NewProductUseCase(repo repository.ProductRepository, uomUc UomUseCase) ProductUseCase {
	return &productUserCase{
		repo:  repo,
		uomUC: uomUc,
	}
}
