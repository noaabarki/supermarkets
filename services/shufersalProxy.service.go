package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"supermarkets/dtos/shufersal"
)

type IShufersalProxy interface {
	// getProductPrice(productName string) shufersal.ShufersalProduct
	GetProducts(productName string) []shufersal.ShufersalProduct
}

type ShufersalProxy struct {
}

const endpointURL string = "https://www.shufersal.co.il/online/he/search/results"
const query string = "?q=%d&limit=10"

func MakeShufersalProxy() ShufersalProxy {
	return ShufersalProxy{}
}

// func (proxy ShufersalProxy) getProductPrice(productName string) float32 {
// 	productPrice := shufersal.ShufersalProductPrice{value: 9.90, formattedValue: "9.90#"}
// 	product := shufersal.ShufersalProduct{name: productName, price: productPrice, code: "P_22"}
// 	return product.price.value
// }

func (proxy ShufersalProxy) GetProducts(productName string) []shufersal.ShufersalProduct {
	// get from server
	var shufersalProducts []shufersal.ShufersalProduct
	var client http.Client

	query := fmt.Sprintf("?q=%s&limit=10", url.QueryEscape(productName))
	URL := "https://www.shufersal.co.il/online/he/search/results" + query
	response, error := client.Get(URL)

	if error != nil {
		log.Fatal(error)
	}

	// log.Println(bodyString)
	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
	}

	bodyString := string(body)

	var shufersalResponseObj shufersal.ShufersalResponse
	error = json.Unmarshal([]byte(bodyString), &shufersalResponseObj)

	if error != nil {
		log.Fatal(error)
	}
	for _, item := range shufersalResponseObj.Results {
		productPrice := shufersal.MakeShufersalProductPrice(item.Price.Value, item.Price.FormattedValue)
		product := shufersal.MakeShufersalProduct(productName, item.Code, productPrice)
		shufersalProducts = append(shufersalProducts, product)
	}
	// productPrice := ShufersalProductPrice{value: 9.90, formattedValue: "9.90#"}
	// product := ShufersalProduct{name: productName, price: productPrice, code: "P_22"}

	return shufersalProducts
}
