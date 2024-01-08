package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/paulosarmento/Arquitetura-Hexagonal/adapters/db"
	"github.com/paulosarmento/Arquitetura-Hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemplo", 30)

	productService.Enable(product)
}
