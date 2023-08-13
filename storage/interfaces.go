package storage

import "github.com/mostafasolati/catalog/models"

type (
	Product interface {
		Create(*models.Product) error
		Update(*models.Product) error
		Find(int) (*models.Product, error)
		Delete(int) error
		List(categoryID int) ([]*models.Product, error)
	}

	Category interface {
		Create(*models.Category) error
		Update(*models.Category) error
		Find(int) (*models.Category, error)
		Delete(int) error
		List() ([]*models.Category, error)
	}
)
