package controller

import (
	"app/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ControllerTicketDefault struct {
	// sv represents the service of the tickets
	sv *service.ServiceTicketDefault
}

func NewControllerTicketDefault(sv *service.ServiceTicketDefault) *ControllerTicketDefault {
	return &ControllerTicketDefault{
		sv: sv,
	}
}

func (c *ControllerTicketDefault) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func (c *ControllerTicketDefault) GetAverageByCountry(w http.ResponseWriter, r *http.Request) {
	dest := chi.URLParam(r, "dest")
	total, err := c.sv.GetAverageByCountry(dest)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		msg := fmt.Sprintf("No tickets found to %s\n", dest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Average of tickets to %s: %.02f%%\n", dest, total * 100)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(msg))
}

func (c *ControllerTicketDefault) GetByCountry(w http.ResponseWriter, r *http.Request) {
	dest := chi.URLParam(r, "dest")
	total, err := c.sv.GetByCountry(dest)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		msg := fmt.Sprintf("No tickets found to %s\n", dest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(msg))
		return
	}
	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Total of tickets to %s: %d\n", dest, total)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(msg))
}
