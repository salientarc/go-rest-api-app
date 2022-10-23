package main

import (
	"net/http"

	mods "go-rest-api-app/db_handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// get all products
	// http://localhost:8083/products
	router.GET("/products", getProducts)

	// get a specific product using product ID
	// http://localhost:8083/product/<product-code>
	router.GET("/product/:code", getProduct)

	// add products
	// http://localhost:8083/product/add
	//	{
	//		"code": "<product-code>,
	//		"name": "<product-name>",
	//		"qty": <product-quantity>
	//	}
	router.POST("/product/add", addProduct)

	// delete a product using product code
	// http://localhost:8083/product/delete/<product-code>
	router.DELETE("/product/delete/:code", deleteProduct)

	// update product details using product ID
	// http://localhost:8083/product/update/<product-code>
	//	{
	//		"code": "<product-code>,
	//		"name": "<product-name>",
	//		"qty": <product-quantity>
	//	}
	router.PUT("/product/update/:code", updateProduct)

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

func updateProduct(c *gin.Context) {
	code := c.Param("code")
	var prod mods.Product

	err := c.BindJSON(&prod)
	mods.ErrorCheck(err)

	update := mods.UpdateProduct(code, prod)

	if !update {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, update)
	}
}
