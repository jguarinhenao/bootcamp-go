package repository

import (
	"app/scaffolding/internal"
	"encoding/json"
	"fmt"
	"os"
)

func NewProductMap(db map[int]internal.Product, lastId int) *ProductMap {

	productsFromFile, err := loadProductsFromJSON("/Users/jguarinhenao/Desktop/bootcamp-go/supermarket-api/code/internal/products.json")
	if err != nil {
		fmt.Printf("Error loading products from JSON: %v\n", err)
		return nil
	}

	productMap := &ProductMap{
		db:     db,
		lastId: lastId,
	}

	(*productMap).initializeProducts(productsFromFile)

	return productMap
}

type ProductMap struct {
	db     map[int]internal.Product
	lastId int
}

func (r *ProductMap) GetAll() ([]*internal.Product, error) {
	products := make([]*internal.Product, 0, len(r.db))
	for _, product := range r.db {
		products = append(products, &product)
	}
	return products, nil
}

func (r *ProductMap) SearchByPrice(priceGT float64) ([]*internal.Product, error) {
	products := make([]*internal.Product, 0, len(r.db))
	for _, product := range r.db {
		if product.Price > priceGT {
			products = append(products, &product)
		}
	}
	return products, nil
}

func (r *ProductMap) Get(id int) (*internal.Product, error) {
	product, ok := r.db[id]
	if !ok {
		return nil, internal.ErrProductNotFound
	}
	return &product, nil
}

func (r *ProductMap) Create(p *internal.Product) (*internal.Product, error) {

	for _, v := range (*r).db {

		if v.CodeValue == p.CodeValue {
			return nil, fmt.Errorf("%w: code_value", internal.ErrCodeValueAlreadyExists)
		}
	}

	(*r).lastId++
	(*p).ID = (*r).lastId
	(*r).db[(*r).lastId] = *p
	return p, nil
}

func (r *ProductMap) Update(p *internal.Product) (*internal.Product, error) {
	_, ok := r.db[(*p).ID]
	if !ok {
		return nil, fmt.Errorf("%w", internal.ErrProductNotFound)
	}
	r.db[(*p).ID] = *p
	return p, nil

}

func (r *ProductMap) Delete(id int) error {
	_, ok := r.db[id]
	if !ok {
		return internal.ErrProductNotFound
	}
	delete(r.db, id)
	return nil
}

func loadProductsFromJSON(filePath string) ([]internal.Product, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var products []internal.Product

	if err := json.Unmarshal(fileContent, &products); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return products, nil
}

func (r *ProductMap) initializeProducts(productsFromFile []internal.Product) {
	for _, product := range productsFromFile {
		r.db[product.ID] = product
		r.lastId = product.ID
	}
}
