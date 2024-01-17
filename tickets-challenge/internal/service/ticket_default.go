package service

import (
	"context"
	"fmt"
	"tickets-challenge/internal"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	//service third
	// ....
	//buisness logic
	// ....
	tickets, err := s.rp.Get(context.Background())
	if err != nil {
		return 0, fmt.Errorf("error getting the tickets: %w", err)
	}
	return len(tickets), nil
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(destination string) (total int, err error) {
	//service third
	// ....
	//buisness logic
	if destination == "" {
		err = internal.ErrDestinationEmpty
		return
	}
	tickets, err := s.rp.GetTicketsByDestinationCountry(context.Background(), destination)
	if err != nil {
		return 0, fmt.Errorf("error getting the tickets: %w", err)
	}
	return len(tickets), nil
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(destination string) (percentage float64, err error) {
	//service third
	// ....
	//buisness logic
	if destination == "" {
		err = internal.ErrDestinationEmpty
		return
	}

	//destinationTickets is the amount of tickets filtered by destination country
	destinationTickets, err := s.rp.GetTicketsByDestinationCountry(context.Background(), destination)
	if err != nil {
		return 0, fmt.Errorf("error getting the tickets: %w", err)
	}

	//allTickets is the amount of tickets
	allTickets, err := s.rp.Get(context.Background())
	if err != nil {
		return 0, fmt.Errorf("error getting the tickets: %w", err)
	}

	totalDestinationTicket := len(destinationTickets)
	totalTicket := len(allTickets)

	return float64(totalDestinationTicket) / float64(totalTicket), nil
}
