package http

import (
	"github.com/relationskat/helpdesk-go/src/internal/config"
	"github.com/relationskat/helpdesk-go/src/internal/service"
	nethttp "net/http"
)

type Deps struct {
	Config  *config.Config
	Tickets service.TicketsService
}

type Server struct {
	server *nethttp.Server
}

func NewServer(d Deps) *Server {
	mux := nethttp.NewServeMux()
	registerRoutes(mux, d)
	return &Server{server: &nethttp.Server{Addr: d.Config.HTTP.Addr, Handler: mux}}
}

func (s *Server) Start() error { return s.server.ListenAndServe() }

func (s *Server) Handler() nethttp.Handler { return s.server.Handler }
