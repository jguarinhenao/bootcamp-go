package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products = make(map[int]Product)
var nextID = 1

func main() {
	startAPI("products.json")
}

func startAPI(filePath string) {
	productsFromFile, err := loadProductsFromJSON(filePath)
	if err != nil {
		fmt.Printf("Error loading products from JSON: %v\n", err)
		return
	}

	initializeProducts(productsFromFile)

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Status OK (200)
		w.Write([]byte("pong"))
	})

	r.Route("/products", func(rt chi.Router) {

		rt.Get("/", getAllProducts)

		rt.Get("/{id}", getProductByID)

		rt.Get("/search", getProductsByQueryParamsPriceGT)

		rt.Post("/", createProduct)
	})

	http.ListenAndServe(":8080", r)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	productsList := make([]Product, 0, len(products))
	for _, product := range products {
		productsList = append(productsList, product)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Status OK (200)
	json.NewEncoder(w).Encode(productsList)
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product, found := products[id]
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Status OK (200)
	json.NewEncoder(w).Encode(product)
}

func getProductsByQueryParamsPriceGT(w http.ResponseWriter, r *http.Request) {

	price, err := strconv.Atoi(r.URL.Query().Get("price"))

	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		fmt.Printf("Error getting price from query params: %v\n", err)
		return
	}

	productsList := make([]Product, 0, len(products))
	for _, product := range products {
		if product.Price > float64(price) {
			productsList = append(productsList, product)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Status OK (200)
	json.NewEncoder(w).Encode(productsList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if newProduct.Name == "" || newProduct.Quantity <= 0 || newProduct.CodeValue == "" {
		http.Error(w, "Name, Quantity, and CodeValue cannot be empty", http.StatusBadRequest)
		return
	}

	_, err = time.Parse("02/01/2006", newProduct.Expiration)
	if err != nil {
		http.Error(w, "Invalid date format for expiration", http.StatusBadRequest)
		return
	}

	newProduct.ID = nextID
	nextID++

	for _, existingProduct := range products {
		if existingProduct.CodeValue == newProduct.CodeValue {
			http.Error(w, "CodeValue must be unique", http.StatusBadRequest)
			return
		}
	}

	products[newProduct.ID] = newProduct

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Status Created (201)
	json.NewEncoder(w).Encode(newProduct)
}

func loadProductsFromJSON(filePath string) ([]Product, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var products []Product

	if err := json.Unmarshal(fileContent, &products); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return products, nil
}

func initializeProducts(productsFromFile []Product) {
	for _, product := range productsFromFile {
		products[product.ID] = product
		if product.ID >= nextID {
			nextID = product.ID + 1
		}
	}
}
