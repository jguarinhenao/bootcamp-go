package handler

import (
	"errors"
	"net/http"
	"tickets-challenge/internal"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type TicketDefault struct {
	sv internal.ServiceTicket
}

func NewTicketDefault(sv internal.ServiceTicket) *TicketDefault {
	return &TicketDefault{
		sv: sv,
	}
}

func (h *TicketDefault) GetTotalTickets() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		totalTickets, err := h.sv.GetTotalAmountTickets()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "get total tickets successfull",
			"data":    totalTickets,
		})
	}
}

func (h *TicketDefault) GetTicketsAmountByDestinationCountry() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		destination := chi.URLParam(r, "dest")

		totalTickets, err := h.sv.GetTicketsAmountByDestinationCountry(destination)
		if err != nil {
			if errors.Is(err, internal.ErrDestinationEmpty) {
				response.Error(w, http.StatusBadRequest, "invalid destination")
				return
			}
			response.Error(w, http.StatusInternalServerError, "unknown error")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "get total tickets successfull",
			"data":    totalTickets,
		})
	}
}

func (h *TicketDefault) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		destination := chi.URLParam(r, "dest")

		percentageTickets, err := h.sv.GetPercentageTicketsByDestinationCountry(destination)
		if err != nil {
			if errors.Is(err, internal.ErrDestinationEmpty) {
				response.Error(w, http.StatusBadRequest, "invalid destination")
				return
			}
			response.Error(w, http.StatusInternalServerError, "unknown error")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "get percentage tickets successfull",
			"data":    percentageTickets,
		})
	}
}
