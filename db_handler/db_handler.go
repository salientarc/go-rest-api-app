package dbhandler

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	ErrorCheck(err)
}

func GetProducts() []Product {
	LoadEnv()
	db, err := sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@tcp(127.0.0.1:3306)/"+os.Getenv("DBNAME"))
	ErrorCheck(err)

	defer db.Close()

	results, err := db.Query("SELECT * FROM product")

	ErrorCheck(err)

	products := []Product{}

	for results.Next() {
		var prod Product

		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)
		ErrorCheck(err)

		products = append(products, prod)
		//fmt.Println("product.code :", prod.Code+" : "+prod.Name)
	}

	return products
}

func GetProduct(code string) *Product {
	LoadEnv()
	db, err := sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@tcp(127.0.0.1:3306)/"+os.Getenv("DBNAME"))

	prod := &Product{}
	ErrorCheck(err)

	defer db.Close()

	results, err := db.Query("SELECT * FROM product where code=?", code)
	ErrorCheck(err)

	if results.Next() {
		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return prod
}

func AddProduct(product Product) {
	LoadEnv()
	db, err := sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@tcp(127.0.0.1:3306)/"+os.Getenv("DBNAME"))
	ErrorCheck(err)

	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO product (code,name,qty,last_updated) VALUES (?,?,?, now())",
		product.Code, product.Name, product.Qty)
	ErrorCheck(err)

	defer insert.Close()
}

func DeleteProduct(code string) bool {
	LoadEnv()
	db, err := sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@tcp(127.0.0.1:3306)/"+os.Getenv("DBNAME"))
	ErrorCheck(err)

	defer db.Close()

	results, err := db.Query("DELETE FROM product WHERE code=?", code)
	ErrorCheck(err)

	defer results.Close()

	return true
}
