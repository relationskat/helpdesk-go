package http

import "net/http"

func registerRoutes(mux *http.ServeMux, d Deps) {
	tickets := &ticketsHandler{tickets: d.Tickets}
	mux.HandleFunc("GET /healthz", health)
	mux.HandleFunc("GET /tickets", tickets.list)
	mux.HandleFunc("POST /tickets", tickets.create)
}
