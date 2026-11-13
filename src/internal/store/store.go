package store

import "github.com/relationskat/helpdesk-go/src/internal/domain"

type TicketStore interface {
	CreateTicket(title string) domain.Ticket
	ListTickets() []domain.Ticket
}
