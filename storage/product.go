package storage

import (
	"database/sql"
	"errors"

	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/validator"
)

type ProductStorage struct {
	db *sql.DB
}

func NewProductSQL(db *sql.DB) Product {
	return &ProductStorage{
		db: db,
	}
}

func (db *ProductStorage) Find(id int) (*models.Product, error) {
	query := `
		SELECT id, price, category_id, weight, title, image, pdf,description
		FROM products
		WHERE id = $1`
	row := db.db.QueryRow(query, id)

	product := &models.Product{}
	err := row.Scan(
		&product.ID,
		&product.Price,
		&product.CategoryID,
		&product.Weight,
		&product.Title,
		&product.Image,
		&product.PDF,
		&product.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(validator.ErrProductNotFound)
		}
		return nil, err
	}

	return product, nil
}

func (db *ProductStorage) Delete(id int) error {
	if _, err := db.Find(id); err != nil {
		return err
	}

	query := `
		DELETE FROM products
		WHERE id = $1`
	_, err := db.db.Exec(query, id)
	return err
}

func (db *ProductStorage) List(categoryID int) ([]*models.Product, error) {
	query := `
		SELECT id, price, category_id, weight, title, image, pdf,description
		FROM products
		WHERE category_id = $1`
	rows, err := db.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*models.Product, 0)
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Price,
			&product.CategoryID,
			&product.Weight,
			&product.Title,
			&product.Image,
			&product.PDF,
			&product.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (db *ProductStorage) Create(product *models.Product) error {
	if product.ID != 0 {
		return errors.New(validator.ErrAlreadyHasID)
	}

	query := `
		INSERT INTO products (price, category_id, weight, title, description, image, pdf)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.db.QueryRow(query, product.Price, product.CategoryID, product.Weight, product.Title, product.Description, product.Image, product.PDF).
		Scan(&product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db *ProductStorage) Update(product *models.Product) error {
	if _, err := db.Find(product.ID); err != nil {
		return err
	}

	query := `
		UPDATE products
		SET price = $1, category_id = $2, weight = $3, title = $4, description = $5, image = $6, pdf = $7
		WHERE id = $8`
	_, err := db.db.Exec(query, product.Price, product.CategoryID, product.Weight, product.Title, product.Description, product.Image, product.PDF, product.ID)
	return err
}
