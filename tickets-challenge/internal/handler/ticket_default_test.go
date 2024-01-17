package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"tickets-challenge/internal"
	"tickets-challenge/internal/handler"
	"tickets-challenge/internal/repository"
	"tickets-challenge/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// TestServiceTicketDefault_GetTotalTickets test the GetTotalTickets method
func TestServiceTicketDefault_GetTotalTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - database: map
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
		}
		// - repository: map
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		// - handler
		hd := handler.NewTicketDefault(sv)

		//req is a http request
		req := httptest.NewRequest("GET", "/tickets", nil)
		//res is a http response
		res := httptest.NewRecorder()

		// act
		hd.GetTotalTickets()(res, req)

		// assert
		expectedStatusCode := http.StatusOK
		expectedBody := `{"data":1,"message":"get total tickets successfull"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

}

// TestServiceTicketDefault_GetTicketsAmountByDestinationCountry test the GetTicketsAmountByDestinationCountry method
func TestServiceTicketDefault_GetTicketsAmountByDestinationCountry(t *testing.T) {
	t.Run("success to get total tickets by destination country", func(t *testing.T) {
		// arrange
		// - database: map
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "Finland",
				Hour:    "10:00",
				Price:   100,
			},
		}

		// - repository: map
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		// - handler
		hd := handler.NewTicketDefault(sv)

		//act
		req := httptest.NewRequest("GET", "/ticket/getByCountry", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("dest", "Finland")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		hd.GetTicketsAmountByDestinationCountry()(res, req)

		//expectedStatusCode := http.StatusOK
		expectedBody := `{
						  "data":1,
						  "message":"get total tickets successfull"
						}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		//require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("error invalid destination", func(t *testing.T) {
		//arrange
		// - database: map
		db := make(map[int]internal.TicketAttributes, 0)
		// - repository: map
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		// - handler
		hd := handler.NewTicketDefault(sv)

		req := httptest.NewRequest("GET", "/ticket/getByCountry", nil)
		res := httptest.NewRecorder()

		hd.GetTicketsAmountByDestinationCountry()(res, req)

		expectedStatusCode := http.StatusBadRequest
		expectedBody := `{"message":"invalid destination", "status":"Bad Request"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

// TestServiceTicketDefault_GetPercentageTicketsByDestinationCountry test the GetPercentageTicketsByDestinationCountry method
func TestServiceTicketDefault_GetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get percentage tickets by destination country", func(t *testing.T) {
		// arrange
		// - database: map
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "Finland",
				Hour:    "10:00",
				Price:   100,
			},
		}

		// - repository: map
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		// - handler
		hd := handler.NewTicketDefault(sv)

		req := httptest.NewRequest("GET", "/ticket/getByCountry", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("dest", "Finland")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		hd.GetPercentageTicketsByDestinationCountry()(res, req)

		expectedStatusCode := http.StatusOK
		expectedBody := `{"data":1,"message":"get percentage tickets successfull"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})

	t.Run("error invalid destination", func(t *testing.T) {
		//arrange
		// - database: map
		db := make(map[int]internal.TicketAttributes, 0)
		// - repository: map
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		// - handler
		hd := handler.NewTicketDefault(sv)

		req := httptest.NewRequest("GET", "/ticket/getByCountry", nil)
		res := httptest.NewRecorder()

		hd.GetPercentageTicketsByDestinationCountry()(res, req)

		expectedStatusCode := http.StatusBadRequest
		expectedBody := `{"message":"invalid destination", "status":"Bad Request"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}
