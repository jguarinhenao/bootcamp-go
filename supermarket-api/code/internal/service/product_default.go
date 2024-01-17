package service

import (
	"app/scaffolding/internal"
	"fmt"
	"time"
)

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

type ProductDefault struct {
	rp internal.ProductRepository
}

func (s *ProductDefault) GetAll() ([]*internal.Product, error) {
	return s.rp.GetAll()
}

func (s *ProductDefault) SearchByPrice(priceGT float64) ([]*internal.Product, error) {

	if priceGT < 0 {
		return nil, fmt.Errorf("%w: price", internal.ErrInvalidPrice)
	}

	return s.rp.SearchByPrice(priceGT)
}

func (s *ProductDefault) Get(id int) (*internal.Product, error) {

	if id < 0 {
		return nil, fmt.Errorf("%w", internal.ErrInvalidId)
	}

	product, err := s.rp.Get(id)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return product, nil
}

func (s *ProductDefault) Create(p *internal.Product) (*internal.Product, error) {

	if err := ValidateProduct(p); err != nil {
		return nil, fmt.Errorf("%w: id", err)
	}

	product, err := s.rp.Create(p)
	if err != nil {
		switch err {
		case internal.ErrIdAlreadyExists:
			err = fmt.Errorf("%w: id", internal.ErrIdAlreadyExists)
		case internal.ErrCodeValueAlreadyExists:
			err = fmt.Errorf("%w: code_value", internal.ErrCodeValueAlreadyExists)
		}
	}

	return product, err
}

func ValidateProduct(p *internal.Product) (err error) {
	if (*p).Name == "" {
		return fmt.Errorf("%w: name", internal.ErrFieldRequired)
	}

	if (*p).Quantity < 0 {
		return fmt.Errorf("%w: quantity", internal.ErrFieldRequired)
	}

	if (*p).CodeValue == "" {
		return fmt.Errorf("%w: code_value", internal.ErrFieldRequired)
	}

	if (*p).Price < 0 {
		return fmt.Errorf("%w: price", internal.ErrFieldRequired)
	}

	_, err = time.Parse("02/01/2006", (*p).Expiration)
	if err != nil {
		return fmt.Errorf("%w: Invalid date format for expiration", internal.ErrFieldRequired)
	}

	return
}

func (s *ProductDefault) Update(p *internal.Product) (product *internal.Product, err error) {

	if err = ValidateProduct(p); err != nil {
		return
	}

	product, err = s.rp.Update(p)

	if err != nil {

		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		case internal.ErrCodeValueAlreadyExists:
			err = fmt.Errorf("%w: code_value", internal.ErrCodeValueAlreadyExists)
		}
	}
	return

}

func (s *ProductDefault) Delete(id int) error {
	return s.rp.Delete(id)
}
