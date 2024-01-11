package handler

import (
	"app/scaffolding/internal"
	"app/scaffolding/platform/web/request"
	"app/scaffolding/platform/web/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewDefaultProducts(sv internal.ProductService) *DefaultProducts {
	return &DefaultProducts{
		sv: sv,
	}
}

type DefaultProducts struct {
	// sv is a movie service
	sv internal.ProductService
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (d *DefaultProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body BodyRequestProductJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		if _, err := d.sv.Create(&product); err != nil {
			// print error
			fmt.Println(err)

			switch {
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrInvalidCodeValue), errors.Is(err, internal.ErrInvalidExpiration), errors.Is(err, internal.ErrInvalidPrice), errors.Is(err, internal.ErrInvalidQuantity), errors.Is(err, internal.ErrInvalidName), errors.Is(err, internal.ErrInvalidId):
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				response.Text(w, http.StatusConflict, "product already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize MovieJSON
		data := ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "product created",
			"data":    data,
		})
	}
}

func (d *DefaultProducts) Get() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get product
		product, err := d.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize MovieJSON
		data := ProductJSON{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,

			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		// - response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product found",
			"data":    data,
		})
	}
}

func (d *DefaultProducts) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// process
		// - get products
		products, err := d.sv.GetAll()
		if err != nil {
			switch {
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize MovieJSON
		data := make([]ProductJSON, 0, len(products))
		for _, product := range products {
			data = append(data, ProductJSON{
				ID:          product.ID,
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration,
				Price:       product.Price,
			})
		}
		// - response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "products found",
			"data":    data,
		})
	}
}

func (d *DefaultProducts) SearchByPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// request
		// - get id from query Param
		price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid price")
			return
		}

		// process
		// - get product
		products, err := d.sv.SearchByPrice(price)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		data := make([]ProductJSON, 0, len(products))
		for _, product := range products {
			data = append(data, ProductJSON{
				ID:          product.ID,
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration,
				Price:       product.Price,
			})
		}

		// - response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product found",
			"data":    data,
		})
	}

}

func (d *DefaultProducts) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		bytes, err := io.ReadAll(r.Body)

		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		var bodyMap map[string]any

		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		if err := ValidateExistingKey(bodyMap, "name",
			"quantity",
			"code_value",
			"is_published",
			"expiration",
			"price"); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		var body BodyRequestProductJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			// if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// process
		// - serialize internal.Movie
		product := internal.Product{
			ID:          id,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		// - update movie
		if _, err := d.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrInvalidQuantity), errors.Is(err, internal.ErrInvalidName), errors.Is(err, internal.ErrInvalidCodeValue), errors.Is(err, internal.ErrInvalidExpiration), errors.Is(err, internal.ErrInvalidPrice):
				response.Text(w, http.StatusBadRequest, "invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize MovieJSON
		data := ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// - response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})

	}

}

// UpdatePartial updates a product
func (d *DefaultProducts) UpdatePartial() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - get the movie from the service
		product, err := d.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// process
		// - serialize internal.Movie to BodyRequestMovieJSON
		reqBody := BodyRequestProductJSON{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// - get body
		if err := request.JSON(r, &reqBody); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - serialize internal.Product
		product = &internal.Product{
			ID:          id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		// - update movie
		if _, err := d.sv.Update(product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrInvalidQuantity), errors.Is(err, internal.ErrInvalidName), errors.Is(err, internal.ErrInvalidCodeValue), errors.Is(err, internal.ErrInvalidExpiration), errors.Is(err, internal.ErrInvalidPrice):
				response.Text(w, http.StatusBadRequest, "invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize MovieJSON
		data := ProductJSON{
			ID:          id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})

	}

}

func (d *DefaultProducts) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id",
				"data":    nil,
			})
			return
		}

		// proces
		// - delete movie
		if err := d.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.Text(w, http.StatusNoContent, "movie deleted")
	}
}

func ValidateExistingKey(mp map[string]any, keys ...string) (err error) {
	for _, k := range keys {
		if _, ok := mp[k]; !ok {
			return fmt.Errorf("key %s not found", k)
		}
	}
	return
}
