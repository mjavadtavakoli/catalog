package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	filePath := filepath.Join("upload", header.Filename)
	output, err := os.Create(filePath)
	defer output.Close()
	io.Copy(output, file)
	InfoResponse(w, filePath)
}
