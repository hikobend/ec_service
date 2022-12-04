package main

import (
	"database/sql"
	"log"
	"strconv"
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
	r.PATCH("/product/:id", AllUpdate)                     // 全て更新
	r.PATCH("/product/:id/name", NameUpdate)               // 名前を更新
	r.PATCH("/product/:id/category", CategoryUpdate)       // カテゴリーを更新
	r.PATCH("/product/:id/price", PriceUpdate)             // 値段を更新
	r.PATCH("/product/:id/stock", StockUpdate)             // 在庫を更新
	r.PATCH("/product/:id/brand", BrandUpdate)             // ブランドを更新
	r.PATCH("/product/:id/description", DescriptionUpdate) // 説明を更新

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

func NameUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var json Product_JSON
	c.ShouldBindJSON(&json)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	update, err := db.Prepare("UPDATE product SET  name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	update.Exec(json.Name, id)
}

func AllUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var json Product_JSON
	c.ShouldBindJSON(&json)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	update, err := db.Prepare("UPDATE product SET name = ?, category = ?, price = ?, stock = ?, brand = ?, description = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	update.Exec(json.Name, json.Category, json.Price, json.Stock, json.Brand, json.Description, id)
}

func CategoryUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var json Product_JSON
	c.ShouldBindJSON(&json)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	update, err := db.Prepare("UPDATE product SET category = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	update.Exec(json.Category, id)
}

func PriceUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var json Product_JSON
	c.ShouldBindJSON(&json)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	update, err := db.Prepare("UPDATE product SET price = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	update.Exec(json.Price, id)
}

func StockUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var json Product_JSON
	c.ShouldBindJSON(&json)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	update, err := db.Prepare("UPDATE product SET stock = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	update.Exec(json.Stock, id)
}

func BrandUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var json Product_JSON
	c.ShouldBindJSON(&json)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	update, err := db.Prepare("UPDATE product SET brand = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	update.Exec(json.Brand, id)
}

func DescriptionUpdate(c *gin.Context) {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/local?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var json Product_JSON
	c.ShouldBindJSON(&json)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	update, err := db.Prepare("UPDATE product SET description = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	update.Exec(json.Description, id)
}
