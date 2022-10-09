package main

import (
	"net/http"

	mods "go-rest-api-app/db_handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/product/:code", getProduct)
	router.POST("/products", addProduct)
	router.DELETE("/delete/:code", deleteProduct)
	router.Run("localhost:8083")
}

func getProducts(c *gin.Context) {
	products := mods.GetProducts()

	// products == nil || len(products) == 0
	if len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

func getProduct(c *gin.Context) {
	code := c.Param("code")

	product := mods.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func addProduct(c *gin.Context) {
	var prod mods.Product

	err := c.BindJSON(&prod)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		mods.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}
}

func deleteProduct(c *gin.Context) {
	code := c.Param("code")

	del := mods.DeleteProduct(code)

	if !del {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, del)
	}
}
