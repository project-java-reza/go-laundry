package usecase

import (
	"fmt"
	"laundry/model"
	"laundry/repository"
)

type UomUseCase interface {
	CreateNew(payload model.Uom) error
	FindById(id string) (model.Uom, error)
	FindAll() ([]model.Uom, error)
	Update(payload model.Uom) error
	Delete(id string) error
}

type uomUseCase struct {
	repo repository.UomRepository
}

// Create Product
func (u *uomUseCase) CreateNew(payload model.Uom) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Description == "" {
		return fmt.Errorf("name is required")
	}

	err := u.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new uom: %v", err) // v ini mendukung semua tipe data termasuk pointer
	}

	return nil
}

// Get By Id
func (u *uomUseCase) FindById(id string) (model.Uom, error) {
	uom, err := u.repo.FindById(id)
	if err != nil {
		return model.Uom{}, fmt.Errorf("uom not found")
	}
	return uom, nil
}

// Get All
func (u *uomUseCase) FindAll() ([]model.Uom, error) {
	uoms, err := u.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all uom: %v", err)
	}
	return uoms, nil
}

// Update
func (u *uomUseCase) Update(payload model.Uom) error {
	_, err := u.FindById(payload.Id)
	if err != nil {
		return err // hanya return err karena sama dengan validasi err yang ada di find by id
	}

	err = u.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom: %v", err)
	}

	return nil
}

// Delete By Id
func (u *uomUseCase) Delete(id string) error {
	uom, err := u.FindById(id)
	if err != nil {
		return err // hanya return err karena sama dengan validasi err yang ada di find by id
	}

	err = u.repo.DeleteById(uom.Id)
	if err != nil {
		return fmt.Errorf("failed to delete uom: %v", err)
	}
	return nil
}

// Constructor
func NewUomUseCase(repo repository.UomRepository) UomUseCase {
	// kenapa menggunakan pointer "&" karena ngirim by reference
	return &uomUseCase{repo: repo}
}
