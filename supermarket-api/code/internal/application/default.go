package application

import (
	"app/scaffolding/internal"
	"app/scaffolding/internal/handler"
	"app/scaffolding/internal/repository"
	"app/scaffolding/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewDefaultHTTP(addr string) *DefaultHTTP {
	// default config / values
	// ...
	return &DefaultHTTP{
		addr: addr,
	}
}

type DefaultHTTP struct {
	// addr is the address of the http server
	addr string
}

// Run runs the http server
func (h *DefaultHTTP) Run() (err error) {
	// // initialize dependencies
	// // - repository
	rp := repository.NewProductMap(make(map[int]internal.Product), 0)
	// // - service
	sv := service.NewProductDefault(rp)
	// // - handler
	hd := handler.NewDefaultProducts(sv)

	r := chi.NewRouter()

	r.Route("/products", func(rt chi.Router) {

		rt.Get("/", hd.GetAll())

		rt.Get("/{id}", hd.Get())

		rt.Put("/{id}", hd.Update())

		rt.Patch("/{id}", hd.UpdatePartial())

		rt.Get("/search", hd.SearchByPrice())

		rt.Post("/", hd.Create())

		rt.Delete("/{id}", hd.Delete())

	})

	// run http server
	err = http.ListenAndServe(h.addr, r)
	return
}
