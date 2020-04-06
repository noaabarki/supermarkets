package services

import (
	"supermarkets/dtos/shufersal"
	"supermarkets/models"
)

type IProductService interface {
	GetProducts() []models.Product
}

type ProductService struct {
	shufersalProxy ShufersalProxy
}

func MakeProductService() ProductService {
	return ProductService{shufersalProxy: MakeShufersalProxy()}
}

func (productsService ProductService) GetProducts(productName string) []models.Product {
	var results []models.Product
	results = append(results, productsService.getShufersalProducts(productName)...)
	return results
}

func (productsService ProductService) getShufersalProducts(productName string) []models.Product {
	var results []models.Product

	var shufersalProducts []shufersal.ShufersalProduct
	shufersalProducts = productsService.shufersalProxy.GetProducts(productName)

	prices := make([]models.ProductPrice, len(shufersalProducts))
	for _, shufersalProduct := range shufersalProducts {
		price := shufersalProduct.GetPrice()
		prices = append(prices, models.MakeProductPrice("shufersal", price))
	}

	results = append(results, models.MakeProduct(productName, prices))
	return results
}
