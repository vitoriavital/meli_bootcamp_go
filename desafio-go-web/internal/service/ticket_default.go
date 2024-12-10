package service

import (
	"app/internal"
	"context"
	"errors"
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

// GetTotalAmountTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	totalTickets, err := s.rp.Get(context.Background())
	if err != nil {
		return 
	}

	total = len(totalTickets)
	return 
}

// GetByCountry returns the the total number of tickets for that dest
func (s *ServiceTicketDefault) GetByCountry(dest string) (total int, err error) {
	totalTickets, err := s.rp.GetTicketsByDestinationCountry(context.Background(), dest)
	if err != nil {
		return 
	}

	total = len(totalTickets)
	return 
}

// GetAverageByCountry returns the the average of tickets for that dest
func (s *ServiceTicketDefault) GetAverageByCountry(dest string) (average float64, err error) {
	totalTickets, err := s.rp.Get(context.Background())
	if err != nil {
		return 
	}
	totalDestTickets, err := s.rp.GetTicketsByDestinationCountry(context.Background(), dest)
	if err != nil {
		return 
	}
	totalAmount := len(totalTickets)
	destAmount := len(totalDestTickets)
	if destAmount < 1 {
		err = errors.New("no tickets for this dest")
		return 
	}
	average = float64(destAmount) / float64(totalAmount)
	return 
}