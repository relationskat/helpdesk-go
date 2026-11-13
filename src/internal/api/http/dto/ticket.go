package dto

import "github.com/relationskat/helpdesk-go/src/internal/domain"

type CreateTicketRequest struct {
	Title string `json:"title"`
}

type TicketResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func TicketFromDomain(item domain.Ticket) TicketResponse {
	return TicketResponse{ID: item.ID, Title: item.Title, Status: item.Status}
}
