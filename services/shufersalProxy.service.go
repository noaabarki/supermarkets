package services

type ShufersalProductPrice struct {
	value          float32
	formattedValue string
}

type ShufersalProduct struct {
	name  string
	price ShufersalProductPrice
	code  string
}

type IShufersalProxy interface {
	getProductPrice(productName string) ShufersalProduct
	GetProducts(productName string) []ShufersalProduct
}

type ShufersalProxy struct {
}

const endpointURL string = "https://www.shufersal.co.il/online/he/search/results"
const query string = "?q=%d&limit=10"

func MakeShufersalProxy() ShufersalProxy {
	return ShufersalProxy{}
}

func (proxy ShufersalProxy) getProductPrice(productName string) float32 {
	// get from server
	productPrice := ShufersalProductPrice{value: 9.90, formattedValue: "9.90#"}
	product := ShufersalProduct{name: productName, price: productPrice, code: "P_22"}
	return product.price.value
}

func (proxy ShufersalProxy) GetProducts(productName string) []ShufersalProduct {
	// get from server
	productPrice := ShufersalProductPrice{value: 9.90, formattedValue: "9.90#"}
	product := ShufersalProduct{name: productName, price: productPrice, code: "P_22"}
	return []ShufersalProduct{product}
}
