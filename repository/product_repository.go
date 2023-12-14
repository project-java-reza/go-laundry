package repository

import (
	"database/sql"
	"laundry/model"
)

type ProductRepository interface {
	Save(product model.Product) error
	FindById(id string) (model.Product, error)
	FindAll() ([]model.Product, error)
	FindByName(name string) ([]model.Product, error)
	Update(product model.Product) error
	DeleteById(id string) error
}

type productRepository struct {
	db *sql.DB
}

func (p productRepository) Save(product model.Product) error {
	_, err := p.db.Exec("INSERT INTO m_products VALUES ($1, $2, $3, $4)",
		product.Id,
		product.Name,
		product.Price,
		product.Uom.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p productRepository) FindById(id string) (model.Product, error) {
	row := p.db.QueryRow(`
	SELECT p.id, p.name, p.price, u.id, u.description 
	FROM m_products p 
	JOIN m_uom u ON u.id = p.uom_id 
	WHERE p.id = $1`, id)
	product := model.Product{}
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Uom.Id,
		&product.Uom.Description,
	)

	if err != nil {
		return model.Product{}, err
	}
	return product, nil

}

func (p productRepository) FindByName(name string) ([]model.Product, error) {
	rows, err := p.db.Query(`
	SELECT p.id, p.name, p.price, u.id, u.description 
	FROM m_products p 
	JOIN m_uom u ON u.id = p.uom_id 
	WHERE name ILIKE $1`, name)

	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Uom.Id,
			&product.Uom.Description,
		)

		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p productRepository) FindAll() ([]model.Product, error) {
	rows, err := p.db.Query(`SELECT p.id, p.name, p.price, u.id, u.description FROM m_products p 
	JOIN m_uom u ON u.id = p.uom_id`)

	if err != nil {
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Uom.Id,
			&product.Uom.Description,
		)

		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p productRepository) Update(product model.Product) error {
	_, err := p.db.Exec("UPDATE m_products SET name = $2, price = $3, uom_id = $4 WHERE id = $1",
		product.Id,
		product.Name,
		product.Price,
		product.Uom.Id,
	)
	return err
}

func (p productRepository) DeleteById(id string) error {
	_, err := p.db.Exec("DELETE FROM m_products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// constructor
func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
