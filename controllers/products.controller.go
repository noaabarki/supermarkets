package controllers

import (
	"log"
	"net/http"
	"supermarkets/services"
)

func GetProducts(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	productsService := services.MakeProductService()
	keys, ok := request.URL.Query()["name"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	productName := keys[0]
	products := productsService.GetProducts(productName)
	log.Println(products)
}
