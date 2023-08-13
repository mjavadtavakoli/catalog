package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mostafasolati/catalog/models"
)

func (s *app) categoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			category := new(models.Category)
			if Bind(w, r, category) != nil {
				return
			}
			err := s.categoryService.Create(category)
			if err != nil {
				ErrorResponse(w, err)
				return
			}
			Response(w, category)

		case http.MethodPut:
			category := new(models.Category)
			if Bind(w, r, category) != nil {
				return
			}
			err := s.categoryService.Update(category)
			if err != nil {
				ErrorResponse(w, err)
				return
			}
			Response(w, category)

		case http.MethodDelete:
			parts := strings.Split(r.URL.String(), "/")
			id, _ := strconv.Atoi(parts[4])
			err := s.categoryService.Delete(id)
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
				category, err := s.categoryService.Find(id)
				if err != nil {
					ErrorResponse(w, err)
					return
				}
				Response(w, category)
				return
			}

			// list
			categories, err := s.categoryService.List()
			if err != nil {
				ErrorResponse(w, err)
			}
			Response(w, categories)
		}

	}
}
