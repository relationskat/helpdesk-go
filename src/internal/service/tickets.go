package service

import (
	"github.com/relationskat/helpdesk-go/src/internal/domain"
	"github.com/relationskat/helpdesk-go/src/internal/store"
	"strings"
)

type ticketsService struct {
	store store.TicketStore
}

func NewTickets(s store.TicketStore) TicketsService {
	return &ticketsService{store: s}
}

func (s *ticketsService) Create(title string) (domain.Ticket, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return domain.Ticket{}, domain.ErrInvalidTitle
	}
	return s.store.CreateTicket(title), nil
}

func (s *ticketsService) List() []domain.Ticket {
	return s.store.ListTickets()
}
