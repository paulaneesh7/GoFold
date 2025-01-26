package main

import "fmt"

type Product struct {
	ProductId string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price float64 `json:"price"`
	Category Category `json:"category"`
	Manufacturer Manufacturer `json:"manufacturer"`
}

type Category struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"` // A brief description of the category
}

type Manufacturer struct {
	ManufactuererName    string `json:"manufacturer_name"`
	ManufacturingCountry string `json:"manufacturing_country"`
}

func main() {
	fmt.Println("Http Server In Golang")
}
