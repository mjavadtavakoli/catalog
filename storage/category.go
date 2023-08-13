package storage

import (
	"database/sql"
	"errors"

	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/validator"
)

type categorySQL struct {
	db *sql.DB
}

func NewCategorySQL(db *sql.DB) Category {
	return &categorySQL{
		db: db,
	}
}

func (db *categorySQL) Create(category *models.Category) error {
	if category.ID != 0 {
		return errors.New(validator.ErrAlreadyHasID)
	}

	query := "INSERT INTO categories (title, image) VALUES ($1, $2) RETURNING id"
	err := db.db.QueryRow(query, category.Title, category.Image).Scan(&category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (db *categorySQL) Update(category *models.Category) error {
	if _, err := db.Find(category.ID); err != nil {
		return err
	}

	query := "UPDATE categories SET title = $1, image = $2 WHERE id = $3"
	_, err := db.db.Exec(query, category.Title, category.Image, category.ID)
	return err
}

func (db *categorySQL) Find(id int) (*models.Category, error) {
	query := "SELECT id, title, image FROM categories WHERE id = $1"
	row := db.db.QueryRow(query, id)

	category := &models.Category{}
	err := row.Scan(&category.ID, &category.Title, &category.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(validator.ErrCategoryNotFound)
		}
		return nil, err
	}

	return category, nil
}

func (db *categorySQL) List() ([]*models.Category, error) {
	query := "SELECT id, title, image FROM categories"
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]*models.Category, 0)
	for rows.Next() {
		category := &models.Category{}
		err := rows.Scan(&category.ID, &category.Title, &category.Image)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (db *categorySQL) Delete(id int) error {
	if _, err := db.Find(id); err != nil {
		return err
	}

	query := "DELETE FROM categories WHERE id = $1"
	_, err := db.db.Exec(query, id)
	return err
}
