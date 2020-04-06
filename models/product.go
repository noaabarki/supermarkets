package models

type Product struct {
	name   string
	prices []ProductPrice
}

type ProductPrice struct {
	storeName string
	price     float64
}

func MakeProduct(name string, prices []ProductPrice) Product {
	return Product{name, prices}
}

func MakeProductPrice(storeName string, price float64) ProductPrice {
	return ProductPrice{storeName, price}
}

func NewProductPrice(storeName string, price float64) *ProductPrice {
	p := ProductPrice{storeName: storeName}
	p.price = price
	return &p
}

func GetName() string {
	return "name"
}
