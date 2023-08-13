package storage

import (
	"errors"
	"sync"

	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/validator"
)

type categoryInMemory struct {
	categories     map[int]*models.Category
	titles         map[string]*models.Category
	nextCategoryID int
	sync.RWMutex
}

func NewCategoryInMemory() *categoryInMemory {
	return &categoryInMemory{
		categories:     make(map[int]*models.Category),
		titles:         make(map[string]*models.Category),
		nextCategoryID: 1,
	}
}

func (db *categoryInMemory) Create(category *models.Category) error {
	db.Lock()
	defer db.Unlock()

	if category.ID != 0 {
		return errors.New(validator.ErrAlreadyHasID)
	}
	if _, ok := db.titles[category.Title]; ok {
		return errors.New(validator.ErrCategoryExists)
	}

	category.ID = db.nextCategoryID
	db.categories[category.ID] = category
	db.titles[category.Title] = category
	db.nextCategoryID++

	return nil
}

func (db *categoryInMemory) Update(category *models.Category) error {
	db.Lock()
	defer db.Unlock()

	if _, exists := db.categories[category.ID]; !exists {
		return errors.New(validator.ErrCategoryNotFound)
	}

	db.categories[category.ID] = category
	return nil
}

func (db *categoryInMemory) Delete(id int) error {
	db.Lock()
	defer db.Unlock()

	var category *models.Category
	var exists bool

	if category, exists = db.categories[id]; !exists {
		return errors.New(validator.ErrCategoryNotFound)
	}

	delete(db.categories, id)
	delete(db.titles, category.Title)
	return nil
}

func (db *categoryInMemory) Find(id int) (*models.Category, error) {
	db.RLock()
	defer db.RUnlock()

	if category, exists := db.categories[id]; exists {
		return category, nil
	}

	return nil, errors.New(validator.ErrCategoryNotFound)
}

func (db *categoryInMemory) List() ([]*models.Category, error) {
	categories := make([]*models.Category, 0, len(db.categories))
	for _, category := range db.categories {
		categories = append(categories, category)
	}
	return categories, nil
}
