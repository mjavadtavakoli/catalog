package services

import (
	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/storage"
	"github.com/mostafasolati/catalog/validator"
)

type CategoryInterface interface {
	Create(*models.Category) error
	Update(*models.Category) error
	Find(int) (*models.Category, error)
	Delete(int) error
	List() ([]*models.Category, error)
}

type categoryService struct {
	storage storage.Category
}

func NewCategory(storage storage.Category) CategoryInterface {
	return &categoryService{
		storage: storage,
	}
}

func (s *categoryService) Create(category *models.Category) error {
	err := validator.New().
		AddRule(validator.String("عنوان", category.Title)).
		AddRule(validator.String("تصویر", category.Image)).
		Validate()
	if err != nil {
		return err
	}
	return s.storage.Create(category)
}

func (s *categoryService) Update(category *models.Category) error {
	err := validator.New().
		AddRule(validator.String("عنوان", category.Title)).
		AddRule(validator.String("تصویر", category.Image)).
		Validate()
	if err != nil {
		return err
	}
	return s.storage.Update(category)
}

func (s *categoryService) Find(id int) (*models.Category, error) {
	err := validator.New().
		AddRule(validator.Number("آیدی", id)).
		Validate()
	if err != nil {
		return nil, err
	}
	return s.storage.Find(id)
}

func (s *categoryService) Delete(id int) error {
	err := validator.New().
		AddRule(validator.Number("آیدی", id)).
		Validate()
	if err != nil {
		return err
	}
	return s.storage.Delete(id)
}

func (s *categoryService) List() ([]*models.Category, error) {
	return s.storage.List()
}
