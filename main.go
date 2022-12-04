package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Product_JSON struct { // JSON
	Name        string `json:"name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
}

type Product struct { // DB
	Id          int
	Name        string
	Category    string
	Price       int
	Stock       int
	Brand       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	r := gin.Default()
	r.POST("/create", Insert)

	r.Run()
}

func Insert(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var product Product_JSON
	c.ShouldBindJSON(&product)

	insert, err := db.Prepare("INSERT INTO product(name, category, price, stock, brand, description) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	insert.Exec(product.Name, product.Category, product.Price, product.Stock, product.Brand, product.Description)
}
