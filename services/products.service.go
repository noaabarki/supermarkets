package services

import "supermarkets/models"

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

	shufersalProducts := productsService.shufersalProxy.GetProducts(productName)

	prices := make([]models.ProductPrice, len(shufersalProducts))
	for _, shufersalProduct := range shufersalProducts {
		prices = append(prices, models.MakeProductPrice("shufersal", shufersalProduct.price.value))
	}

	results = append(results, models.MakeProduct(productName, prices))
	return results
}
