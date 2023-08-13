package services

import (
	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/storage"
	"github.com/mostafasolati/catalog/validator"
)

type (
	ProductInterface interface {
		Create(*models.Product) error
		Update(*models.Product) error
		Find(int) (*models.Product, error)
		Delete(int) error
		List(categoryID int) ([]*models.Product, error)
	}

	productService struct {
		storage storage.Product
	}
)

func NewProduct(storage storage.Product) ProductInterface {
	return &productService{
		storage: storage,
	}
}

func (s *productService) Create(product *models.Product) error {
	err := validator.New().
		AddRule(validator.Number("مبلغ", product.Price)).
		AddRule(validator.Number("وزن", product.Weight)).
		AddRule(validator.String("عنوان", product.Title)).
		AddRule(validator.String("فایل پی دی اف", product.PDF)).
		AddRule(validator.Number("دسته بندی", product.CategoryID)).
		AddRule(validator.String("توضیحات", product.Description)).
		Validate()
	if err != nil {
		return err
	}

	return s.storage.Create(product)
}

func (s *productService) Update(product *models.Product) error {

	err := validator.New().
		AddRule(validator.Number("id", product.ID)).
		AddRule(validator.Number("مبلغ", product.Price)).
		AddRule(validator.Number("وزن", product.Weight)).
		AddRule(validator.String("عنوان", product.Title)).
		AddRule(validator.String("فایل پی دی اف", product.PDF)).
		AddRule(validator.Number("دسته بندی", product.CategoryID)).
		AddRule(validator.String("توضیحات", product.Description)).
		Validate()

	if err != nil {
		return err
	}

	return s.storage.Update(product)
}

func (s *productService) Find(id int) (*models.Product, error) {
	err := validator.New().AddRule(validator.Number("id", id)).Validate()
	if err != nil {
		return nil, err
	}
	return s.storage.Find(id)
}

func (s *productService) Delete(id int) error {
	err := validator.New().AddRule(validator.Number("id", id)).Validate()
	if err != nil {
		return err
	}
	return s.storage.Delete(id)
}

func (s *productService) List(categoryID int) ([]*models.Product, error) {
	return s.storage.List(categoryID)
}
