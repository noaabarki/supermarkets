package controllers

import (
	"supermarkets/models"
	"supermarkets/services"
)

func getProductPrices(productName string) []models.Product {
	productsService := services.MakeProductService()

	// productPrices := []models.ProductPrice{models.MakeProductPrice("shufersal", 9.90)}
	// products := []models.Product{models.MakeProduct(productName, productPrices)}

	products := productsService.GetProducts(productName)
	return products
}
