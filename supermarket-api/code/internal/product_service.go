package internal

import "errors"

var (
	ErrFieldRequired        = errors.New("field required")
	ErrInvalidField         = errors.New("invalid field")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrInvalidId            = errors.New("invalid id")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidQuantity      = errors.New("invalid quantity")
	ErrInvalidCodeValue     = errors.New("invalid code value")
	ErrInvalidIsPublished   = errors.New("invalid is published")
	ErrInvalidExpiration    = errors.New("invalid expiration")
	ErrInvalidPrice         = errors.New("invalid price")
)

type ProductService interface {
	GetAll() ([]*Product, error)
	Get(id int) (*Product, error)
	Create(p *Product) (*Product, error)
	Update(p *Product) (*Product, error)
	Delete(id int) error
	SearchByPrice(priceGT float64) ([]*Product, error)
}
