package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (s *app) websiteIndexHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := s.categoryService.List()
	if err != nil {
		log.Fatal(err)
	}
	s.tmpl.ExecuteTemplate(w, "index.html", map[string]any{
		"categories": categories,
	})
}

func (s *app) websiteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	id, _ := strconv.Atoi(parts[2])
	products, err := s.productService.List(id)
	if err != nil {
		log.Fatal(err)
	}

	categories, err := s.categoryService.List()
	if err != nil {
		log.Fatal(err)
	}

	category, err := s.categoryService.Find(id)
	if err != nil {
		log.Fatal(err)
	}

	s.tmpl.ExecuteTemplate(w, "products.html", map[string]any{
		"categories": categories,
		"category":   category,
		"products":   products,
	})
}

func (s *app) websiteProductHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	id, _ := strconv.Atoi(parts[2])

	product, err := s.productService.Find(id)
	if err != nil {
		log.Fatal(err)
	}

	categories, err := s.categoryService.List()
	if err != nil {
		log.Fatal(err)
	}

	category, err := s.categoryService.Find(product.CategoryID)
	if err != nil {
		log.Fatal(err)
	}

	s.tmpl.ExecuteTemplate(w, "single.html", map[string]any{
		"categories": categories,
		"category":   category,
		"product":    product,
	})
}
