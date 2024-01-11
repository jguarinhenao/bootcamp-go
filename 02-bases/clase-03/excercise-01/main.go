package main

import "fmt"

// Product estructura
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Products slice global
var Products []Product

// Save agrega el producto al slice Products
func (p *Product) Save() {
	Products = append(Products, *p)
}

// GetAll imprime todos los productos guardados
func (p *Product) GetAll() {
	fmt.Println("Lista de productos:")
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

// getById retorna el producto correspondiente al ID proporcionado
func getById(id int) *Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func main() {
	// Instanciar algunos productos
	product1 := Product{ID: 1, Name: "Producto 1", Price: 19.99, Description: "Descripción 1", Category: "Categoría 1"}
	product2 := Product{ID: 2, Name: "Producto 2", Price: 29.99, Description: "Descripción 2", Category: "Categoría 2"}

	// Ejecutar el método Save para agregar los productos al slice
	product1.Save()
	product2.Save()

	// Ejecutar el método GetAll para imprimir todos los productos
	var tempProduct Product
	tempProduct.GetAll()

	// Ejecutar la función getById para obtener un producto por ID
	idToFind := 1
	foundProduct := getById(idToFind)
	if foundProduct != nil {
		fmt.Printf("\nProducto encontrado por ID (%d):\n", idToFind)
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n",
			foundProduct.ID, foundProduct.Name, foundProduct.Price, foundProduct.Description, foundProduct.Category)
	} else {
		fmt.Printf("\nProducto con ID %d no encontrado.\n", idToFind)
	}
}
