package internal

import "errors"

var (
	ErrIdAlreadyExists = errors.New("id already exists")
	ErrProductNotFound = errors.New("product not found")

	ErrCodeValueAlreadyExists = errors.New("code value already exists")
)

type ProductRepository interface {
	GetAll() ([]*Product, error)
	Get(id int) (*Product, error)
	Create(p *Product) (*Product, error)
	Update(p *Product) (*Product, error)
	Delete(id int) error
	SearchByPrice(priceGT float64) ([]*Product, error)
}
