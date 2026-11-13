package store

import (
	"github.com/relationskat/helpdesk-go/src/internal/domain"
	"sync"
)

type Memory struct {
	mu      sync.Mutex
	tickets []domain.Ticket
}

func NewMemory() *Memory {
	return &Memory{tickets: []domain.Ticket{}}
}

func (m *Memory) CreateTicket(title string) domain.Ticket {
	m.mu.Lock()
	defer m.mu.Unlock()
	item := domain.Ticket{ID: len(m.tickets) + 1, Title: title, Status: domain.StatusOpen}
	m.tickets = append(m.tickets, item)
	return item
}

func (m *Memory) ListTickets() []domain.Ticket {
	m.mu.Lock()
	defer m.mu.Unlock()
	items := make([]domain.Ticket, len(m.tickets))
	copy(items, m.tickets)
	return items
}
