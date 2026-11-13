package http

import (
	"encoding/json"
	"github.com/relationskat/helpdesk-go/src/internal/api/http/dto"
	"github.com/relationskat/helpdesk-go/src/internal/service"
	"net/http"
)

type ticketsHandler struct {
	tickets service.TicketsService
}

func (h *ticketsHandler) list(w http.ResponseWriter, r *http.Request) {
	items := h.tickets.List()
	response := make([]dto.TicketResponse, len(items))
	for i, item := range items {
		response[i] = dto.TicketFromDomain(item)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ticketsHandler) create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateTicketRequest
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "expected JSON with non-empty title (max 4096 bytes)", http.StatusBadRequest)
		return
	}
	item, err := h.tickets.Create(input.Title)
	if err != nil {
		http.Error(w, "expected JSON with non-empty title (max 4096 bytes)", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.TicketFromDomain(item))
}
