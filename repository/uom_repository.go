package repository

import (
	"database/sql"
	"laundry/model"
)

type UomRepository interface {
	// CRUD
	Save(uom model.Uom) error              // INSERT
	FindById(id string) (model.Uom, error) // SELECT By Id
	FindAll() ([]model.Uom, error)         // SELECT * FROM m_uom
	Update(uom model.Uom) error            // Update
	DeleteById(id string)                  // Delete FROM m_uom where id = ?
}

// tidak dipakai di file lain
type uomRepository struct {
	db *sql.DB // field DB
}

func (u uomRepository) Save(uom model.Uom) error {
	_, err := u.db.Exec("INSERT INTO m_uom VALUES  ($1, $2)", uom.Id, uom.Description)
	if err != nil {
		return err
	}
	return nil
}

func (u uomRepository) FindById(id string) (model.Uom, error) {
	row := u.db.QueryRow("SELECT id, description FROM m_uom WHERE id = $1", id)
	var uom model.Uom
	err := row.Scan(&uom.Id, &uom.Description)
	if err != nil {
		return model.Uom{}, err
	}
	return uom, nil
}

func (u uomRepository) FindAll() ([]model.Uom, error) {
	//TODO implement me
	panic("implement me")
}

func (u uomRepository) Update(uom model.Uom) error {
	//TODO implement me
	panic("implement me")
}

func (u uomRepository) DeleteById(id string) {
	//TODO implement me
	panic("implement me")
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
