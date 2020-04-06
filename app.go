package main

import (
	"fmt"
	"net/http"
	"supermarkets/controllers"
)

func main() {
	fmt.Println("supermarkets")
	http.HandleFunc("/products", controllers.GetProducts)
	http.ListenAndServe(":8080", nil)
}
