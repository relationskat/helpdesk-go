package service

import "github.com/relationskat/helpdesk-go/src/internal/domain"

type TicketsService interface {
	Create(title string) (domain.Ticket, error)
	List() []domain.Ticket
}
