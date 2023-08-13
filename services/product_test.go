package services_test

import (
	"testing"

	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/services"
	"github.com/mostafasolati/catalog/storage"
)

func AssertNotError(t *testing.T, err error, i int) {
	if err == nil {
		return
	}
	t.Fatalf("[test case %d] expected nil got %v", i, err)
}

func AssertError(t *testing.T, err error, i int) {
	if err != nil {
		return
	}
	t.Fatalf("[test case %d] expected error got nil", i)
}

var (
	productStorage = storage.NewProductInMemory()
	productService = services.NewProduct(productStorage)
	testProducts   = []*models.Product{
		{Price: 1000, Weight: 1.5, Title: "Samsung Galaxy S24", CategoryID: 1, PDF: "file1.pdf", Description: "Desc1"},
		{Price: 2000, Weight: 5.7, Title: "Playstation 5", CategoryID: 2, PDF: "file2.pdf", Description: "Desc2"},
		{Price: 3000, Weight: 2.1, Title: "Galaxy Buds Pro", CategoryID: 3, PDF: "file3.pdf", Description: "Desc3"},
	}
)

func TestProductCRUD(t *testing.T) {

	t.Run("Create Product", func(t *testing.T) {
		for i, product := range testProducts {
			err := productService.Create(product)
			if err != nil {
				AssertNotError(t, err, i)
			}
		}
	})

	t.Run("Create Product Invalid", func(t *testing.T) {
		for i, product := range []*models.Product{
			{},
			{Price: 0, Weight: 1.5, Title: "Samsung Galaxy S24"},
			{Price: 1000, Weight: 0, Title: "Samsung Galaxy S24"},
			{Price: 1000, Weight: 1.5},
			{Price: 1000, Weight: 1, Title: "Samsung Galaxy S24", CategoryID: 1, PDF: "file.pdf"},
			{Price: 1000, Weight: 2, Title: "Samsung Galaxy S24", CategoryID: 1, Description: "desc"},
			{Price: 1000, Weight: 3, Title: "Samsung Galaxy S24", PDF: "file.pdf", Description: "desc"},
		} {
			err := productService.Create(product)
			AssertError(t, err, i)
		}
	})

	t.Run("Update Product", func(t *testing.T) {
		for i, product := range testProducts {
			product.Title += "_UPDATED"
			err := productService.Update(product)
			AssertNotError(t, err, i)
		}
	})

	t.Run("Update Product Invalid", func(t *testing.T) {
		for i, product := range []*models.Product{
			{},
			{ID: 1, Price: 0, Weight: 1.5, Title: "Samsung Galaxy S24"},
			{ID: 2, Price: 1000, Weight: 0, Title: "Samsung Galaxy S24"},
			{ID: 3, Price: 1000, Weight: 1.5, Title: ""},
		} {
			err := productService.Update(product)
			AssertError(t, err, i)
		}
	})

	t.Run("Find Product", func(t *testing.T) {
		product, err := productService.Find(3)
		AssertNotError(t, err, 0)
		if product.ID != 3 {
			t.Fatal("mismatch IDs")
		}
	})

	t.Run("Find Product Invalid", func(t *testing.T) {
		for i, id := range []int{-1, 0, 300} {
			_, err := productService.Find(id)
			AssertError(t, err, i)
		}
	})

	t.Run("Delete Product", func(t *testing.T) {
		err := productService.Delete(3)
		AssertNotError(t, err, 0)
		_, err = productService.Find(3)
		AssertError(t, err, 0)
	})

	t.Run("Delete Product Invalid", func(t *testing.T) {
		for i, id := range []int{-1, 0, 3, 300} {
			err := productService.Delete(id)
			AssertError(t, err, i)
		}
	})

}
