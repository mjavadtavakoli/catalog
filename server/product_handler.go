package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mostafasolati/catalog/models"
)

func (s *app) productHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			product := new(models.Product)
			if Bind(w, r, product) != nil {
				return
			}
			err := s.productService.Create(product)
			if err != nil {
				ErrorResponse(w, err)
				return
			}
			Response(w, product)

		case http.MethodPut:
			product := new(models.Product)
			if Bind(w, r, product) != nil {
				return
			}
			err := s.productService.Update(product)
			if err != nil {
				ErrorResponse(w, err)
				return
			}
			Response(w, product)

		case http.MethodDelete:
			parts := strings.Split(r.URL.String(), "/")
			id, _ := strconv.Atoi(parts[4])
			err := s.productService.Delete(id)
			if err != nil {
				ErrorResponse(w, err)
				return
			}
			InfoResponse(w, "با موفقیت حذف شد")

		case http.MethodGet:
			parts := strings.Split(r.URL.String(), "/")
			// if id present
			if len(parts) == 5 {
				id, _ := strconv.Atoi(parts[4])
				user, err := s.productService.Find(id)
				if err != nil {
					ErrorResponse(w, err)
					return
				}
				Response(w, user)
				return
			}

			// list
			users, err := s.productService.List(0)
			if err != nil {
				ErrorResponse(w, err)
			}
			Response(w, users)
		}

	}
}
