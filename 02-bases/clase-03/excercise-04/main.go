package main

import "fmt"

// Product interfaz
type Product interface {
	Price() float64
}

// ProductFactory función factory para crear productos
func ProductFactory(productType string, price float64) Product {
	switch productType {
	case "Small":
		return &SmallProduct{BasePrice: price}
	case "Medium":
		return &MediumProduct{BasePrice: price}
	case "Large":
		return &LargeProduct{BasePrice: price}
	default:
		return nil
	}
}

// SmallProduct estructura para productos Small
type SmallProduct struct {
	BasePrice float64
}

// Price método para calcular el precio total de productos Small
func (p *SmallProduct) Price() float64 {
	return p.BasePrice
}

// MediumProduct estructura para productos Medium
type MediumProduct struct {
	BasePrice float64
}

// Price método para calcular el precio total de productos Medium
func (p *MediumProduct) Price() float64 {
	return p.BasePrice * 1.03
}

// LargeProduct estructura para productos Large
type LargeProduct struct {
	BasePrice float64
}

// Price método para calcular el precio total de productos Large
func (p *LargeProduct) Price() float64 {
	return p.BasePrice*1.06 + 2500
}

func main() {
	// Crear productos utilizando la función factory
	smallProduct := ProductFactory("Small", 100.0)
	mediumProduct := ProductFactory("Medium", 200.0)
	largeProduct := ProductFactory("Large", 300.0)

	// Calcular y mostrar el precio total de cada producto
	fmt.Printf("Small Product Price: $%.2f\n", smallProduct.Price())
	fmt.Printf("Medium Product Price: $%.2f\n", mediumProduct.Price())
	fmt.Printf("Large Product Price: $%.2f\n", largeProduct.Price())
}
