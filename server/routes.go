package server

import (
	"net/http"
	"path"
)

func (s *app) routes() {

	// api
	http.HandleFunc("/api/v1/upload", uploadFileHandler)
	http.HandleFunc("/api/v1/category", s.categoryHandler())
	http.HandleFunc("/api/v1/category/", s.categoryHandler())
	http.HandleFunc("/api/v1/products", s.productHandler())
	http.HandleFunc("/api/v1/products/", s.productHandler())

	// public folders
	for _, folderPath := range []string{"template/css", "template/js", "template/fonts", "upload"} {
		folder := path.Base(folderPath)
		fs := http.FileServer(http.Dir(folderPath))
		folder = "/" + folder + "/"
		http.Handle(folder, http.StripPrefix(folder, fs))
	}

	//website
	http.HandleFunc("/", s.websiteIndexHandler)
	http.HandleFunc("/category/", s.websiteCategoryHandler)
	http.HandleFunc("/product/", s.websiteProductHandler)
}
