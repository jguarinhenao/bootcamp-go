package internal

import (
	"errors"
)

var (
	ErrDestinationEmpty = errors.New("service: error destination is empty")
)

type ServiceTicket interface {
	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalAmountTickets() (total int, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	GetTicketsAmountByDestinationCountry(destination string) (total int, err error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(destination string) (percentage float64, err error)
}
