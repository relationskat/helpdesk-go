package http

import (
	"encoding/json"
	"github.com/relationskat/helpdesk-go/src/internal/api/http/dto"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.HealthResponse{Status: "ok"})
}
