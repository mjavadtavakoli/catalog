package services_test

import (
	"testing"

	"github.com/mostafasolati/catalog/models"
	"github.com/mostafasolati/catalog/services"
	"github.com/mostafasolati/catalog/storage"
)

func TestCategoryCRUD(t *testing.T) {
	categoryStorage := storage.NewCategoryInMemory()
	categoryService := services.NewCategory(categoryStorage)
	testCategories := []*models.Category{
		{Title: "Electronics", Image: "Pic1"},
		{Title: "Clothing", Image: "Pic2"},
		{Title: "Books", Image: "Pic3"},
	}

	t.Run("Create Category", func(t *testing.T) {
		for i, category := range testCategories {
			err := categoryService.Create(category)
			if err != nil {
				AssertNotError(t, err, i)
			}
		}
	})

	t.Run("Create Category Invalid", func(t *testing.T) {
		for i, category := range []*models.Category{
			{},
			{Title: ""},
			{Title: "ABC"},
			{Image: "Pic1"},
		} {
			err := categoryService.Create(category)
			AssertError(t, err, i)
		}
	})

	t.Run("Update Category", func(t *testing.T) {
		for i, category := range testCategories {
			category.Title += "_UPDATED"
			err := categoryService.Update(category)
			AssertNotError(t, err, i)
		}
	})

	t.Run("Update Category Invalid", func(t *testing.T) {
		for i, category := range []*models.Category{
			{},
			{ID: 1, Title: ""},
		} {
			err := categoryService.Update(category)
			AssertError(t, err, i)
		} 
	})

	t.Run("Find Category", func(t *testing.T) {
		category, err := categoryService.Find(3)
		AssertNotError(t, err, 0)
		if category.ID != 3 {
			t.Fatal("mismatch IDs")
		}
	})

	t.Run("Find Category Invalid", func(t *testing.T) {
		for i, id := range []int{-1, 0, 300} {
			_, err := categoryService.Find(id)
			AssertError(t, err, i)
		}
	})

	t.Run("Delete Category", func(t *testing.T) {
		err := categoryService.Delete(3)
		AssertNotError(t, err, 0)
		_, err = categoryService.Find(3)
		AssertError(t, err, 0)
	})

	t.Run("Delete Category Invalid", func(t *testing.T) {
		for i, id := range []int{-1, 0, 3, 300} {
			err := categoryService.Delete(id)
			AssertError(t, err, i)
		}
	})
}
