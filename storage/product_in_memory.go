package storage

import (
	"errors"
	"sync"

	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/validator"
)

type productInMemory struct {
	products      map[int]*models.Product
	titles        map[string]*models.Product
	nextProductID int
	sync.RWMutex
}

func NewProductInMemory() *productInMemory {
	return &productInMemory{
		products:      make(map[int]*models.Product),
		titles:        make(map[string]*models.Product),
		nextProductID: 1,
	}
}

func (db *productInMemory) Create(product *models.Product) error {
	db.Lock()
	defer db.Unlock()

	if product.ID != 0 {
		return errors.New(validator.ErrAlreadyHasID)
	}
	if _, ok := db.titles[product.Title]; ok {
		return errors.New(validator.ErrProductExists)
	}

	product.ID = db.nextProductID
	db.products[product.ID] = product
	db.titles[product.Title] = product
	db.nextProductID++

	return nil
}

func (db *productInMemory) Update(product *models.Product) error {

	db.Lock()
	defer db.Unlock()

	if _, exists := db.products[product.ID]; !exists {
		return errors.New(validator.ErrProductNotFound)
	}

	db.products[product.ID] = product
	return nil
}

func (db *productInMemory) Delete(id int) error {
	db.Lock()
	defer db.Unlock()

	var product *models.Product
	var exists bool

	if product, exists = db.products[id]; !exists {
		return errors.New(validator.ErrProductNotFound)
	}

	delete(db.products, id)
	delete(db.titles, product.Title)
	return nil
}

func (db *productInMemory) Find(id int) (*models.Product, error) {
	db.RLock()
	defer db.RUnlock()

	if product, exists := db.products[id]; exists {
		return product, nil
	}

	return nil, errors.New(validator.ErrProductNotFound)
}

func (db *productInMemory) List(categoryID int) ([]*models.Product, error) {
	products := make([]*models.Product, 0, len(db.products))
	for _, product := range db.products {
		products = append(products, product)
	}
	return products, nil
}
