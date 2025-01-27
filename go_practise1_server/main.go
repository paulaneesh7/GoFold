package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductId string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price float64 `json:"price"`
	Category *Category `json:"category"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}

type Category struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"` // A brief description of the category
}

type Manufacturer struct {
	ManufactuererName    string `json:"manufacturer_name"`
	ManufacturingCountry string `json:"manufacturing_country"`
}

func (p *Product) IsEmpty() bool {
	return p.ProductId == "" && p.ProductName == "" && p.Price == 0 && p.Category == nil && p.Manufacturer == nil
}


var products []Product

func main() {
	fmt.Println("Http Server In Golang")

	r := mux.NewRouter();

	r.HandleFunc("/api/products", CreateProduct).Methods("POST")
	r.HandleFunc("/", HomeRoute).Methods("GET")
	r.HandleFunc("/api/products", GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductById).Methods("GET")
	r.HandleFunc("/api/products/{id}", UpdateProductById).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductById).Methods("DELETE")
	


	// seed some dummy data into the products slice
	products = append(products, Product{"1", "Laptop", 699, &Category{"Laptop", "A laptop"}, &Manufacturer{"Dell", "USA"}})
	products = append(products, Product{"2", "Phone", 299, &Category{"Phone", "A phone"}, &Manufacturer{"Apple", "USA"}})
	products = append(products, Product{"3", "TV", 399, &Category{"TV", "A TV"}, &Manufacturer{"Samsung", "Korea"}})
	products = append(products, Product{"4", "Mac", 1099, &Category{"Laptop", "An Apple Laptop"}, &Manufacturer{"Apple", "USA"}})

	log.Fatal(http.ListenAndServe(":8080", r))
}



func HomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Route")
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("Home Route"))
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all products")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}


func GetProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting a Product by Id")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for _, product := range products {
		if product.ProductId == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("No product found with id %v", id))
}


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Product")
	w.Header().Set("Content-Type", "application/json")
	
	var product Product
	
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide input")
		return 
	}

	json.NewDecoder(r.Body).Decode(&product)

	if product.IsEmpty() {
		json.NewEncoder(w).Encode("No data found in the body")
		return
	}

	// generate a unique id (for product) and convert it into string (just generated it through autogeneration)
	product.ProductId = strconv.Itoa(len(products) + 1)

	// append product into products
	products = append(products, product)

	json.NewEncoder(w).Encode(product)
}


func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a Product by Id")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide input")
		return
	}

	var updatedProduct Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)

	for index, p := range products {
		if p.ProductId == id {
			products[index].ProductId = id
			products[index].ProductName = updatedProduct.ProductName
			products[index].Price = updatedProduct.Price
			products[index].Category = updatedProduct.Category
			products[index].Manufacturer = updatedProduct.Manufacturer
			json.NewEncoder(w).Encode(products[index])
			return
		}
	}


	json.NewEncoder(w).Encode(fmt.Sprintf("No product found with id %v", id))
}


func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a Product by Id")
	w.Header().Set("Content-Type", "applcation/json")

	params := mux.Vars(r)
	id := params["id"]

	for index, p := range products {
		if p.ProductId == id {
			products = append(products[:index], products[index+1:]...)
			json.NewEncoder(w).Encode("Product deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("No product found with id %v", id))
}


