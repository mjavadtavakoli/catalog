package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mostafasolati/catalog/server"
	"github.com/mostafasolati/catalog/services"
	"github.com/mostafasolati/catalog/storage"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("upload")
	if os.IsNotExist(err) {
		os.Mkdir("upload", 0777)
	}

	data, err := os.ReadFile("setup.sql")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(string(data))
	if err != nil {
		log.Fatal("couldn't run queries:", err)
	}

	categoryStorage := storage.NewCategorySQL(db)
	categoryService := services.NewCategory(categoryStorage)

	productStorage := storage.NewProductSQL(db)
	productService := services.NewProduct(productStorage)

	app := server.New(
		categoryService,
		productService,
	)

	app.Start("localhost:8080")
}
