package shufersal

type ShufersalProductPrice struct {
	value          float64
	formattedValue string
}

func MakeShufersalProductPrice(value float64, formattedValue string) ShufersalProductPrice {
	return ShufersalProductPrice{value, formattedValue}
}

type ShufersalProduct struct {
	name  string
	price ShufersalProductPrice
	code  string
}

func MakeShufersalProduct(name string, code string, price ShufersalProductPrice) ShufersalProduct {
	return ShufersalProduct{name, price, code}
}

func (r *ShufersalProduct) GetPrice() float64 {
	return r.price.value
}
