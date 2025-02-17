package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FactResponse struct {
	Teks string `json:"text"`
	Tipe string `json:"type"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", listProducts)
	mux.HandleFunc("POST /products", createProduct)
	// mux.HandleFunc("PUT /products/{id}", updateProduct)
	// mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	port := os.Getenv("APP_PORT")

	fmt.Println("running from port :", port)

	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	server.ListenAndServe()
}

var database = map[int]Product{
	1: {ID: 1, Name: "Product 1", Price: 10000},
	2: {ID: 2, Name: "Product 2", Price: 20000},
}

var lastID int = 1

func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product

	for _, p := range database {
		products = append(products, p)
	}

	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi Kesalahan"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
	}

	var product Product
	err = json.Unmarshal(bodyByte, &product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
	}

	lastID++

	product.ID = lastID

	database[int(product.ID)] = product

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("Sukses membuat data product"))
}

// func updateProduct(w http.ResponseWriter, r *http.Request) {
// 	productId := r.PathValue("id")

// 	productIDInt, err := strconv.Atoi(productId)
// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(400)
// 		w.Write([]byte("Kesalahan dalam request"))
// 	}

// 	bodyByte, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(400)
// 		w.Write([]byte("Kesalahan dalam request"))
// 	}

// 	var product Product
// 	err = json.Unmarshal(bodyByte, &product)
// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(400)
// 		w.Write([]byte("Kesalahan dalam request"))
// 	}

// 	product.ID = productIDInt
// 	database[productIDInt] = product

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	w.Write([]byte("Sukses mengubah data product"))
// }

// func deleteProduct(w http.ResponseWriter, r *http.Request) {
// 	productId := r.PathValue("id")

// 	productIDInt, err := strconv.Atoi(productId)
// 	if err != nil {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(400)
// 		w.Write([]byte("Kesalahan dalam request"))
// 	}

// 	delete(database, productIDInt)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(204)
// }
