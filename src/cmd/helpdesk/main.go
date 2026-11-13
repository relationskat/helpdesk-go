package main

import (
	api "github.com/relationskat/helpdesk-go/src/internal/api/http"
	"github.com/relationskat/helpdesk-go/src/internal/config"
	"github.com/relationskat/helpdesk-go/src/internal/service"
	"github.com/relationskat/helpdesk-go/src/internal/store"
	"log"
)

func main() {
	cfg := config.Load()
	tickets := service.NewTickets(store.NewMemory())
	server := api.NewServer(api.Deps{Config: cfg, Tickets: tickets})
	log.Printf("HelpDesk listening on %s", cfg.HTTP.Addr)
	log.Fatal(server.Start())
}
