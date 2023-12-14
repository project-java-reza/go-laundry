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
	DeleteById(id string) error            // Delete FROM m_uom where id = ?
}

// tidak dipakai di file lain
// class
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
	rows, err := u.db.Query("SELECT * FROM m_uom")
	if err != nil {
		return nil, err
	}
	var uoms []model.Uom
	for rows.Next() {
		var uom model.Uom
		err := rows.Scan(&uom.Id, &uom.Description)
		if err != nil {
			return nil, err
		}
		uoms = append(uoms, uom)
	}

	// pesan error jika kegagalan dalam pengambilan query
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return uoms, nil
}

func (u uomRepository) Update(uom model.Uom) error {
	_, err := u.db.Exec("UPDATE m_uom SET description= $1 where id= $2", uom.Description, uom.Id)
	if err != nil {
		return err
	}
	return nil
}

//func (u uomRepository) Update(uom model.Uom) error {
//	query := "UPDATE m_uom set description = $1 where id=$2"
//
//	_, err := u.db.Exec(query, uom.Description, uom.Id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (u uomRepository) DeleteById(id string) error {
	_, err := u.db.Exec("DELETE FROM m_uom WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
