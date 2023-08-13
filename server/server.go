package server

import (
	"html/template"
	"net/http"

	"github.com/mostafasolati/catalog/services"
)

type (
	App interface {
		Start(string) error
	}

	app struct {
		categoryService services.CategoryInterface
		productService  services.ProductInterface
		tmpl            *template.Template
	}
)

func New(
	categoryService services.CategoryInterface,
	productService services.ProductInterface,
) App {
	tmpl, _ := template.New("website").ParseGlob("template/*.html")
	return &app{
		categoryService: categoryService,
		productService:  productService,
		tmpl:            tmpl,
	}
}

func (s *app) Start(host string) error {
	s.routes()
	return http.ListenAndServe(host, nil)
}
