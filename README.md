# REST API

A REST API written in Go-Lang

- Lang: Go-Lang
- DB: MySQL

Routes
```go
// get all products
// http://localhost:8083/products
router.GET("/products", getProducts)

// get a specific product using product ID
// http://localhost:8083/product/<product-code>
router.GET("/product/:code", getProduct)

// add products
// http://localhost:8083/product/add
//	{
//		"code": "<product-code>",
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
//		"code": "<product-code>",
//		"name": "<product-name>",
//		"qty": <product-quantity>
//	}
router.PUT("/product/update/:code", updateProduct)
```
